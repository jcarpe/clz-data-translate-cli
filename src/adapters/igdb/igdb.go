package igdb

import (
	"encoding/json"
	"fmt"
	"main/src/domain"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	igdbBaseUrl string
	authToken   string
	clientID    string
)

type authResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type igdbFuzzySearchGameData struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Platforms []int  `json:"platforms"`
}

func retrieveAuthToken(baseUrl string, path string, id string, secret string) string {
	request, _ := http.NewRequest(http.MethodPost, baseUrl+path, nil)
	request.Header.Add("Content-Type", "application/json")

	// Add query parameters
	query := url.Values{}
	query.Add("client_id", id)
	query.Add("client_secret", secret)
	query.Add("grant_type", "client_credentials")
	request.URL.RawQuery = query.Encode()

	// http client
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)

	if err != nil || response == nil || response.StatusCode != http.StatusOK {
		return "failed"
	}
	defer response.Body.Close()

	var authRes authResponse
	if err := json.NewDecoder(response.Body).Decode(&authRes); err != nil {
		fmt.Printf("error decoding response body: %v\n", err)
		return "failed"
	}

	return authRes.AccessToken
}

func initIGDBRequestObject(path string, filter *strings.Reader) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, igdbBaseUrl+path, filter)
	request.Header.Add("Client-ID", clientID)
	request.Header.Add("Authorization", "Bearer "+authToken)

	return request
}

func getGameData(gameIDs []int) []IGDBGameData {
	ids := strings.Join(strings.Fields(strings.Trim(fmt.Sprint(gameIDs), "[]")), ",")
	query := fmt.Sprintf("fields *, platforms.name, cover.url, cover.width; where id = (%s);", ids)
	request := initIGDBRequestObject("/games", strings.NewReader(query))

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)

	if err != nil || response == nil || response.StatusCode != http.StatusOK {
		fmt.Printf("error getting game data: %v\n, %v", err, response)
		return []IGDBGameData{}
	}
	defer response.Body.Close()

	var gameData []IGDBGameData
	if err := json.NewDecoder(response.Body).Decode(&gameData); err != nil {
		fmt.Printf("error decoding response body: %v\n", err)
		return []IGDBGameData{}
	}

	return gameData
}

func fuzzyFindIGDBGameByTitle(title string, clzPlatformName string) int {
	// Normalize the game title
	normalizedTitle := GameTitleNormalization(title)

	fmt.Printf("FuzzyFind for title: %s\n", normalizedTitle)

	// Search for the game by normalized title
	gamesData := fuzzySearchByTerm(normalizedTitle)
	if len(gamesData) == 0 {
		fmt.Printf("FuzzyFind failed for title: %s\n", normalizedTitle)
		return 0
	}

	// if multiple games returned, find the entry with a matching platform
	if len(gamesData) > 1 {
		for _, game := range gamesData {
			fmt.Printf("FuzzyFind found game: %s\n", game.Name)
			for _, platform := range game.Platforms {
				fmt.Printf("Platform: %s\n", domain.PlatformMap.IGDBToCLZ[platform])
				if platform == domain.PlatformMap.CLZToIGDB[clzPlatformName] {
					return game.ID
				}
			}
		}
	}

	// Return the first game found
	return gamesData[0].ID
}

func fuzzySearchByTerm(searchTerm string) []igdbFuzzySearchGameData {
	request := initIGDBRequestObject("/games", strings.NewReader(fmt.Sprintf("search \"%s\"; fields id, name, platforms;", searchTerm)))

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)

	if err != nil || response == nil || response.StatusCode != http.StatusOK {
		fmt.Printf("error getting game data: %v\n, %v", err, response)
		return []igdbFuzzySearchGameData{}
	}
	defer response.Body.Close()

	var searchResults []igdbFuzzySearchGameData

	if err := json.NewDecoder(response.Body).Decode(&searchResults); err != nil {
		fmt.Printf("error decoding response body: %v\n", err)
		return []igdbFuzzySearchGameData{}
	}

	return searchResults
}

func fuzzyFindGamesList(gameList []domain.Game) []domain.Game {
	// Rate limit duration for IGDB API
	rateLimitStr := os.Getenv("IGDB_API_RATE_LIMIT")
	rateLimit, err := strconv.Atoi(rateLimitStr)
	if err != nil {
		fmt.Printf("Invalid IGDB_API_RATE_LIMIT value: %v -- setting to 0\n", err)
		rateLimit = 0 // Default to 0 second if parsing fails
	}

	sleepTime := time.Duration(rateLimit) * time.Millisecond

	for i, game := range gameList {
		if i != 0 {
			fmt.Printf("%d Sleeping for rate limit...\n", i)
			// Sleep for a short duration to avoid hitting the rate limit
			time.Sleep(sleepTime)
		}

		// Normalize the game title
		normalizedTitle := GameTitleNormalization(game.Title)

		// Search for the game by normalized title
		gameIgdbId := fuzzyFindIGDBGameByTitle(normalizedTitle, string(game.Platform))
		if gameIgdbId == 0 {
			fmt.Printf("No games found in FuzzyFind for title: %s\n", game.Title)
			continue
		}

		// Update the game ID in the game list
		gameList[i].IGDB_ID = gameIgdbId
	}

	return gameList
}

// GameTitleNormalization normalizes the game title by removing special characters and converting to lowercase.
//
// Parameters:
//   - title: The game title string.
//
// Returns:
//   - The normalized game title string.
func GameTitleNormalization(title string) string {
	// Normalize the game title by removing special characters and converting to lowercase
	normalizedTitle := strings.ToLower(title)
	normalizedTitle = strings.ReplaceAll(normalizedTitle, "(greatest hits)", "")
	normalizedTitle = strings.TrimSpace(normalizedTitle)
	return normalizedTitle
}

// NewIGDBAdapter initializes a new IGDBAdapter with the provided authentication details.
// Authenticates with the IGDB API with an access token.
//
// Parameters:
//   - init: IGDBAdapterInit struct containing the following fields:
//   - AuthBaseUrl: The base URL for the authentication endpoint.
//   - AuthUrlPath: The URL path for the authentication endpoint.
//   - AuthClientId: The client ID for authentication.
//   - AuthClientSecret: The client secret for authentication.
//
// Returns:
//   - A pointer to an IGDBAdapter instance with the retrieved authentication token and a function to get game data.
func NewIGDBAdapter(init IGDBAdapterInit) *IGDBAdapter {
	authToken = retrieveAuthToken(init.AuthBaseUrl, init.AuthUrlPath, init.AuthClientId, init.AuthClientSecret)
	clientID = init.AuthClientId
	igdbBaseUrl = init.IGDBBaseUrl

	return &IGDBAdapter{
		GetGameData:          func(gameIDs []int) []IGDBGameData { return getGameData(gameIDs) },
		FuzzyFindGameByTitle: func(title string, clzPlatform string) int { return fuzzyFindIGDBGameByTitle(title, clzPlatform) },
		FuzzyFindGamesList:   func(gameList []domain.Game) []domain.Game { return fuzzyFindGamesList(gameList) },
	}
}
