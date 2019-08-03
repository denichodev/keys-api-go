package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

// SuccessResponse is the baseline of all succeed response
type SuccessResponse struct {
	HTTPStatusCode int         `json:"-"`
	Status         string      `json:"status"`
	Data           interface{} `json:"data"`
}

// Render is chi renderer for basic success response
func (res *SuccessResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, res.HTTPStatusCode)
	res.Status = "success"

	return nil
}
