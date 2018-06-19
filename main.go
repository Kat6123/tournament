package main

import (
	"log"
	"net/http"

	"github.com/kat6123/tournament/routes"
)

func main() {
	routes.Set()

	if err := http.ListenAndServe(":8080", routes.Router); err != nil {
		log.Fatal(err)
	}
}
