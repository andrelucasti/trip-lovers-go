package main

import (
	"github.com/andrelucasti/trip-lovers-go/app"
	"log"
)

func main() {

	log.Println("Starting server")
	app.HandleRequest()
}
