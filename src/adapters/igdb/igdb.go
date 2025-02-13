package igdb

import "net/http"

type IGDBAdapter struct {
	authToken   string
	GetGameData func(int)
}

func retrieveAuthToken() string {
	httpClient := &http.Client{}
	httpClient.Post("https://api.igdb.com/v4/games", "application/json", nil)

	return "authToken"
}

func NewIGDBAdapter() *IGDBAdapter {

	return &IGDBAdapter{
		GetGameData: func(gameID int) {},
	}
}
