package app

import (
	"net/http"

	"github.com/msound/example001/pkg/model"
)

type App struct {
	booker model.Booker
}

func NewApp(booker model.Booker) *App {
	app := App{booker: booker}

	return &app
}

func (app *App) GetRouter() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("GET /", app.indexHandler)
	r.HandleFunc("GET /random-book", app.randomBookHandler)

	return r
}
