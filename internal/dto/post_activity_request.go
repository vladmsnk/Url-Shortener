package dto

import (
	"time"
	"vladmsnk/taskrec/internal/entity"
)

type ActivityDTO struct {
	Title         string    `json:"s_title"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	AvailableFrom time.Time `json:"available_from"`
	AvailableTo   time.Time `json:"available_to"`
}

func (p ActivityDTO) FromDto() entity.Activity {
	return entity.Activity{
		Title:         p.Title,
		Description:   p.Description,
		Price:         p.Price,
		AvailableFrom: p.AvailableFrom,
		AvailableTo:   p.AvailableTo,
	}
}

func (p ActivityDTO) ToDto(activity entity.Activity) ActivityDTO {
	return ActivityDTO{Title: activity.Title,
		Description:   activity.Description,
		Price:         activity.Price,
		AvailableFrom: activity.AvailableFrom,
		AvailableTo:   activity.AvailableTo}
}
