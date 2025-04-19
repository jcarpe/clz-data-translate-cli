package igdb

import (
	"encoding/json"
	"fmt"
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

func getPlatformData() []IGDBPlatformData {
	request := initIGDBRequestObject("/platforms", strings.NewReader("fields name, updated_at;"))

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)

	if err != nil || response == nil || response.StatusCode != http.StatusOK {
		fmt.Printf("error getting platform data: %v\n, %v", err, response)
		return []IGDBPlatformData{}
	}
	defer response.Body.Close()

	var platformData []IGDBPlatformData
	if err := json.NewDecoder(response.Body).Decode(&platformData); err != nil {
		fmt.Printf("error decoding platform response body: %v\n", err)
		return []IGDBPlatformData{}
	}

	return platformData
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

func searchByTerm(searchTerm string) []IGDBGameData {
	request := initIGDBRequestObject("/games", strings.NewReader(fmt.Sprintf("search \"%s\"; fields *, platforms.name, cover.url, cover.width;", searchTerm)))

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)

	if err != nil || response == nil || response.StatusCode != http.StatusOK {
		fmt.Printf("error getting game data: %v\n, %v", err, response)
		return []IGDBGameData{}
	}
	defer response.Body.Close()

	var searchResults []IGDBGameData

	if err := json.NewDecoder(response.Body).Decode(&searchResults); err != nil {
		fmt.Printf("error decoding response body: %v\n", err)
		return []IGDBGameData{}
	}

	return searchResults
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
		GetPlatformData:  func() []IGDBPlatformData { return getPlatformData() },
		GetGameData:      func(gameID int) IGDBGameData { return getGameData(gameID) },
		SearchGameByTerm: func(searchTerm string) []IGDBGameData { return searchByTerm(searchTerm) },
	}
}
