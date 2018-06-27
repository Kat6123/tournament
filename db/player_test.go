package db

import (
	"testing"

	"github.com/globalsign/mgo"
	"github.com/kat6123/tournament/model"
	"github.com/stretchr/testify/assert"
)

const (
	URL = ":27017"
	DB  = "test"
)

func dropCollection(pc *Players, t *testing.T) {
	if err := pc.DropCollection(); err != nil {
		t.Fatalf("can't drop collection: %v", err)
	}
}

func TestPlayers_CreatePlayer(t *testing.T) {
	pc, _, err := New(URL, DB)
	if err != nil {
		t.Fatalf("dial with db has failed: %v", err)
	}
	defer dropCollection(pc, t)

	p := &model.Player{
		ID:      1,
		Balance: 300,
	}

	if err = pc.CreatePlayer(p); err != nil {
		t.Fatalf("unable to create player %s: %v", p, err)
	}

	cursor := pc.FindId(p.ID)
	n, err := cursor.Count()
	if err != nil {
		t.Fatalf("unable to count players: %v", err)
	}
	if n != 1 {
		t.Fatalf("create more than one player")
	}

	gotP := new(model.Player)
	if err := cursor.One(gotP); err != nil {
		t.Fatalf("unnable to fetch player: %v", err)
	}

	assert.Equal(t, p, gotP)
}

func TestPlayers_LoadPlayer(t *testing.T) {
	pc, _, err := New(URL, DB)
	if err != nil {
		t.Fatalf("dial with db has failed: %v", err)
	}
	defer dropCollection(pc, t)

	p := &model.Player{
		ID:      1,
		Balance: 300,
	}

	if err = pc.CreatePlayer(p); err != nil {
		t.Fatalf("unable to create player %s: %v", p, err)
	}

	gotP, err := pc.LoadPlayer(p.ID)
	if err != nil {
		t.Fatalf("unnable to load player: %v", err)
	}

	assert.Equal(t, p, gotP)
}

func TestPlayers_SavePlayer(t *testing.T) {
	pc, _, err := New(URL, DB)
	if err != nil {
		t.Fatalf("dial with db has failed: %v", err)
	}
	defer dropCollection(pc, t)

	p := &model.Player{
		ID:      1,
		Balance: 300,
	}

	if err = pc.CreatePlayer(p); err != nil {
		t.Fatalf("unable to create player %s: %v", p, err)
	}

	p.Balance += 100
	if err = pc.SavePlayer(p); err != nil {
		t.Fatalf("unable to save player %s: %v", p, err)
	}

	gotP, err := pc.LoadPlayer(p.ID)
	if err != nil {
		t.Fatalf("unnable to load player: %v", err)
	}

	assert.Equal(t, p, gotP)
}

func TestPlayers_DeletePlayer(t *testing.T) {
	pc, _, err := New(URL, DB)
	if err != nil {
		t.Fatalf("dial with db has failed: %v", err)
	}
	defer dropCollection(pc, t)

	p := &model.Player{
		ID:      1,
		Balance: 300,
	}

	if err = pc.CreatePlayer(p); err != nil {
		t.Fatalf("unable to create player %s: %v", p, err)
	}

	if err = pc.DeletePlayer(p.ID); err != nil {
		t.Fatalf("unable to delete player %s: %v", p, err)
	}

	// Why struct should be empty?
	got, err := pc.LoadPlayer(p.ID)
	if err != nil {
		assert.Equal(t, mgo.ErrNotFound, err, "should get error not found")
	}

	assert.Equal(t, model.Player{}, *got, "load player was successful when should be deleted")
}
