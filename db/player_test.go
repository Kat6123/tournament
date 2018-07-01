package db

import (
	"testing"

	"github.com/kat6123/tournament/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	URL = ":27017"
	DB  = "test"
)

func TestChangePlayer(t *testing.T) {
	pc, _, err := New(URL, DB)
	require.NoError(t, err, "dial with db has failed: %v", err)

	defer func() {
		err = pc.delete(1)
		require.NoError(t, err, "can't delete test tour: %v", err)
	}()

	player := &model.Player{
		ID:      1,
		Balance: 300,
	}

	pc.create(player)
	require.NoError(t, err, "unable to create a player %s: %v", player, err)

	player.Balance += 100
	pc.Save(player)
	require.NoError(t, err, "unable to save a player %s: %v", player, err)

	var gotP *model.Player
	gotP, err = pc.ByID(player.ID)
	require.NoError(t, err, "unable to save a tour %s: %v", player, err)

	assert.Equal(t, player, gotP)
}
