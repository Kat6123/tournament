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

type Builder struct {
	URL string
}

func New(b Builder) api.Repository {
	return &mgoSession{
		url: b.URL,
	}
}

func (m *mgoSession) Connect() (err error) {
	m.session, err = mgo.Dial(m.url)
	if err != nil {
		return err
	}

	return nil
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
	if err := c.UpdateId(p.ID, &p); err != nil {
		return err
	}
	return nil
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
	if err := c.UpdateId(t.ID, &t); err != nil {
		return err
	}
	return nil
}

func (m *mgoSession) CreateTournament(t *model.Tournament) error {
	s := m.session.Copy()
	defer s.Close()

	c := s.DB("tournament").C("tours")
	if err := c.Insert(t); err != nil {
		return nil
	}

	return nil
}
