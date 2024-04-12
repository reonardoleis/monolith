package views_repository

import domain "github.com/reonardoleis/views/internal/core/ports/views"

type Repository struct{}

func New() domain.ViewRepository {
	return &Repository{}
}
