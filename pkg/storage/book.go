package storage

import (
	"context"

	"github.com/msound/example001/pkg/model"
)

func (stor *Storage) GetBooks(ctx context.Context) ([]model.Book, error) {
	books := make([]model.Book, 0)

	books = append(books, model.Book{Title: "Prophet Song", Author: "Paul Lynch"})
	books = append(books, model.Book{Title: "The Seven Moons of Malli Almeida", Author: "Shehan Karunatilaka"})
	books = append(books, model.Book{Title: "The Promise", Author: "Damon Galgut"})
	books = append(books, model.Book{Title: "Shuggie Bain", Author: "Douglas Stuart"})

	return books, nil
}
