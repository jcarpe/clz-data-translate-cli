package igdb

// IGDBAdapter is an adapter for the IGDB API.
//
// Fields:
//   - GetGameData: A function that retrieves game data from the IGDB API.
type IGDBAdapter struct {
	// GetPlatformData retrieves a list of platforms from the IGDB API.
	//
	// Returns:
	//   - A list of IGDBPlatformData instances representing the platforms.
	GetPlatformData func() []IGDBPlatformData

	// GetGameData takes a unique game ID value and returns the requested game details.
	//
	// Fields:
	//   - gameID: The ID int value of the game.
	//
	// Returns:
	//   - An IGDBGameData instance representing the requested game.
	GetGameData func(int) IGDBGameData

	// SearchGameByTerm takes a search term and returns a list of games that match the search term by title.
	//
	// Fields:
	//   - searchTerm: The search term string.
	//
	// Returns:
	//   - A list of IGDBGameData instances representing the games that match the search term.
	SearchGameByTerm func(string) []IGDBGameData
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

type IGDBPlatformData struct {
	Name string `json:"name"`
}

// IGDBGameData represents the data structure for a game retrieved from the IGDB API.
//
// Fields:
//   - ID: The unique ID value of the game.
//   - Name: The name of the game.
type IGDBGameData struct {
	Artworks           []int  `json:"artworks"`
	Cover              int    `json:"cover"`
	First_release_date int    `json:"first_release_date"`
	Franchise          int    `json:"franchise"`
	Game_status        int    `json:"game_status"`
	Game_type          int    `json:"game_type"`
	Genres             []int  `json:"genres"`
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Platforms          []int  `json:"platforms"`
	Storyline          string `json:"storyline"`
	Summary            string `json:"summary"`
	Videos             []int  `json:"videos"`
}
