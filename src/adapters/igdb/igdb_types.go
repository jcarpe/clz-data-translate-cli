package igdb

import (
	"main/src/domain"
)

// IGDBAdapter is an adapter for the IGDB API.
//
// Fields:
//   - GetGameData: A function that retrieves game data from the IGDB API.
//   - FuzzyFindGameByTitle: A function that searches for a game by title and platform.
type IGDBAdapter struct {
	// GetGameData takes a unique game ID value and returns the requested game details.
	//
	// Fields:
	//   - gameID: The ID int value of the game.
	//
	// Returns:
	//   - An IGDBGameData instance representing the requested game.
	GetGameData func(int) IGDBGameData

	// FuzzyFindGameByTitle takes a game title and platform name, and returns the ID of the game that matches the title and platform.
	//
	// Fields:
	//   - title: The game title string.
	//   - clzPlatform: The platform name string.
	//
	// Returns:
	//   - The ID int value of the game that matches the title and platform.
	FuzzyFindGameByTitle func(string, string) int

	// FuzzyFindGamesList takes a game title and returns a list of games that match the title.
	//
	// Fields:
	//   - gamesList: The game collection list
	//
	// Returns:
	//   - The games list with entires updated with IGDB ID values.
	FuzzyFindGamesList func([]domain.Game) []domain.Game
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

// IGDBPlatformData represents the data structure for a platform retrieved from the IGDB API.
//
// Fields:
//   - ID: The unique ID value of the platform.
//   - Name: The name of the platform.
type IGDBPlatformData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type IGDBCover struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Width int    `json:"width"`
}

// IGDBGameData represents the data structure for a game retrieved from the IGDB API.
//
// Fields:
//   - ID: The unique ID value of the game.
//   - Name: The name of the game.
type IGDBGameData struct {
	Artworks           []int              `json:"artworks"`
	Cover              IGDBCover          `json:"cover"`
	First_release_date int                `json:"first_release_date"`
	Franchise          int                `json:"franchise"`
	Game_status        int                `json:"game_status"`
	Game_type          int                `json:"game_type"`
	Genres             []int              `json:"genres"`
	ID                 int                `json:"id"`
	Name               string             `json:"name"`
	Platforms          []IGDBPlatformData `json:"platforms"`
	Storyline          string             `json:"storyline"`
	Summary            string             `json:"summary"`
	Videos             []int              `json:"videos"`
}
