package clz_translate

import (
	"main/src/adapters/_test/mocks"
	"main/src/domain"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Setup code here if needed
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

func TestTranslateCLZ(t *testing.T) {
	data, err := os.ReadFile("../_test/data/game-data-list.xml")
	if err != nil {
		t.Errorf("error reading test data: %v", err)
	}

	input := string(data)
	expectedOutput := domain.Game{
		Boxset: false,
		Completeness: domain.Completeness{
			HasBox:    false,
			HasManual: false,
			HasGame:   true,
		},
		Condition:    "",
		DateAcquired: time.Time{},
		Developers:   []string{"Sony Interactive Studios America"},
		Edition:      "Greatest Hits",
		Format:       "CD-ROM",
		Genres:       []string{"Racing", "Sports"},
		HardwareType: "Game",
		Links: []domain.Link{
			{
				Description: "1Xtreme at Core for Games",
				URL:         "http://core.collectorz.com/games/ps1/1xtreme",
			},
			{
				Description: "1Xtreme at PriceCharting.com",
				URL:         "https://www.pricecharting.com/game/Playstation/1Xtreme",
			},
		},
		Multiplayer:        false,
		Platform:           domain.PlayStation,
		PricechartingValue: 5.97,
		Publishers:         []string{"Sony Computer Entertainment America", "And Another One"},
		Quantity:           1,
		Region:             "",
		ReleaseDate:        time.Time{},
		Series:             "",
		Title:              "1Xtreme (Greatest Hits)",
	}

	actualOutput := TranslateCLZ(input, false)

	if len(actualOutput.Games) != 8 {
		t.Errorf("expected 8 games, got %d", len(actualOutput.Games))
	}

	if !reflect.DeepEqual(actualOutput.Games[0], expectedOutput) {
		t.Errorf("\nexpected \n%#v,\ngot \n%#v", expectedOutput, actualOutput.Games[0])
	}

	actualOutputWithIGDBSupplement := TranslateCLZ(input, true)

	if len(actualOutputWithIGDBSupplement.Games) != 8 {
		t.Errorf("expected 8 games, got %d", len(actualOutputWithIGDBSupplement.Games))
	}

	if actualOutputWithIGDBSupplement.Games[0].Cover.ID != 136520 {
		t.Errorf("expected cover ID to be 136520, got %d", actualOutputWithIGDBSupplement.Games[0].Cover.ID)
	}

	if actualOutputWithIGDBSupplement.Games[0].Storyline != "Gear up for some incredible Xtreme racing" {
		t.Errorf("expected storyline to be 'Gear up for some incredible Xtreme racing', got '%s'", actualOutputWithIGDBSupplement.Games[0].Storyline)
	}

	if actualOutputWithIGDBSupplement.Games[0].Summary != "Extreme racing." {
		t.Errorf("expected summary to be 'Extreme racing.', got '%s'", actualOutputWithIGDBSupplement.Games[0].Summary)
	}
}

// func TestTranslateCLZWithIGDBSupplement(t *testing.T) {
// 	data, err := os.ReadFile("../_test/data/game-data-list.xml")
// 	if err != nil {
// 		t.Errorf("error reading test data: %v", err)
// 	}

// 	input := string(data)

// 	expectedOutput := domain.Game{
// 		Boxset: false,
// 		Completeness: domain.Completeness{
// 			HasBox:    false,
// 			HasManual: false,
// 			HasGame:   true,
// 		},
// 		Condition:    "",
// 		DateAcquired: time.Time{},
// 		Developers:   []string{"Sony Interactive Studios America"},
// 		Edition:      "Greatest Hits",
// 		Format:       "CD-ROM",
// 		Genres:       []string{"Racing", "Sports"},
// 		HardwareType: "Game",
// 		Links: []domain.Link{
// 			{
// 				Description: "1Xtreme at Core for Games",
// 				URL:         "http://core.collectorz.com/games/ps1/1xtreme",
// 			},
// 			{
// 				Description: "1Xtreme at PriceCharting.com",
// 				URL:         "https://www.pricecharting.com/game/Playstation/1Xtreme",
// 			},
// 		},
// 		Multiplayer:        false,
// 		Platform:           domain.PlayStation,
// 		PricechartingValue: 5.97,
// 		Publishers:         []string{"Sony Computer Entertainment America", "And Another One"},
// 		Quantity:           1,
// }
