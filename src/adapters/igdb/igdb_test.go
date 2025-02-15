package igdb

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestIGDBServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "access12345token",
			"expires_in":   5587808,
			"token_type":   "bearer",
		})
	}))
}

func TestGetGameData(t *testing.T) {
	testServer := getTestIGDBServer()
	defer testServer.Close()

	igdbAdapter := NewIGDBAdapter(IGDBAdapterInit{
		AuthBaseUrl:      testServer.URL,
		AuthUrlPath:      "/oauth2/token",
		AuthClientId:     "client12345id",
		AuthClientSecret: "client123secret",
	})
	gameID := 1942

	// Execution
	gameData := igdbAdapter.GetGameData(gameID)

	// Validation
	if igdbAdapter.AuthToken != "access12345token" {
		t.Errorf("Expected auth token to be authToken, but got %s", igdbAdapter.AuthToken)
	}

	if gameData.Name != "1942" {
		t.Errorf("Expected game name to be 1942, but got %s", gameData.Name)
	}
}
