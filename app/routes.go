package app

import (
	"fmt"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/trip", NewTrip)

	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Home Page Health")
	})

	log.Println("Server started - listening in 8181 port")
	err := http.ListenAndServe(":8181", nil)

	if err != nil {
		log.Panic(err)
	}
}
