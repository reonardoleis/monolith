package di

import (
	views_repository "github.com/reonardoleis/views/internal/adapters/database/postgres/repository/views"
	"github.com/reonardoleis/views/internal/adapters/transport/http"
	views_handlers "github.com/reonardoleis/views/internal/adapters/transport/http/handlers/views"
	views_service "github.com/reonardoleis/views/internal/core/services/views"
)

func RestAPI() *http.Server {
	viewsRepository := views_repository.New()
	viewsUsecase := views_service.New(viewsRepository)
	viewsHandlers := views_handlers.New(viewsUsecase)

	return http.NewServer(viewsHandlers)
}
