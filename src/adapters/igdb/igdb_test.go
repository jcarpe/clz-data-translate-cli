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

func TestSearchGameName(t *testing.T) {
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
	searchTerm := "tokobot"

	// Execution
	gamesData := igdbAdapter.SearchGameByTerm(searchTerm)

	// Assertion
	if gamesData == nil {
		t.Errorf("Expected game data to be returned, but got nil")
	}

	if gamesData[1].Name != "Tokobot Plus: Mysteries of the Karakuri" {
		t.Errorf("Expected game name to be Tokobot Plus: Mysteries of the Karakuri, but got %s", gamesData[1].Name)
	}
}
