package routing

import (
	"learn/handlers"
	"net/http"
)

var Mux = http.NewServeMux()

func InitializeRoutes() {

	Mux.HandleFunc("/players", handlers.GetPlayers)

	Mux.HandleFunc("/player/{id}", handlers.GeyPlayerById)

	Mux.HandleFunc("/add-player", handlers.AddPlayerHandler().ServeHTTP)

	Mux.HandleFunc("/remove-player/{id}", handlers.RemovePlayerHandler().ServeHTTP)

}
