package main

import (
	"log/slog"
	"net/http"

	"github.com/msound/example001/pkg/app"
	"github.com/msound/example001/pkg/service"
	"github.com/msound/example001/pkg/storage"
)

const appVersion = "v1.0.0"

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Info("Example001", "version", appVersion)

	stor := storage.NewStorage()
	svc := service.NewService(stor)
	app := app.NewApp(svc)
	slog.Info("Listening")
	slog.Error("Stopping", "err", http.ListenAndServe(":8080", app.GetRouter()))
}
