package views_domain

import (
	"context"

	views_domain "github.com/reonardoleis/views/internal/core/domain/views"
	views_dto "github.com/reonardoleis/views/internal/core/dto/views"
)

type ViewService interface {
	AddView(ctx context.Context, ip string, req *views_dto.AddViewRequest) (*views_dto.AddViewResponse, error)
	GetViewCount(ctx context.Context) (*views_dto.GetViewCountResponse, error)
}

type ViewRepository interface {
	CreateView(ctx context.Context, ip, origin string) (*views_domain.View, error)
	CountViews(ctx context.Context) (int, error)
}
