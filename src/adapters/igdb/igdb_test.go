package igdb

import (
	"testing"
)

func TestGetGameData(t *testing.T) {
	// Setup
	igdbAdapter := NewIGDBAdapter()
	gameID := 1942

	// Execution
	igdbAdapter.GetGameData(gameID)

	// Validation
	// assert.Nil(t, err)
	// assert.Equal(t, "1942", game.Name)
	// assert.Equal(t, "1984-12-01", game.FirstReleaseDate)
	// assert.Equal(t, "1942", game.Slug)
}
