package infrastructure

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sh0e1/wire/interface/handler"
)

func Route(sqlHandler *SQLHandler) *chi.Mux {
	todoHandler := handler.InitializeTodoHandler(sqlHandler)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/todos", func(r chi.Router) {
		r.Post("/", todoHandler.Create)
		r.Get("/", todoHandler.List)
		r.Get("/{id}", todoHandler.Get)
	})

	return r
}
