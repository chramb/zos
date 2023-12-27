package container

import (
	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", get)
	r.Put("/", put)
	r.Post("/", post)
	r.Head("/", head)
	r.Delete("/", delete)

	return r
}
