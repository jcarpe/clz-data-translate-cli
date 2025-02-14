package igdb

import "net/http"

type IGDBAdapter struct {
	AuthToken   string
	GetGameData func(int) IGDBGameData
}

func retrieveAuthToken() string {
	httpClient := &http.Client{}

	request, _ := http.NewRequest(http.MethodPost, "http://localhost:3000/v4/games", nil)
	request.Header.Add("Content-Type", "application/json")

	response, err := httpClient.Do(request)
	if err != nil || response == nil || response.StatusCode != http.StatusOK {
		return "failed"
	}

	return "authToken"
}

func getGameData(gameID int) IGDBGameData {
	// Retrieve game data from IGDB
	return IGDBGameData{
		Name: "1942",
	}
}

func NewIGDBAdapter() *IGDBAdapter {

	retrievedToken := retrieveAuthToken()

	return &IGDBAdapter{
		AuthToken:   retrievedToken,
		GetGameData: getGameData,
	}
}
