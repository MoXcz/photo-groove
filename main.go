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

func (app *application) folders(w http.ResponseWriter, r *http.Request) {
	folders := map[string]any{
		"name":   "All Photos",
		"photos": map[string]any{},
		"subfolders": []any{
			map[string]any{
				"name":   "2016",
				"photos": map[string]any{},
				"subfolders": []any{
					map[string]any{
						"name": "Aloha",
						"photos": map[string]any{
							"2turtles": map[string]any{
								"title": "Turtles & sandals",
								"related_photos": []string{
									"beach",
								},
								"size": 27,
							},
							"beach": map[string]any{
								"title": "At Chang’s Beach",
								"related_photos": []string{
									"wake",
									"2turtles",
								},
								"size": 36,
							},
							"wake": map[string]any{
								"title": "First day on Maui",
								"related_photos": []string{
									"beach",
								},
								"size": 21,
							},
						},
						"subfolders": []any{},
					},
					map[string]any{
						"name": "Metal",
						"photos": map[string]any{
							"epica": map[string]any{
								"title": "Simone Simons",
								"related_photos": []string{
									"Noora",
									"Joakim",
								},
								"size": 20,
							},
							"Noora": map[string]any{
								"title": "Noora Louhimo",
								"related_photos": []string{
									"epica",
									"Joakim",
								},
								"size": 29,
							},
							"Joakim": map[string]any{
								"title": "Joakim Brodén",
								"related_photos": []string{
									"Noora",
									"epica",
								},
								"size": 29,
							},
						},
						"subfolders": []any{},
					},
				},
			},
			map[string]any{
				"name":       "2017",
				"photos":     map[string]any{},
				"subfolders": []any{},
			},
		},
	}

	app.respJSON(w, http.StatusOK, folders)
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
	mux.HandleFunc("GET /folders/list", app.folders)

	return mux
}
