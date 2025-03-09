package igdb

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
	IGDBBaseUrl      string
}

type IGDBGameData struct {
	ID   int
	Name string
}
