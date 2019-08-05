package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"github.com/denichodev/keys-api/keys"
	"github.com/go-chi/chi"
)

type KeysHandler struct{}

// Return a new handler
func NewKeysHandler() *KeysHandler {
	return &KeysHandler{}
}

// GetRoutes returns all the routes for the keys handler
func (h *KeysHandler) GetRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.GetKey)

	return r
}

// GetKey will return a single key generated before
func (h *KeysHandler) GetKey(w http.ResponseWriter, r *http.Request) {
	rand := keys.String(6)
	fmt.Printf("Get key: %s\n", rand)

	payload := SuccessResponse{
		HTTPStatusCode: 200,
		Data:           rand,
	}

	render.Render(w, r, &payload)
}
