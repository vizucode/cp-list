package models

import "time"

type Article struct {
	Id        int
	Title     string
	Content   string
	Category  string
	Status    string
	UpdatedAt time.Time
	CreatedAt time.Time
}
