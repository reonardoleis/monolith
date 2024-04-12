package views_domain

import "time"

type View struct {
	ID        int
	IP        string
	Origin    string
	Visits    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
