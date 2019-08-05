package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// CreateRouter creates root chi router
func CreateRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/ping", Ping)

	keysHandler := NewKeysHandler()

	r.Mount("/keys", keysHandler.GetRoutes())

	return r
}

// Ping handle the healthcheck of the app
func Ping(w http.ResponseWriter, r *http.Request) {
	payload := SuccessResponse{
		HTTPStatusCode: 200,
		Data:           "pong",
		Status:         "success",
	}

	render.Render(w, r, &payload)
}
