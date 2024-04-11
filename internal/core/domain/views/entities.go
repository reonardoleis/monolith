package views_domain

import "time"

type View struct {
	ID        int
	IP        string
	Origin    string
	CreatedAt time.Time
}
