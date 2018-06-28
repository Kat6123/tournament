package db

import (
	"testing"

	"github.com/kat6123/tournament/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testID = 1

// Is it better to use require thane err != nil?
// Can one test run in parallel?
func TestChangeTour(t *testing.T) {
	_, tc, err := New(URL, DB)
	require.Nil(t, err, "dial with db has failed: %v", err)

	defer func() {
		err = tc.delete(testID)
		require.Nil(t, err, "can't delete test tour: %v", err)
	}()

	tour := model.NewTour(testID, 300)

	tc.Create(tour)
	require.Nil(t, err, "unable to create a tour %s: %v", tour, err)

	tour.Deposit += 100
	tc.Save(tour)
	require.Nil(t, err, "unable to save a tour %s: %v", tour, err)

	var gotT *model.Tournament
	gotT, err = tc.ByID(tour.ID)
	require.Nil(t, err, "unable to save a tour %s: %v", tour, err)

	assert.Equal(t, tour, gotT)
}
