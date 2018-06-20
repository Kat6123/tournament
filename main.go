package main

import (
	"log"
	"net/http"

	"github.com/kat6123/tournament/db"
	"github.com/kat6123/tournament/routes"
)

func main() {
	db.Connect()
	defer db.Close()

	routes.Set()

	if err := http.ListenAndServe(":8080", routes.Router); err != nil {
		log.Fatal(err)
	}
}
