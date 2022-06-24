package main

import (
	"botStuff/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func routes() http.Handler{
	mux := chi.NewRouter()

	mux.Get("/bots", handlers.BotSupervisor)

	mux.Get("/bots/travelBot", handlers.TravelBot)
	mux.Get("/bots/algoliaBot", handlers.AlgoliaBot)
	mux.Get("/bots/firebaseConnector", handlers.FirebaseConnector)

	return mux
}
