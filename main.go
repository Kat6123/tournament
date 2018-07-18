package main

import (
	"log"
	"net/http"

	"github.com/kat6123/tournament/db"
	"github.com/kat6123/tournament/handler"
	"github.com/kat6123/tournament/logic"
)

func main() {
	//conf, err := config.Get()
	//if err != nil{
	//	log.Fatalf("parse config has failed: %v", err)
	//}
	//pc, tc, err := db.New(conf.URL, conf.DB)
	//if err != nil {
	//	log.Fatalf("dial with db has failed: %v", err)
	//}
	//
	//service := logic.New(logic.Builder{PP: pc, TP: tc})
	//api := handler.New(service)
	//
	//log.Fatal(http.ListenAndServe(":3001", api.Router()))

	pc, tc, err := db.New(":27017", "tours")
	if err != nil {
		log.Fatalf("dial with db has failed: %v", err)
	}

	service := logic.New(logic.Builder{PP: pc, TP: tc})
	api := handler.New(service)

	log.Fatal(http.ListenAndServe(":3001", api.Router()))
}
