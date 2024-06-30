package service

import (
	"context"
	"log/slog"
	"math/rand"

	"github.com/msound/example001/pkg/model"
)

func (s *Service) GetRandomBook(ctx context.Context) (*model.Book, error) {
	slog.Debug("Getting random book")
	books, err := s.storer.GetBooks(ctx)
	if err != nil {
		return nil, err
	}

	slog.Debug("Got books", "count", len(books))

	n := rand.Intn(len(books))

	slog.Debug("Selected random book", "rand", n)

	return &books[n], nil
}
