package views_service

import (
	"context"
	"log"

	dto "github.com/reonardoleis/views/internal/core/dto/views"
)

func (u service) AddView(ctx context.Context, ip string, req *dto.AddViewRequest) (*dto.AddViewResponse, error) {
	_, err := u.repo.CreateView(ctx, ip, req.Origin)
	if err != nil {
		log.Println("error creating view", err)
		return nil, ErrInternal
	}

	return &dto.AddViewResponse{}, nil
}

func (u service) GetViewCount(ctx context.Context) (*dto.GetViewCountResponse, error) {
	count, err := u.repo.CountViews(ctx)
	if err != nil {
		log.Println("error counting views", err)
		return nil, ErrInternal
	}

	return &dto.GetViewCountResponse{Views: count}, nil
}
