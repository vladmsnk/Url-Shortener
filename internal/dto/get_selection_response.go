package dto

import "vladmsnk/taskrec/internal/entity"

type GetSelectionResponse struct {
	Title      string `json:"title"`
	Activities []ActivityDTO
}

func (s GetSelectionResponse) ToDto(title string, activities []entity.Activity) GetSelectionResponse {
	var selection GetSelectionResponse

	selection.Title = title
	for _, activity := range activities {
		activityDTO := ActivityDTO{}.ToDto(activity)
		selection.Activities = append(selection.Activities, activityDTO)
	}
	return selection
}
