package views_service

import domain "github.com/reonardoleis/views/internal/core/ports/views"

type service struct {
	repo domain.ViewRepository
}

func New(repo domain.ViewRepository) domain.ViewService {
	return &service{
		repo: repo,
	}
}
