package igdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type IGDBAdapter struct {
	AuthToken   string
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

func getGameData(gameID int) IGDBGameData {
	// Retrieve game data from IGDB
	return IGDBGameData{
		Name: "1942",
	}
}

func NewIGDBAdapter(init IGDBAdapterInit) *IGDBAdapter {
	retrievedToken := retrieveAuthToken(init.AuthBaseUrl, init.AuthUrlPath, init.AuthClientId, init.AuthClientSecret)

	return &IGDBAdapter{
		AuthToken:   retrievedToken,
		GetGameData: getGameData,
	}
}
