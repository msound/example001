package app

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (app *App) randomBookHandler(w http.ResponseWriter, r *http.Request) {
	slog.Debug("randonBookHandler")
	book, err := app.booker.GetRandomBook(r.Context())
	if err != nil {
		slog.Error("Error getting random book", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		slog.Error("Error encoding json")
	}
}
