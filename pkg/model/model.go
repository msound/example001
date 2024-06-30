package model

import "context"

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Booker interface {
	GetRandomBook(ctx context.Context) (*Book, error)
}

type Storer interface {
	GetBooks(ctx context.Context) ([]Book, error)
}
