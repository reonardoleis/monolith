package postgres

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var Pool *pgxpool.Pool

func Connect() error {
	connString := os.Getenv("POSTGRES_CONN_STRING")

	var err error

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Println("error parsing connection string", err)
		return err
	}

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Println("error connecting to database", err)
		return err
	}

	return nil
}
