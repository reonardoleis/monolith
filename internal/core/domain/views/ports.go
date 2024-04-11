package views_domain

import (
	"context"

	views_dto "github.com/reonardoleis/views/internal/core/dto/views"
)

type ViewUsecase interface {
	AddView(ctx context.Context, ip string, req *views_dto.AddViewRequest) (*views_dto.AddViewResponse, error)
}

type ViewRepository interface {
	CreateView(ctx context.Context, ip, origin string) (*View, error)
}
