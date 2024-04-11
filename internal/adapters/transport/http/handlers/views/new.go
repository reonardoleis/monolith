package views_handlers

import domain "github.com/reonardoleis/views/internal/core/domain/views"

type ViewsHandler struct {
	usecase domain.ViewUsecase
}

func New(usecase domain.ViewUsecase) *ViewsHandler {
	return &ViewsHandler{usecase}
}
