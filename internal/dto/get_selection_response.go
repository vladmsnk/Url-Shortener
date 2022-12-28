package dto

type GetSelectionResponse struct {
	Title      string `json:"title"`
	Activities [3]PostActivityRequest
}
