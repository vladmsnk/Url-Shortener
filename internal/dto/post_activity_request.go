package dto

import (
	"time"
	"vladmsnk/taskrec/internal/entity"
)

type PostActivityRequest struct {
	Title         string    `json:"s_title"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	AvailableFrom time.Time `json:"available_from"`
	AvailableTo   time.Time `json:"available_to"`
}

func (p PostActivityRequest) FromDto() entity.Activity {
	return entity.Activity{
		Title:         p.Title,
		Description:   p.Description,
		Price:         p.Price,
		AvailableFrom: p.AvailableFrom.String(),
		AvailableTo:   p.AvailableTo.String(),
	}
}
