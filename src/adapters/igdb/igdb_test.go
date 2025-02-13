package igdb

import (
	"testing"
)

func TestGetGameData(t *testing.T) {
	// Setup
	igdbAdapter := NewIGDBAdapter()
	gameID := 1942

	// Execution
	gameData := igdbAdapter.GetGameData(gameID)

	// Validation
	if gameData.Name != "1942" {
		t.Errorf("Expected game name to be 1942, but got %s", gameData.Name)
	}

	// assert.Nil(t, err)
	// assert.Equal(t, "1942", game.Name)
	// assert.Equal(t, "1984-12-01", game.FirstReleaseDate)
	// assert.Equal(t, "1942", game.Slug)
}
