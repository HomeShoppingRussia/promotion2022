package server

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	pb "hsr/loto/api/proto/loto"
	"hsr/loto/internal/entity"
	"log"
	"os"
	"strconv"
	"sync"

	//es "hsr/test-service/api/proto/external/external_service"
	"hsr/loto/internal/logger"
	"hsr/loto/internal/service"
)

const ChunkSize = 50

type Server struct {
	*pb.UnimplementedLotoServer
	*service.TicketService
	*service.PrizeService
	//ExternalServiceClient es.ExternalServiceClient
}

func (s *Server) GetWinners(ctx context.Context, in *pb.LotoRequest) (*pb.LotoResponse, error) {
	fmt.Println("GetWinners called")

	//header := metadata.Pairs("Access-Control-Allow-Origin", "*", "Access-Control-Allow-Methods", "GET,POST,PATCH,UPDATE,DELETE,OPTIONS", "Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization", "Access-Control-Allow-Credentials", "true")
	grpc.SendHeader(ctx, metadata.Pairs("Access-Control-Allow-Origin", "*"))

	fmt.Println("GetTicketList calling")
	t, err := s.GetTicketList()
	if err != nil {
		return nil, err
	}
	count, err := s.CheckDrawStatus(ctx)
	if err != nil {
		return nil, err
	}
	if count != 0 {
		return &pb.LotoResponse{
			Status: false,
			Error:  "Розыгрыш уже проводился. Все победители определены",
			List:   nil,
		}, nil
	}
	p, err := s.GetPrizeList()
	if err != nil {
		fmt.Println("GetPrizeList error" + err.Error())
		return nil, err
	}
	fmt.Println("ticket length: " + strconv.Itoa(len(*t)) + " prize length: " + strconv.Itoa(len(*p)))
	var services []*pb.Participants
	for i, ticket := range *t {
		prize := make([]entity.Prize, len(*p))
		copy(prize, *p)
		fmt.Println("prize in if: " + strconv.Itoa(len(prize)))
		if len(prize) >= 120 {
			err := s.SetPrize(prize[i].Id, ticket)
			if err != nil {
				fmt.Println("SetPrize error" + err.Error())
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("does not enough tickets")
		}

		participant := &pb.Participants{
			TicketNumber: ticket.Id,
			Name:         ticket.Name,
			Surname:      ticket.Surname,
			MiddleName:   ticket.MiddleName,
			Phone:        ticket.Phone,
			Spi:          ticket.Spi,
			Prize:        prize[i].Prize,
		}
		services = append(services, participant)
	}
	log.Print(services)
	return &pb.LotoResponse{
		Status: true,
		Error:  "",
		List:   services,
	}, nil
}

// Define the struct to hold the data from csv
type Record struct {
	Field1     string
	Id         string
	Phone      string
	FirstName  string
	LastName   string
	MiddleName string
	//Gender     string
	UserID string
}

func (s *Server) UploadData(ctx context.Context, in *pb.UploadDataRequest) (*pb.UploadDataResponse, error) {
	t, _ := s.GetTicketList()
	logger.Info.Print("Numbers of elements: " + strconv.Itoa(len(*t)))
	if len(*t) > 0 {
		return nil, fmt.Errorf("DB is already filled")
	}
	fmt.Println(t)
	logger.Info.Print("Upload data from CSV")
	// Open the CSV file
	file, err := os.Open("data.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Create a new CSV reader
	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ';'
	allRowsCount := 0
	var wg sync.WaitGroup
	if err != nil {
		return nil, err
	}
	// Create a slice to hold the structs
	var tickets []entity.Ticket

	// Wrap parsing into go-rutines for faster interact
	wg.Add(1)
	go func() {
		err := func() error {
			wg.Add(1)
			counter := 0
			// Loop through each record and parse the data
			for {
				// Read next row from reader
				row, err := reader.Read()
				if err != nil {
					break
				}

				if counter >= ChunkSize {
					err := s.AddTicketList(tickets)
					if err != nil {
						return err
					}
					// Clean up var tickets
					tickets = []entity.Ticket{}
					counter = 0
				}

				id, err := strconv.ParseInt(row[5], 10, 64)
				tickets = append(tickets, entity.Ticket{
					Spi:        row[0],
					Phone:      row[1],
					Name:       row[2],
					MiddleName: row[3],
					Surname:    row[4],
					Id:         id,
				})
				counter++
				allRowsCount++
				if allRowsCount == 140 {
					fmt.Println("debug")
				}
			}
			wg.Done()
			return nil
		}()
		if err != nil {
		}
	}()
	wg.Wait()
	fmt.Println("allRowsCount is: ")
	fmt.Println(allRowsCount)

	return &pb.UploadDataResponse{
		Success: true,
		Message: "",
	}, nil

}

func (s *Server) FillPrizesTable(ctx context.Context, in *pb.UploadDataRequest) (*pb.UploadDataResponse, error) {
	fmt.Println("FillPrizesTable started")
	err := s.SeedPrizeTable()
	if err != nil {
		fmt.Println("FillPrizesTable error: " + err.Error())
		return &pb.UploadDataResponse{
			Success: false,
			Message: "Table prozes is filled",
		}, err
	}
	return &pb.UploadDataResponse{
		Success: true,
		Message: "Table prozes is filled",
	}, nil
}
