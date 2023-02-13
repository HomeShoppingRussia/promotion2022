package service

import (
	"hsr/loto/internal/entity"
	"hsr/loto/internal/repository"
)

type PrizeService struct {
	r *repository.PrizeRepository
}

type PrizeInterface interface {
	GetPrize(int) (*entity.Prize, error)
	GetPrizeList() (*[]entity.Prize, error)
	AddPrize(m *entity.Prize) (*entity.Prize, error)
	DeletePrize(id int) (bool, error)
}

func NewPrize(r *repository.PrizeRepository) *PrizeService {
	return &PrizeService{r}
}

func (s *PrizeService) GetPrize(id int) (*entity.Prize, error) {
	return s.r.GetPrize(id)
}

func (s *PrizeService) GetPrizeList() (*[]entity.Prize, error) {
	return s.r.GetPrizes()
}

func (s *PrizeService) AddPrize(m *entity.Prize) (*entity.Prize, error) {
	return s.r.AddPrize(m)
}

func (s *PrizeService) DeletePrize(id int) (bool, error) {
	return s.r.DeletePrize(id)
}

func (s *TicketService) SeedPrizeTable() error {
	return s.r.SeedPrizeTable()
}
