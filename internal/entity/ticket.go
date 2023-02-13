package entity

import "hsr/loto/pkg/nulltype"

type Ticket struct {
	Id         int64              `json:"id"`
	PrizeId    nulltype.NullInt64 `json:"prize_id, omitempty"`
	Name       string             `json:"name"`
	Surname    string             `json:"surname"`
	MiddleName string             `json:"middle_name"`
	Phone      string             `json:"phone"`
	Spi        string             `json:"spi"`
}
