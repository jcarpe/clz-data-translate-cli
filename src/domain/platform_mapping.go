package domain

type PlatformMapping struct {
	CLZName  string
	IGDBName string
}

// CLZPlatformMap is a mapping of CLZ platform names to their corresponding IGDB platform names.
var CLZPlatformMap = map[string]string{
	"NES":                  "Nintendo Entertainment System",
	"Saturn":               "Sega Saturn",
	"Genesis / Mega Drive": "Sega Mega Drive/Genesis",
	"Game Boy":             "Game Boy",
	"PlayStation":          "PlayStation",
	"PlayStation 2":        "PlayStation 2",
}
