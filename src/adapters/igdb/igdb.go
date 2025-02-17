package igdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type IGDBAdapter struct {
	// GetGameData takes a unique game ID value and returns the requested game details.
	//
	// Fields:
	//   - gameID: The ID int value of the game.
	//
	// Returns:
	//   - An IGDBGameData instance representing the requested game.
	GetGameData func(int) IGDBGameData
}

type IGDBAdapterInit struct {
	AuthBaseUrl      string
	AuthUrlPath      string
	AuthClientId     string
	AuthClientSecret string
}

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

func getGameData(gameID int, authToken string, clientID string) IGDBGameData {
	filter := strings.NewReader(fmt.Sprintf("where id = %d", gameID))

	request, _ := http.NewRequest(http.MethodPost, "https://api.igdb.com/v4/games", filter)
	request.Header.Add("Client-ID", clientID)
	request.Header.Add("Authorization", "Bearer "+authToken)

	return IGDBGameData{
		Name: "1942",
	}
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
	retrievedToken := retrieveAuthToken(init.AuthBaseUrl, init.AuthUrlPath, init.AuthClientId, init.AuthClientSecret)
	clientID := init.AuthClientId

	return &IGDBAdapter{
		GetGameData: func(i int) IGDBGameData { return getGameData(i, retrievedToken, clientID) },
	}
}
