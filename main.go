package main

import (
	"log"
	"net/http"

	"github.com/kat6123/tournament/db"
	"github.com/kat6123/tournament/handler"
	"github.com/kat6123/tournament/logic"
)

func main() {
	pc, tc, err := db.New(":27017", "tours")
	if err != nil {
		log.Fatalf("dial with db has failed: %v", err)
	}

	service := logic.New(logic.Builder{PP: pc, TP: tc})
	api := handler.New(service)

	log.Fatal(http.ListenAndServe(":3001", api.Router()))
}
