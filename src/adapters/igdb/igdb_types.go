package igdb

// IGDBAdapter is an adapter for the IGDB API.
//
// Fields:
//   - GetGameData: A function that retrieves game data from the IGDB API.
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

// IGDBAdapterInit contains the initialization parameters for the IGDBAdapter.
//
// Fields:
//   - AuthBaseUrl: The base URL for the authentication endpoint.
//   - AuthUrlPath: The URL path for the authentication endpoint.
//   - AuthClientId: The client ID for authentication.
//   - AuthClientSecret: The client secret for authentication.
//   - IGDBBaseUrl: The base URL for the IGDB API.
type IGDBAdapterInit struct {
	AuthBaseUrl      string
	AuthUrlPath      string
	AuthClientId     string
	AuthClientSecret string
	IGDBBaseUrl      string
}

// IGDBGameData represents the data structure for a game retrieved from the IGDB API.
//
// Fields:
//   - ID: The unique ID value of the game.
//   - Name: The name of the game.
type IGDBGameData struct {
	ID   int
	Name string
}
