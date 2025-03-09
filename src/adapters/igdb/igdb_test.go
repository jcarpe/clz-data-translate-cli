package igdb

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestTwitchAuthServer() *httptest.Server {
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

func getTestIGDBServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]interface{}{
			{
				"id":   1068,
				"name": "Super Mario Bros. 3",
			},
		})
	}))
}

func TestGetGameData(t *testing.T) {
	testAuthServer := getTestTwitchAuthServer()
	defer testAuthServer.Close()

	testIGDBServer := getTestIGDBServer()
	defer testIGDBServer.Close()

	igdbAdapter := NewIGDBAdapter(IGDBAdapterInit{
		AuthBaseUrl:      testAuthServer.URL,
		AuthUrlPath:      "/oauth2/token",
		AuthClientId:     "clientID123",
		AuthClientSecret: "clientSecret123",
		IGDBBaseUrl:      testIGDBServer.URL,
	})
	gameID := 1068 // <-- Super Mario Bros 3 ID value in IGDB

	// Execution
	gameData := igdbAdapter.GetGameData(gameID)

	if gameData.Name != "Super Mario Bros. 3" {
		t.Errorf("Expected game name to be Super Mario Bros. 3, but got %s", gameData.Name)
	}
}
