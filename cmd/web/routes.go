package main

import (
	"botStuff/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func routes() http.Handler{
	mux := chi.NewRouter()

	mux.Get("/bots/123", handlers.TravelBot)
	mux.Get("/bots/789", handlers.AlgoliaBot)

	return mux
}
