package igdb

// import "net/http"

type IGDBAdapter struct {
	// authToken   string
	GetGameData func(int) IGDBGameData
}

// func retrieveAuthToken() string {
// 	httpClient := &http.Client{}
// 	httpClient.Post("https://api.igdb.com/v4/games", "application/json", nil)

// 	return "authToken"
// }

func getGameData(gameID int) IGDBGameData {
	// Retrieve game data from IGDB
	return IGDBGameData{
		Name: "1942",
	}
}

func NewIGDBAdapter() *IGDBAdapter {

	return &IGDBAdapter{
		GetGameData: getGameData,
	}
}
