package model

import "github.com/google/uuid"

type Book struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Price float64   `json:"price"`
}

func NewUUID() uuid.UUID {
	return uuid.New()
}

