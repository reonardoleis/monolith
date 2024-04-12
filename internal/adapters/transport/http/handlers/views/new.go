package views_handlers

import domain "github.com/reonardoleis/views/internal/core/ports/views"

type ViewsHandler struct {
	svc domain.ViewService
}

func New(svc domain.ViewService) *ViewsHandler {
	return &ViewsHandler{svc}
}
