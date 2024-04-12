package main

import (
	"log"

	"github.com/reonardoleis/views/internal/adapters/database/postgres"
	"github.com/reonardoleis/views/internal/di"
)

func main() {
	err := postgres.Connect()
	if err != nil {
		log.Fatalln("error connecting to database:", err)
	}

	restAPI := di.RestAPI()
	if err := restAPI.Run(); err != nil {
		log.Fatalln("error running server:", err)
	}
}
