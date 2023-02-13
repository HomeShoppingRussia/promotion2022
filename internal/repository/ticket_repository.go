package repository

import (
	"context"
	"hsr/loto/internal/db"
	"hsr/loto/internal/entity"
	"hsr/loto/internal/utils"
	"strconv"
)

const PRIZE_COUNT = 120

type TicketRepository struct {
	db *db.DB
}

type TicketInterface interface {
	GetTicket(id int) (*entity.Ticket, error)
	GetTickets() (*[]entity.Ticket, error)
	ListTickets() (*[]entity.Ticket, error)
	AddTicket(e *entity.Ticket) (*entity.Ticket, error)
	DeleteTicket(id int) (bool, error)
}

func NewTicketRepository(d *db.DB) *TicketRepository {
	return &TicketRepository{db: d}
}

func (r *TicketRepository) GetTicket(id int) (*entity.Ticket, error) {
	e := entity.Ticket{}

	//err := r.db.Client.
	//	QueryRow(context.Background(), "SELECT id FROM test WHERE id = $1 AND deleted_at IS NULL", id).
	//	Scan(&e.Id)

	//if err != nil {
	//	return nil, err
	//}

	return &e, nil
}

/**
 * Функция выбора рандомных билетов. Это можно ускорить, см.https://stackoverflow.com/questions/8674718/best-way-to-select-random-rows-postgresql
 */
func (r *TicketRepository) GetTickets() (*[]entity.Ticket, error) {
	d := []entity.Ticket{}

	rows, err := r.db.Client.
		Query(context.Background(), "SELECT id, prize_id, name, surname, middle_name, phone, spi FROM hsrloto.tickets order by random() limit 120")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.Ticket
		err := rows.Scan(&e.Id, &e.PrizeId, &e.Name, &e.Surname, &e.MiddleName, &e.Phone, &e.Spi)

		if err != nil {
			return nil, err
		}

		d = append(d, e)
	}

	return &d, nil
}

/**
 * Функция назначения билету приза
 */
func (r *TicketRepository) SetPrize(prizeId int, ticket entity.Ticket) error {
	rows, err := r.db.Client.
		Query(context.Background(), "UPDATE hsrloto.tickets SET prize_id = "+strconv.Itoa(prizeId)+" WHERE hsrloto.tickets.id = "+strconv.FormatInt(ticket.Id, 10))
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (r *TicketRepository) ListTickets() (*[]entity.Ticket, error) {
	var d []entity.Ticket

	//rows, err := r.db.Client.Query(context.Background(), "SELECT id, status FROM test WHERE deleted_at IS NULL")

	//if err != nil {
	//	return nil, err
	//}

	//for rows.Next() {
	//	var e entity.Test
	//	err := rows.Scan(&e.Id, &e.Status)

	//	if err != nil {
	//		return nil, err
	//	}

	//	d = append(d, e)
	//}

	return &d, nil
}

func (r *TicketRepository) AddTicket(e *entity.Ticket) (*entity.Ticket, error) {
	err := r.db.Client.
		QueryRow(context.Background(), "INSERT INTO hsrloto.tickets (id, prize_id, name, surname, middle_name) VALUES ($1, $2, $3, $4, $5) RETURNING id", e.Id, e.PrizeId.Get(), e.Name, e.Surname, e.MiddleName).
		Scan(&e.Id)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *TicketRepository) AddTicketList(e []entity.Ticket) error {
	for _, t := range e {
		rows, err := r.db.Client.
			Query(context.Background(),
				"INSERT INTO hsrloto.tickets (name, surname, middle_name, phone, spi) VALUES ($1, $2, $3, $4, $5);", t.Name, t.Surname, t.MiddleName, t.Phone, t.Spi)

		if err != nil {
			return err
		}

		rows.Close()
	}

	return nil
}

func (r *TicketRepository) DeleteTicket(id int) (bool, error) {
	if _, err := r.GetTicket(id); err != nil {
		return false, err
	}

	//_, err := r.db.Client.
	//	Query(context.Background(), "UPDATE test SET deleted_at = now() WHERE id = $1", id)

	//if err != nil {
	//	return false, err
	//}

	return true, nil
}

func (r *TicketRepository) SeedTicketTable() error {
	//var ticketNumber int
	names := make(map[int]string)
	names[0] = "Роман"
	names[1] = "Владимир"
	names[2] = "Константин"
	names[3] = "Анна"
	names[4] = "Мария"
	names[5] = "Виктория"

	surnames := make(map[int]string)
	surnames[0] = "Ким"
	surnames[1] = "Вик"
	surnames[2] = "Шапиро"
	surnames[3] = "Никитенко"
	surnames[4] = "Черных"
	surnames[5] = "Долгих"

	middleNames := make(map[int]string)
	middleNames[0] = "Павлович"
	middleNames[1] = "Владимирович"
	middleNames[2] = "Олегович"
	middleNames[3] = "Петрович"
	middleNames[4] = "Иванович"
	middleNames[5] = "Романович"
	for i := 0; i < 120; i++ {
		name := names[utils.RandInt(0, 5)]
		surname := surnames[utils.RandInt(0, 5)]
		middleName := middleNames[utils.RandInt(0, 5)]
		phone := "*******" + strconv.Itoa(utils.RandInt(0, 9)) + strconv.Itoa(utils.RandInt(0, 9)) + strconv.Itoa(utils.RandInt(0, 9)) + strconv.Itoa(utils.RandInt(0, 9))
		spiStr := "111753753" + strconv.Itoa(utils.RandInt(0, 9)) + strconv.Itoa(utils.RandInt(0, 9)) + strconv.Itoa(utils.RandInt(0, 9)) + strconv.Itoa(utils.RandInt(0, 9)) + strconv.Itoa(utils.RandInt(0, 9))
		spi, err := strconv.Atoi(spiStr)
		if err != nil {
			return err
		}
		rows, err := r.db.Client.
			Query(context.Background(),
				"INSERT INTO hsrloto.tickets (name, surname, middle_name, phone, spi) VALUES ($1, $2, $3, $4, $5);", name, surname, middleName, phone, spi)

		if err != nil {
			return err
		}

		rows.Close()
	}

	return nil
}
