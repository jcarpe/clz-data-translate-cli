package igdb

import (
	"main/src/adapters/_test/mocks"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Setup code here (if needed)
	testAuthServer := mocks.GetTestTwitchAuthServer()
	defer testAuthServer.Close()

	testIGDBServer := mocks.GetTestIGDBServer()
	defer testIGDBServer.Close()

	// Set environment variables for IGDB
	os.Setenv("IGDB_AUTH_BASE_URL", testAuthServer.URL)
	os.Setenv("IGDB_AUTH_PATH", "/oauth2/token")
	os.Setenv("IGDB_CLIENT_ID", "test_client_id")
	os.Setenv("IGDB_CLIENT_SECRET", "test_client_secret")
	os.Setenv("IGDB_BASE_URL", testIGDBServer.URL)

	// Run the tests
	exitCode := m.Run()

	// Teardown code here (if needed)

	// Exit with the appropriate code
	os.Exit(exitCode)
}

func TestGetPlatformData(t *testing.T) {
	igdbAdapter := NewIGDBAdapter(IGDBAdapterInit{
		AuthBaseUrl:      os.Getenv("IGDB_AUTH_BASE_URL"),
		AuthUrlPath:      os.Getenv("IGDB_AUTH_PATH"),
		AuthClientId:     os.Getenv("IGDB_CLIENT_ID"),
		AuthClientSecret: os.Getenv("IGDB_CLIENT_SECRET"),
		IGDBBaseUrl:      os.Getenv("IGDB_BASE_URL"),
	})

	// Execution
	platformData := igdbAdapter.GetPlatformData()

	if len(platformData) == 0 {
		t.Errorf("Expected platform data to be returned, but got nil")
		t.FailNow()
	}

	if platformData[0].Name != "Nintendo Entertainment System" {
		t.Errorf("Expected platform name to be Nintendo Entertainment System, but got %s", platformData[0].Name)
	}
}

func TestGetGameData(t *testing.T) {
	igdbAdapter := NewIGDBAdapter(IGDBAdapterInit{
		AuthBaseUrl:      os.Getenv("IGDB_AUTH_BASE_URL"),
		AuthUrlPath:      os.Getenv("IGDB_AUTH_PATH"),
		AuthClientId:     os.Getenv("IGDB_CLIENT_ID"),
		AuthClientSecret: os.Getenv("IGDB_CLIENT_SECRET"),
		IGDBBaseUrl:      os.Getenv("IGDB_BASE_URL"),
	})
	gameID := 1068 // <-- Super Mario Bros 3 ID value in IGDB

	// Execution
	gameData := igdbAdapter.GetGameData(gameID)

	if gameData.Name != "Super Mario Bros. 3" {
		t.Errorf("Expected game name to be Super Mario Bros. 3, but got %s", gameData.Name)
	}
}

func TestSearchGameName(t *testing.T) {
	igdbAdapter := NewIGDBAdapter(IGDBAdapterInit{
		AuthBaseUrl:      os.Getenv("IGDB_AUTH_BASE_URL"),
		AuthUrlPath:      os.Getenv("IGDB_AUTH_PATH"),
		AuthClientId:     os.Getenv("IGDB_CLIENT_ID"),
		AuthClientSecret: os.Getenv("IGDB_CLIENT_SECRET"),
		IGDBBaseUrl:      os.Getenv("IGDB_BASE_URL"),
	})
	searchTerm := "tokobot"

	// Execution
	gamesData := igdbAdapter.SearchGameByTerm(searchTerm)

	// Assertion
	if gamesData == nil {
		t.Errorf("Expected game data to be returned, but got nil")
	}

	if len(gamesData) != 3 {
		t.Errorf("Expected 3 games to be returned, but got %d", len(gamesData))
	}

	if gamesData[1].Name != "Tokobot Plus: Mysteries of the Karakuri" {
		t.Errorf("Expected game name to be Tokobot Plus: Mysteries of the Karakuri, but got %s", gamesData[1].Name)
	}

	if gamesData[0].Name != "Tokobot" {
		t.Errorf("Expected game name to be Tokobot, but got %s", gamesData[0].Name)
	}
}
