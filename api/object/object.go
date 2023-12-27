package object

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	chi.RegisterMethod("COPY")

	r.Get("/", get)
	r.Put("/", put)
	r.MethodFunc("COPY", "/", copy)
	r.Post("/", post)
	r.Head("/", head)
	r.Delete("/", delete)

	return r
}

func get(w http.ResponseWriter, r *http.Request)    {}
func put(w http.ResponseWriter, r *http.Request)    {}
func copy(w http.ResponseWriter, r *http.Request)   {}
func post(w http.ResponseWriter, r *http.Request)   {}
func head(w http.ResponseWriter, r *http.Request)   {}
func delete(w http.ResponseWriter, r *http.Request) {}
