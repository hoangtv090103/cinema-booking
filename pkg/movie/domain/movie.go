package domain

import "time"

type Movie struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PosterPath  string    `json:"poster_path"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    int       `json:"duration"` //Duration in minutes
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Active      bool      `json:"active"`
}

type MovieCreate struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PosterPath  string    `json:"poster_path"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    int       `json:"duration"` //Duration in minutes
}
