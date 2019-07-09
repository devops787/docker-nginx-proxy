package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"
)

func main() {
	fmt.Println("Listening on port 3000")

	// router
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		serverId := os.Getenv("SERVER_ID")
		if serverId == "" {
			serverId = "0"
		}

		response := fmt.Sprintf("Server %s", serverId)
		writer.Write([]byte(response))
	})

	// server
	server := http.Server{
		Handler: router,
		Addr: "0.0.0.0:3000",
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
