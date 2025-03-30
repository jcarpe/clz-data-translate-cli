package igdb

import (
	"main/src/adapters/_test/mocks"
	"testing"
)

func TestGetGameData(t *testing.T) {
	testAuthServer := mocks.GetTestTwitchAuthServer()
	defer testAuthServer.Close()

	testIGDBServer := mocks.GetTestIGDBServer(t)
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

func TestSearchGameName(t *testing.T) {
	testAuthServer := mocks.GetTestTwitchAuthServer()
	defer testAuthServer.Close()

	testIGDBServer := mocks.GetTestIGDBServer(t)
	defer testIGDBServer.Close()

	igdbAdapter := NewIGDBAdapter(IGDBAdapterInit{
		AuthBaseUrl:      testAuthServer.URL,
		AuthUrlPath:      "/oauth2/token",
		AuthClientId:     "clientID123",
		AuthClientSecret: "clientSecret123",
		IGDBBaseUrl:      testIGDBServer.URL,
	})
	searchTerm := "tokobot"

	// Execution
	gamesData := igdbAdapter.SearchGameByTerm(searchTerm)

	// Assertion
	if gamesData == nil {
		t.Errorf("Expected game data to be returned, but got nil")
	}

	if len(gamesData) != 2 {
		t.Errorf("Expected 2 games to be returned, but got %d", len(gamesData))
	}

	if gamesData[1].Name != "Tokobot Plus: Mysteries of the Karakuri" {
		t.Errorf("Expected game name to be Tokobot Plus: Mysteries of the Karakuri, but got %s", gamesData[1].Name)
	}

	if gamesData[0].Name != "Tokobot" {
		t.Errorf("Expected game name to be Tokobot, but got %s", gamesData[0].Name)
	}
}
