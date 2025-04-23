package mocks

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

var (
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
			"storyline": "A storyline supplement from mocked IGDB data for test.",
			"summary":   "A summary supplement from mocked IGDB data for test.",
			"videos":    []int{35343, 20256},
		},
	}

	testFuzzyFindGameDataResponse = []map[string]interface{}{
		{
			"id":        1,
			"name":      "Super Mario Bros. 3+",
			"platforms": []int{6},
		},
		{
			"id":        2,
			"name":      "Tokobot",
			"platforms": []int{6},
		},
		{
			"id":        3,
			"name":      "Super Mario Bros. 3",
			"platforms": []int{18},
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

		if strings.Contains(string(body), "fields id, name, platforms") {
			json.NewEncoder(w).Encode(testFuzzyFindGameDataResponse)
		} else {
			json.NewEncoder(w).Encode(testGetGameDataResponse)
		}
	}))
}
