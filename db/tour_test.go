package db

import (
	"testing"

	"github.com/kat6123/tournament/model"
	"github.com/stretchr/testify/assert"
)

func TestChangeTour(t *testing.T) {
	_, tc, err := New(URL, DB)
	if err != nil {
		t.Fatalf("dial with db has failed: %v", err)
	}
	defer func() {
		if err = tc.DeleteTournament(1); err != nil {
			t.Fatalf("can't delete test tour: %v", err)
		}
	}()

	tour := model.NewTour(1, 300)

	if err = tc.CreateTournament(tour); err != nil {
		t.Fatalf("unable to create a tour %s: %v", tour, err)
	}

	tour.Deposit += 100
	if err = tc.SaveTournament(tour); err != nil {
		t.Fatalf("unable to save a tour %s: %v", tour, err)
	}

	var gotT *model.Tournament
	if gotT, err = tc.LoadTournament(tour.ID); err != nil {
		t.Fatalf("unable to save a tour %s: %v", tour, err)
	}

	assert.Equal(t, tour, gotT)
}
