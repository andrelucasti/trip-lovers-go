package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/trip/{userId}/new", NewTrip)
	r.HandleFunc("/trip/{userId}", ListAllTripsByUserId)

	r.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Home Page Health")
	})

	log.Println("Server started - listening in 8181 port")
	err := http.ListenAndServe(":8181", r)

	if err != nil {
		log.Panic(err)
	}
}
