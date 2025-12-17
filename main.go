package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{logger: logger}

	srv := http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	app.logger.Info("Starting server on :8080")
	err := srv.ListenAndServe()
	if err != nil {
		logger.Error("Server failed to start", "error", err)
		return
	}
}

type Photos struct {
	List []string `json:"photos"`
}

func (app *application) photos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "1.jpeg,2.jpeg,3.jpeg,4.jpeg")
}

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /", fs)

	mux.HandleFunc("GET /photos/list", app.photos)

	return mux
}
