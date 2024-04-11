package main

import (
	"log"

	"github.com/reonardoleis/views/internal/adapters/database/postgres"
	views_repository "github.com/reonardoleis/views/internal/adapters/database/postgres/repository/views"
	"github.com/reonardoleis/views/internal/adapters/transport/http"
	views_usecase "github.com/reonardoleis/views/internal/core/usecases/views"
)

func main() {
	err := postgres.Connect()
	if err != nil {
		log.Fatalln("error connecting to database:", err)
	}

	viewsRepository := views_repository.New()
	viewsUsecase := views_usecase.New(viewsRepository)

	server := http.NewServer(viewsUsecase)
	if err := server.Run(); err != nil {
		log.Fatalln("error running server:", err)
	}
}
