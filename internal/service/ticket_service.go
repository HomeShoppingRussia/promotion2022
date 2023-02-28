package service

import (
	"context"
	"hsr/loto/internal/db"
	"hsr/loto/internal/entity"
	"hsr/loto/internal/repository"
)

type Repository struct {
	db *db.DB
}

type TicketService struct {
	r *repository.TicketRepository
}

type TicketInterface interface {
	GetTicket(int) (*entity.Ticket, error)
	GetTicketList() (*[]entity.Ticket, error)
	AddTicket(m *entity.Ticket) (*entity.Ticket, error)
	AddTicketList(m []*entity.Ticket) ([]*entity.Ticket, error)
	DeleteTicket(id int) (bool, error)
	SeedTicketTable() error
}

func NewTicket(r *repository.TicketRepository) *TicketService {
	return &TicketService{r}
}

func (s *TicketService) GetTicket(id int) (*entity.Ticket, error) {
	return s.r.GetTicket(id)
}

func (s *TicketService) GetTicketList() (*[]entity.Ticket, error) {
	return s.r.GetTickets()
}

func (s *TicketService) SetPrize(prizeId int, ticket entity.Ticket) error {
	return s.r.SetPrize(prizeId, ticket)
}

func (s *TicketService) AddTicket(m *entity.Ticket) (*entity.Ticket, error) {
	return s.r.AddTicket(m)
}

func (s *TicketService) AddTicketList(m []entity.Ticket) error {
	return s.r.AddTicketList(m)
}

func (s *TicketService) DeleteTicket(id int) (bool, error) {
	return s.r.DeleteTicket(id)
}

func (s *TicketService) SeedTicketTable() error {
	return s.r.SeedTicketTable()
}

func (s *TicketService) CheckDrawStatus(ctx context.Context) (int, error) {
	return s.r.CheckDrawStatus(ctx)
}
