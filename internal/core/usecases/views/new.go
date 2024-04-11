package views_usecase

import domain "github.com/reonardoleis/views/internal/core/domain/views"

type usecase struct {
	repo domain.ViewRepository
}

func New(repo domain.ViewRepository) domain.ViewUsecase {
	return &usecase{
		repo: repo,
	}
}
