package db

import (
	"testing"

	"github.com/kat6123/tournament/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChangeTour(t *testing.T) {
	_, tc, err := New(URL, tours, players)
	require.NoError(t, err, "dial with db has failed: %v", err)

	defer func() {
		err = tc.delete(1)
		require.NoError(t, err, "can't delete test tour: %v", err)
	}()

	tour := model.NewTour(1, 300)

	tc.Create(tour)
	require.NoError(t, err, "unable to create a tour %s: %v", tour, err)

	tour.Deposit += 100
	tc.Save(tour)
	require.NoError(t, err, "unable to save a tour %s: %v", tour, err)

	var gotT *model.Tournament
	gotT, err = tc.ByID(tour.ID)
	require.NoError(t, err, "unable to save a tour %s: %v", tour, err)

	assert.Equal(t, tour, gotT)
}
