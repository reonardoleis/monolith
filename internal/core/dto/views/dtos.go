package views_dto

type AddViewRequest struct {
	Origin string `json:"origin"`
}

type AddViewResponse struct{}

type GetViewCountResponse struct {
	Views int `json:"views"`
}
