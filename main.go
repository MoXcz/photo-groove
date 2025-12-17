package main

import (
	"encoding/json"
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
		os.Exit(1)
	}
}

type Photo struct {
	Url   string `json:"url"`
	Size  int    `json:"size"`
	Title string `json:"title,omitempty"`
}

func (app *application) photos(w http.ResponseWriter, r *http.Request) {
	photos := []Photo{
		{Url: "1.jpeg", Size: 36, Title: "Beachside"},
		{Url: "2.jpeg", Size: 19, Title: "Epica, live at the Agora"},
		{Url: "3.jpeg", Size: 41},
		{Url: "4.jpeg", Size: 41, Title: "City Museum"},
		{Url: "5.jpeg", Size: 25, Title: ""},
		{Url: "6.jpeg", Size: 37, Title: "Boat in Glass"},
	}

	app.respJSON(w, http.StatusOK, photos)
}

func (app *application) respJSON(w http.ResponseWriter, code int, jsonResp any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, err := json.Marshal(jsonResp)
	if err != nil {
		app.logger.Error("Error marshalling JSON", "err", err)
	}
	w.Write(data)
}

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /", fs)

	mux.HandleFunc("GET /photos/list", app.photos)

	return mux
}
