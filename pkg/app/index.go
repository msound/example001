package app

import "net/http"

func (app *App) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("example001"))
}
