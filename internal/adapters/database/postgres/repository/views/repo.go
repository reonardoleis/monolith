package views_repository

import (
	"context"
	"log"
	"time"

	"github.com/reonardoleis/views/internal/adapters/database/postgres"
	domain "github.com/reonardoleis/views/internal/core/domain/views"
)

func (r Repository) CreateView(ctx context.Context, ip, origin string) (*domain.View, error) {
	conn, err := postgres.Pool.Acquire(ctx)
	if err != nil {
		log.Println("error acquiring connection", err)
		return nil, err
	}

	defer conn.Release()

	query := `INSERT INTO views (origin, ip, created_at) 
			  VALUES ($1, $2, $3) 
			  ON CONFLICT (ip) DO UPDATE
			  SET created_at = $3,
			  	  visits = views.visits + 1
			  RETURNING id, ip, origin, created_at, visits`
	result := conn.QueryRow(ctx, query, origin, ip, time.Now())

	view := domain.View{}
	err = result.Scan(&view.ID, &view.IP, &view.Origin, &view.CreatedAt, &view.Visits)
	if err != nil {
		log.Println("error scanning result", err)
		return nil, err
	}

	return &view, nil
}

func (r Repository) CountViews(ctx context.Context) (int, error) {
	conn, err := postgres.Pool.Acquire(ctx)
	if err != nil {
		log.Println("error acquiring connection", err)
		return 0, err
	}

	defer conn.Release()

	query := `SELECT SUM(visits) FROM views`
	var count int
	err = conn.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		log.Println("error scanning result", err)
		return 0, err
	}

	return count, nil
}
