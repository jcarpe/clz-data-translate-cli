package clz_translate

import (
	"encoding/json"
	"io"
	"main/src/domain"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

var (
	testGetGameDataResponse = []map[string]interface{}{
		{
			"artworks":           []int{358989},
			"cover":              358989,
			"first_release_date": 593568000,
			"franchise":          24,
			"game_status":        0,
			"game_type":          0,
			"genres":             []int{8},
			"id":                 1068,
			"name":               "Super Mario Bros. 3",
			"platforms":          []int{52, 37, 5, 99, 41, 18},
			"storyline":          "The Mushroom Kingdom has been a peaceful place thanks to the brave deeds of Mario and Luigi. The Mushroom Kingdom forms an entrance to the Mushroom World where all is not well. Bowser has sent his 7 children to make mischief as they please in the normally peaceful Mushroom World. They stole the royal magic wands from each country in the Mushroom World and used them to turn their kings into animals. Mario and Luigi must recover the royal magic wands from Bowser's 7 kids to return the kings to their true forms. \"Goodbye and good luck!,\" said the Princess and Toad as Mario and Luigi set off on their journey deep into the Mushroom World.",
			"summary":            "Super Mario Bros. 3, the third entry in the Super Mario Bros. series and Super Mario franchise, sees Mario or Luigi navigate a nonlinear world map containing platforming levels and optional minigames and challenges. The game features more diverse movement options and new items alongside more complex level designs and boss battles.",
			"videos":             []int{35343, 20256},
		},
	}

	testSearchGameNameResponse = []map[string]interface{}{
		{
			"artworks":           nil,
			"cover":              90786,
			"first_release_date": 1136419200,
			"franchise":          0,
			"game_status":        0,
			"game_type":          0,
			"genres":             []int{8},
			"id":                 19692,
			"name":               "Tokobot",
			"platforms":          []int{46, 38},
			"storyline":          "",
			"summary":            "Players take on the role of the young hero Bolt, a quick thinking agent who has discovered some friendly, highly advanced robots called \"Tokobots\" during his explorations of ancient ruins. With the help of the loyal Tokobots, Bolt will reveal mysteries and save the world from a horrible plot, as the Tokobots faithfully follow him on his journey, helping him to avoid obstacles, traps and enemies, by working together to create \"team combos\". During these actions, the Tokobots team up to take on different combinations in order to simulate everything from a ladder that Bolt climbs to wings that allow him to fly over large obstacles. The player will have to use strategy and skill to create these team combos in order to complete each level and succeed in the game.",
			"videos":             nil,
		},
		{
			"artworks":           nil,
			"cover":              136520,
			"first_release_date": 1163721600,
			"franchise":          0,
			"game_status":        0,
			"game_type":          10,
			"genres":             []int{8},
			"id":                 20663,
			"name":               "Tokobot Plus: Mysteries of the Karakuri",
			"platforms":          []int{8},
			"storyline":          "",
			"summary":            "Bolt is an apprentice Treasure Master working at Mr. Canewood's laboratory. With the help of mysterious Tokobots, he will explore many prehistoric ruins to find ancient relics that will put him one step further on his path to be a Treasure Master and find his father. However, he will also change his destiny in other ways.",
			"videos":             nil,
		},
	}
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

func getTestIGDBServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Failed to read request body: %v", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if strings.Contains(string(body), "search") {
			json.NewEncoder(w).Encode(testSearchGameNameResponse)
		} else {
			json.NewEncoder(w).Encode(testGetGameDataResponse)
		}
	}))
}

func TestTranslateCLZ(t *testing.T) {
	// Arrange
	testAuthServer := getTestTwitchAuthServer()
	defer testAuthServer.Close()

	testIGDBServer := getTestIGDBServer(t)
	defer testIGDBServer.Close()

	// Set environment variables for IGDB
	os.Setenv("IGDB_AUTH_BASE_URL", testAuthServer.URL)
	os.Setenv("IGDB_AUTH_PATH", "/oauth2/token")
	os.Setenv("IGDB_CLIENT_ID", "test_client_id")
	os.Setenv("IGDB_CLIENT_SECRET", "test_client_secret")
	os.Setenv("IGDB_BASE_URL", testIGDBServer.URL)

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
		PricechartingValue: 0.00,
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
}
