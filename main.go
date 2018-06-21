package main

import (
	"log"
	"net/http"

	"github.com/kat6123/tournament/db"
	"github.com/kat6123/tournament/route"
)

func main() {
	db.Connect()
	defer db.Close()

	log.Fatal(
		http.ListenAndServe(":8080", route.Serve()))
}
