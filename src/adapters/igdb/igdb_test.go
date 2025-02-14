package igdb

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGameData(t *testing.T) {
	// Setup
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("--- handle HTTP mock ---")

		if r.URL.Path != "/v4/games" {
			t.Errorf("Expected to request '/v4/games', got: %s", r.URL.Path)
		}
		// if r.Header.Get("Accept") != "application/json" {
		// 	t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		// }

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"value":"fixed"}`))
	}))
	testServer.Config.Addr = "localhost:3000"
	defer testServer.Close()

	igdbAdapter := NewIGDBAdapter()
	gameID := 1942

	// Execution
	gameData := igdbAdapter.GetGameData(gameID)

	// Validation
	if igdbAdapter.AuthToken != "authToken" {
		t.Errorf("Expected auth token to be authToken, but got %s", igdbAdapter.AuthToken)
	}

	if gameData.Name != "1942" {
		t.Errorf("Expected game name to be 1942, but got %s", gameData.Name)
	}

	// assert.Nil(t, err)
	// assert.Equal(t, "1942", game.Name)
	// assert.Equal(t, "1984-12-01", game.FirstReleaseDate)
	// assert.Equal(t, "1942", game.Slug)
}
