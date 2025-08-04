package igdb

import (
	"fmt"
	"main/src/_test/mocks"
	"main/src/domain"
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
	os.Setenv("IGDB_AUTH_BASE_URL", "https://id.twitch.tv")
	os.Setenv("IGDB_AUTH_PATH", "/oauth2/token")
	os.Setenv("IGDB_CLIENT_ID", "ob0sa1skmi4a5vvthwyob8qzm3w4fk")
	os.Setenv("IGDB_CLIENT_SECRET", "lg210dshkx54t3m25l44039xlovvq0")
	os.Setenv("IGDB_BASE_URL", "https://api.igdb.com/v4")
	os.Setenv("IGDB_API_RATE_LIMIT", "500")

	// Run the tests
	exitCode := m.Run()

	// Teardown code here (if needed)

	// Exit with the appropriate code
	os.Exit(exitCode)
}

func TestGetGameData(t *testing.T) {
	igdbAdapter := NewIGDBAdapter(IGDBAdapterInit{
		AuthBaseUrl:      os.Getenv("IGDB_AUTH_BASE_URL"),
		AuthUrlPath:      os.Getenv("IGDB_AUTH_PATH"),
		AuthClientId:     os.Getenv("IGDB_CLIENT_ID"),
		AuthClientSecret: os.Getenv("IGDB_CLIENT_SECRET"),
		IGDBBaseUrl:      os.Getenv("IGDB_BASE_URL"),
	})
	// gameID := []int{1068} // <-- Super Mario Bros 3 ID value in IGDB
	gameID := []int{21838}

	// Execution
	gameData := igdbAdapter.GetGameData(gameID)

	fmt.Println("Game Data count:", len(gameData))
	fmt.Printf("Game Data: %+v\n", gameData[0].Cover)

	if gameData[0].Name != "Super Mario Bros. 3" {
		t.Errorf("Expected game name to be Super Mario Bros. 3, but got %s", gameData[0].Name)
	}

	multipleGameData := igdbAdapter.GetGameData([]int{1068, 1069})

	for _, game := range multipleGameData {
		t.Logf("Game ID: %d, Name: %s", game.ID, game.Name)
	}
}

func TestGameTitleNormalization(t *testing.T) {
	// Test cases
	tests := []struct {
		input    string
		expected string
	}{
		{"Super Mario Bros. 3", "super mario bros. 3"},
		{"ESPN Extreme Games (Greatest Hits)", "espn extreme games"},
	}

	for _, test := range tests {
		result := GameTitleNormalization(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestFuzzyFindIGDBGameByTitleEXAMPLE(t *testing.T) {
	igdbAdapter := NewIGDBAdapter(IGDBAdapterInit{
		AuthBaseUrl:      os.Getenv("IGDB_AUTH_BASE_URL"),
		AuthUrlPath:      os.Getenv("IGDB_AUTH_PATH"),
		AuthClientId:     os.Getenv("IGDB_CLIENT_ID"),
		AuthClientSecret: os.Getenv("IGDB_CLIENT_SECRET"),
		IGDBBaseUrl:      os.Getenv("IGDB_BASE_URL"),
	})
	title := "Guardian Heroes"
	clzPlatform := "Saturn"

	// Execution
	gameID := igdbAdapter.FuzzyFindGameByTitle(title, clzPlatform)

	// Assertion
	if gameID == 0 {
		t.Errorf("Expected game ID to be found, but got 0")
	}

	if gameID != 3 {
		t.Errorf("Expected game ID to be 3, but got %d", gameID)
	}
}

func TestFuzzyFindIGDBGameByTitle(t *testing.T) {
	igdbAdapter := NewIGDBAdapter(IGDBAdapterInit{
		AuthBaseUrl:      os.Getenv("IGDB_AUTH_BASE_URL"),
		AuthUrlPath:      os.Getenv("IGDB_AUTH_PATH"),
		AuthClientId:     os.Getenv("IGDB_CLIENT_ID"),
		AuthClientSecret: os.Getenv("IGDB_CLIENT_SECRET"),
		IGDBBaseUrl:      os.Getenv("IGDB_BASE_URL"),
	})
	title := "Super Mario Bros. 3"
	clzPlatform := "NES"

	// Execution
	gameID := igdbAdapter.FuzzyFindGameByTitle(title, clzPlatform)

	// Assertion
	if gameID == 0 {
		t.Errorf("Expected game ID to be found, but got 0")
	}

	if gameID != 3 {
		t.Errorf("Expected game ID to be 3, but got %d", gameID)
	}
}

func TestFuzzySearchList(t *testing.T) {
	igdbAdapter := NewIGDBAdapter(IGDBAdapterInit{
		AuthBaseUrl:      os.Getenv("IGDB_AUTH_BASE_URL"),
		AuthUrlPath:      os.Getenv("IGDB_AUTH_PATH"),
		AuthClientId:     os.Getenv("IGDB_CLIENT_ID"),
		AuthClientSecret: os.Getenv("IGDB_CLIENT_SECRET"),
		IGDBBaseUrl:      os.Getenv("IGDB_BASE_URL"),
	})

	mockGamesList := []domain.Game{
		{
			Title:    "Super Mario Bros. 3",
			Platform: "NES",
		},
		{
			Title:    "The Legend of Zelda",
			Platform: "NES",
		},
		{
			Title:    "Final Fantasy",
			Platform: "NES",
		},
	}

	// Execution
	fuzzySearchedGames := igdbAdapter.FuzzyFindGamesList(mockGamesList)

	// Assertion
	if len(fuzzySearchedGames) == 0 {
		t.Errorf("Expected game list to be found, but got 0")
	}

	if fuzzySearchedGames[0].IGDB_ID != 3 {
		t.Errorf("Expected game IGDB_ID to be 3, but got %d", fuzzySearchedGames[0].IGDB_ID)
	}
}
