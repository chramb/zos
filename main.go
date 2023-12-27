package main

import (
	"net/http"

	"github.com/chramb/zos/api/account"
	"github.com/chramb/zos/api/container"
	"github.com/chramb/zos/api/object"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Mount("/v1/{account}", account.Routes())
	r.Mount("/v1/{account}/{container}", container.Routes())
	r.Mount("/v1/{account}/{container}/{object}", object.Routes())

	http.ListenAndServe(":8080", r)
}
