package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gookit/color"
	"github.com/gorilla/mux"
	gocache "github.com/patrickmn/go-cache"
	"github.com/PoCFrance/e/network"
)

var (
	// Create the neural network variable to use it everywhere
	neuralNetworks map[string]network.Network
	// Initializes the cache with a 5 minute lifetime
	cache = gocache.New(5*time.Minute, 5*time.Minute)
)

// Serve serves the server in the given port
func Serve(_neuralNetworks map[string]network.Network, port string) {
	// Set the current global network as a global variable
	neuralNetworks = _neuralNetworks

	// Initializes the router
	router := mux.NewRouter()
	// Serve the websocket
	router.HandleFunc("/websocket", SocketHandle)
	// Serve the API
	router.HandleFunc("/api/{locale}/train", Train).Methods("POST")

	magenta := color.FgMagenta.Render
	fmt.Printf("\nServer listening on the port %s...\n", magenta(port))

	// Serves the chat
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}

// Train is the route to re-train the neural network
func Train(w http.ResponseWriter, r *http.Request) {
	// // Checks if the token present in the headers is the right one
	// magenta := color.FgMagenta.Render
	// fmt.Printf("\nRe-training the %s..\n", magenta("neural network"))

	// for locale := range neuralNetworks {
	// 	neuralNetworks[locale] = training.CreateNeuralNetwork(locale, true)
	// }
}
