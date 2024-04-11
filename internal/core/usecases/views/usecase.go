package views_usecase

import (
	"context"

	dto "github.com/reonardoleis/views/internal/core/dto/views"
)

func (u usecase) AddView(ctx context.Context, ip string, req *dto.AddViewRequest) (*dto.AddViewResponse, error) {
	_, err := u.repo.CreateView(ctx, ip, req.Origin)
	if err != nil {
		return nil, err
	}

	return &dto.AddViewResponse{}, nil
}
