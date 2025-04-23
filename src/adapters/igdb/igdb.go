package igdb

import (
	"encoding/json"
	"fmt"
	"main/src/domain"
	"net/http"
	"net/url"
	"strings"
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

func getGameData(gameID int) IGDBGameData {
	request := initIGDBRequestObject("/games", strings.NewReader(fmt.Sprintf("fields *, platforms.name, cover.url, cover.width; where id = %d;", gameID)))

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)

	if err != nil || response == nil || response.StatusCode != http.StatusOK {
		fmt.Printf("error getting game data: %v\n, %v", err, response)
		return IGDBGameData{}
	}
	defer response.Body.Close()

	var gameData []IGDBGameData
	if err := json.NewDecoder(response.Body).Decode(&gameData); err != nil {
		fmt.Printf("error decoding response body: %v\n", err)
		return IGDBGameData{}
	}

	return gameData[0]
}

func fuzzyFindIGDBGameByTitle(title string, clzPlatformName string) int {
	// Normalize the game title
	normalizedTitle := GameTitleNormalization(title)

	// Search for the game by normalized title
	gamesData := fuzzySearchByTerm(normalizedTitle)
	if len(gamesData) == 0 {
		fmt.Printf("No games found in FuzzyFind for title: %s\n", title)
		return 0
	}

	// if multiple games returned, find the entry with a matching platform
	if len(gamesData) > 1 {
		for _, game := range gamesData {
			for _, platform := range game.Platforms {
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
		GetGameData:          func(gameID int) IGDBGameData { return getGameData(gameID) },
		FuzzyFindGameByTitle: func(title string, clzPlatform string) int { return fuzzyFindIGDBGameByTitle(title, clzPlatform) },
	}
}
