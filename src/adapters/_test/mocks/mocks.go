package mocks

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

var (
	testGetPlatformData = []map[string]interface{}{
		{
			"checksum":   "a1b2c3d4e5f6g7h8i9j0",
			"created_at": 1234567890,
			"name":       "Nintendo Entertainment System",
			"updated_at": 1234567890,
		},
	}

	testGetGameDataResponse = []map[string]interface{}{
		{
			"artworks": []int{358989},
			"cover": map[string]interface{}{
				"id":    136520,
				"width": 1000,
				"url":   "//images.igdb.com/igdb/image/upload/t_cover_big/co1j8f.jpg",
			},
			"first_release_date": 593568000,
			"franchise":          24,
			"game_status":        0,
			"game_type":          0,
			"genres":             []int{8},
			"id":                 1068,
			"name":               "Super Mario Bros. 3",
			"platforms": []map[string]interface{}{
				{
					"id":   6,
					"name": "NES",
				},
			},
			"storyline": "The Mushroom Kingdom has been a peaceful place thanks to the brave deeds of Mario and Luigi. The Mushroom Kingdom forms an entrance to the Mushroom World where all is not well. Bowser has sent his 7 children to make mischief as they please in the normally peaceful Mushroom World. They stole the royal magic wands from each country in the Mushroom World and used them to turn their kings into animals. Mario and Luigi must recover the royal magic wands from Bowser's 7 kids to return the kings to their true forms. \"Goodbye and good luck!,\" said the Princess and Toad as Mario and Luigi set off on their journey deep into the Mushroom World.",
			"summary":   "Super Mario Bros. 3, the third entry in the Super Mario Bros. series and Super Mario franchise, sees Mario or Luigi navigate a nonlinear world map containing platforming levels and optional minigames and challenges. The game features more diverse movement options and new items alongside more complex level designs and boss battles.",
			"videos":    []int{35343, 20256},
		},
	}

	testSearchGameNameResponse = []map[string]interface{}{
		{
			"artworks": nil,
			"cover": map[string]interface{}{
				"id":    136520,
				"width": 1000,
				"url":   "//images.igdb.com/igdb/image/upload/t_cover_big/co1j8f.jpg",
			},
			"first_release_date": 1136419200,
			"franchise":          0,
			"game_status":        0,
			"game_type":          0,
			"genres":             []int{8},
			"id":                 19692,
			"name":               "Tokobot",
			"platforms": []map[string]interface{}{
				{
					"id":   6,
					"name": "NES",
				},
			},
			"storyline": "",
			"summary":   "Players take on the role of the young hero Bolt, a quick thinking agent who has discovered some friendly, highly advanced robots called \"Tokobots\" during his explorations of ancient ruins. With the help of the loyal Tokobots, Bolt will reveal mysteries and save the world from a horrible plot, as the Tokobots faithfully follow him on his journey, helping him to avoid obstacles, traps and enemies, by working together to create \"team combos\". During these actions, the Tokobots team up to take on different combinations in order to simulate everything from a ladder that Bolt climbs to wings that allow him to fly over large obstacles. The player will have to use strategy and skill to create these team combos in order to complete each level and succeed in the game.",
			"videos":    nil,
		},
		{
			"artworks": nil,
			"cover": map[string]interface{}{
				"id":    136520,
				"width": 1000,
				"url":   "//images.igdb.com/igdb/image/upload/t_cover_big/co1j8f.jpg",
			},
			"first_release_date": 1163721600,
			"franchise":          0,
			"game_status":        0,
			"game_type":          10,
			"genres":             []int{8},
			"id":                 20663,
			"name":               "Tokobot Plus: Mysteries of the Karakuri",
			"platforms": []map[string]interface{}{
				{
					"id":   6,
					"name": "NES",
				},
			},
			"storyline": "",
			"summary":   "Bolt is an apprentice Treasure Master working at Mr. Canewood's laboratory. With the help of mysterious Tokobots, he will explore many prehistoric ruins to find ancient relics that will put him one step further on his path to be a Treasure Master and find his father. However, he will also change his destiny in other ways.",
			"videos":    nil,
		},
	}
)

func GetTestTwitchAuthServer() *httptest.Server {
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

func GetTestIGDBServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		// if err != nil {
		// 	// fmt.Errorf("failed to read request body: %v", err)
		// 	t.Fail()
		// }

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if strings.Contains(string(body), "search") {
			json.NewEncoder(w).Encode(testSearchGameNameResponse)
		} else if strings.Contains(string(body), "name, updated_at") {
			json.NewEncoder(w).Encode(testGetPlatformData)
		} else {
			json.NewEncoder(w).Encode(testGetGameDataResponse)
		}
	}))
}
