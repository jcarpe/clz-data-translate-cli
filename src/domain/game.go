package domain

import (
	"time"
)

type GameCollection struct {
	Games []Game
}

// Game is the domain model for a video game as defined for our purposes.
type Game struct {
	Boxset             bool
	Completeness       Completeness
	Condition          string
	Cover              Cover
	DateAcquired       time.Time
	Developers         []string
	Edition            string
	FirstReleaseDate   time.Time
	Format             string
	Genres             []string
	HardwareType       string
	IGDB_ID            int
	Links              []Link
	Multiplayer        bool
	Platform           Platform
	PricechartingValue float64
	Publishers         []string
	Quantity           int
	Region             string
	ReleaseDate        time.Time
	Series             string
	Storyline          string
	Summary            string
	Title              string
}

type Cover struct {
	ID    int
	Width int
	URL   string
}

type Completeness struct {
	HasBox    bool
	HasManual bool
	HasGame   bool
}

type Link struct {
	Description string
	URL         string
}

type Platform string

// Platform represents a type for various gaming platforms.
// The constants defined below act as an enumeration of different PlayStation platforms.
const (
	PlayStation  Platform = "PlayStation"
	PlayStation2 Platform = "PlayStation 2"
	PlayStation3 Platform = "PlayStation 3"
	PlayStation4 Platform = "PlayStation 4"
	PlayStation5 Platform = "PlayStation 5"
)
