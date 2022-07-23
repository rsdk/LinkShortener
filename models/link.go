package models

import "time"

type Link struct {
	ID        uint
	ShortLink string
	Url       string
	CreatedAt time.Time
}
