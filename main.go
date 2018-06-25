package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/kat6123/tournament/db"
	"github.com/kat6123/tournament/handler"
	"github.com/kat6123/tournament/logic"
)

func main() {
	repository := db.New(db.Builder{
		URL: "127.0.0.1:27017",
	})
	if err := repository.Connect(); err != nil {
		panic(fmt.Errorf("dial with db has failed: %v", err))
	}
	defer repository.Close()

	logic.Init(logic.Builder{R: repository})

	log.Fatal(http.ListenAndServe(":3001", handler.ServeRoutes()))
}
