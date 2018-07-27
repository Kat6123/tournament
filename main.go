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
	conf := config.Get()
	tourlog.SetLevel(conf.Debug)

	pc, tc, err := db.New(conf.DB.URL, conf.DB.DB,
		conf.DB.TourCollection, conf.DB.PlayerCollection)
	if err != nil {
		log.Fatalf("dial with db has failed: %v", err)
	}

	service := logic.New(logic.Builder{PP: pc, TP: tc})
	api := handler.New(service)

	log.Fatal(http.ListenAndServe(":"+conf.Port, api.Router()))
}
