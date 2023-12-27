package account

import (
	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", get)
	r.Post("/", post)
	r.Head("/", head)
	r.Delete("/", delete)

	return r
}
