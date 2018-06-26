package db

import (
	"github.com/globalsign/mgo"
	"github.com/kat6123/tournament/api"
	"github.com/kat6123/tournament/model"
)

type mgoSession struct {
	url     string
	session *mgo.Session
}

// Builder type should be used to init connection to MongoDB.
type Builder struct {
	URL string
}

// New returns value of api.Repository interface type which was tuned with Builder.
func New(b Builder) api.Repository {
	return &mgoSession{
		url: b.URL,
	}
}

func (m *mgoSession) Connect() (err error) {
	m.session, err = mgo.Dial(m.url)
	return err
}

func (m *mgoSession) Close() {
	// Close doesn't return error ?
	m.session.Close()
}

func (m *mgoSession) LoadPlayer(playerID int) (*model.Player, error) {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB("tournament").C("players")
	var p model.Player

	err := c.FindId(playerID).One(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (m *mgoSession) SavePlayer(p *model.Player) error {
	s := m.session.Copy()
	defer s.Close()

	c := s.DB("tournament").C("players")

	return c.UpdateId(p.ID, &p)
}

func (m *mgoSession) LoadTournament(tourID int) (*model.Tournament, error) {
	s := m.session.Copy()
	defer s.Close()

	c := s.DB("tournament").C("tours")
	var t model.Tournament

	err := c.FindId(tourID).One(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (m *mgoSession) SaveTournament(t *model.Tournament) error {
	s := m.session.Copy()
	defer s.Close()

	c := s.DB("tournament").C("tours")

	return c.UpdateId(t.ID, &t)
}

func (m *mgoSession) CreateTournament(t *model.Tournament) error {
	s := m.session.Copy()
	defer s.Close()

	c := s.DB("tournament").C("tours")

	return c.Insert(t)
}
