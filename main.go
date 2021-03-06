package main

import (
	"log"
	"net/http"

	"github.com/kat6123/tournament/config"
	"github.com/kat6123/tournament/db"
	"github.com/kat6123/tournament/handler"
	tourlog "github.com/kat6123/tournament/log"
	"github.com/kat6123/tournament/logic"
)

func main() {
	// TODO find out config libs and apply
	conf := config.Get()
	// TODO find out log libs and apply
	tourlog.SetLevel(conf.Debug)

	// TODO MySQL?
	pc, tc, err := db.New(conf.DB.URI, conf.DB.TourCollection, conf.DB.PlayerCollection)
	if err != nil {
		log.Fatalf("dial with db has failed: %v", err)
	}

	service := logic.New(logic.Builder{PP: pc, TP: tc})
	api := handler.New(service)

	log.Fatal(http.ListenAndServe(":"+conf.Port, api.Router()))
}
