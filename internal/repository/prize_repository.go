package repository

import (
	"context"
	"hsr/loto/internal/db"
	"hsr/loto/internal/entity"
)

type PrizeRepository struct {
	db *db.DB
}

type PrizeInterface interface {
	Get(id int) (*entity.Prize, error)
	List() (*[]entity.Prize, error)
	GetPrizes() (*[]entity.Prize, error)
	Add(t *entity.Prize) (*entity.Prize, error)
	DeletePrize(id int) (bool, error)
}

func NewPrizeRepository(d *db.DB) *PrizeRepository {
	return &PrizeRepository{db: d}
}

func (r *PrizeRepository) GetPrize(id int) (*entity.Prize, error) {
	e := entity.Prize{}

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
func (r *PrizeRepository) GetPrizes() (*[]entity.Prize, error) {
	d := []entity.Prize{}

	rows, err := r.db.Client.
		Query(context.Background(), "SELECT id, prize FROM hsrloto.prizes limit 120;")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e entity.Prize
		err := rows.Scan(&e.Id, &e.Prize)

		if err != nil {
			return nil, err
		}

		d = append(d, e)
	}

	return &d, nil
}

func (r *PrizeRepository) ListPrizes() (*[]entity.Prize, error) {
	var d []entity.Prize

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

func (r *PrizeRepository) AddPrize(e *entity.Prize) (*entity.Prize, error) {
	//err := r.db.Client.
	//	QueryRow(context.Background(), "INSERT INTO test(status, created_at, updated_at) VALUES ($1, now(), now()) RETURNING id", e.Status).
	//	Scan(&e.Id)

	//if err != nil {
	//	return nil, err
	//}

	return e, nil
}

func (r *PrizeRepository) DeletePrize(id int) (bool, error) {
	if _, err := r.GetPrize(id); err != nil {
		return false, err
	}

	//_, err := r.db.Client.
	//	Query(context.Background(), "UPDATE test SET deleted_at = now() WHERE id = $1", id)

	//if err != nil {
	//	return false, err
	//}

	return true, nil
}

func (r *TicketRepository) SeedPrizeTable() error {
	var prizes []string
	for i := 0; i <= 99; i++ {
		prizes = append(prizes, "Сертификат на 1000 руб")
	}
	for i := 100; i <= 110; i++ {
		prizes = append(prizes, "Сертификат на 4000 руб")
	}
	for i := 111; i <= 115; i++ {
		prizes = append(prizes, "Смартфон")
	}
	for i := 116; i <= 118; i++ {
		prizes = append(prizes, "Планшет")
	}
	prizes = append(prizes, "Суперприз телевизор (100 000)")

	for i := 0; i <= 119; i++ {
		rows, err := r.db.Client.
			Query(context.Background(),
				"INSERT INTO hsrloto.prizes (prize) VALUES ($1);", prizes[i])

		if err != nil {
			return err
		}

		rows.Close()
	}

	return nil
}
