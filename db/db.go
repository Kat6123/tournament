package db

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/kat6123/tournament/model"
)

var session *mgo.Session

func Connect() {
	s, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(fmt.Sprintf("dial with db has failed: %v", err))
	}
	session = s
}

func Close() {
	session.Close()
}

func LoadPlayer(playerId int) (*model.Player, error) {
	s := session.Copy()
	defer s.Close()

	c := session.DB("tournament").C("players")
	var p model.Player

	err := c.FindId(playerId).One(&p)
	if err != nil {
		return nil, fmt.Errorf("load player with id %d from db: %v", playerId, err)
	}

	return &p, nil
}

func SavePlayer(p *model.Player) error {
	s := session.Copy()
	defer s.Close()

	c := session.DB("tournament").C("players")
	if err := c.UpdateId(p.ID, &p); err != nil {
		return fmt.Errorf("save player %d in db: %v", p.ID, err)
	}
	return nil
}

func LoadTournament(tourId int) (*model.Tournament, error) {
	s := session.Copy()
	defer s.Close()

	c := session.DB("tournament").C("tours")
	var t model.Tournament

	err := c.FindId(tourId).One(&t)
	if err != nil {
		return nil, fmt.Errorf("load tournament with id %d from db: %v", tourId, err)
	}

	return &t, nil
}

func SaveTournament(t *model.Tournament) error {
	s := session.Copy()
	defer s.Close()

	c := session.DB("tournament").C("tours")
	if err := c.UpdateId(t.ID, &t); err != nil {
		return fmt.Errorf("save tour %d in db: %v", t.ID, err)
	}
	return nil
}

func CreateTournament(t *model.Tournament) error {
	s := session.Copy()
	defer s.Close()

	c := session.DB("tournament").C("tours")
	if err := c.Insert(t); err != nil {
		return fmt.Errorf("insert tour %s in db: %v", t, err)
	}

	return nil
}
