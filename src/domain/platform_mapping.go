package domain

// CLZPlatformMap is a mapping of CLZ platform names to their corresponding IGDB platform names.
var CLZPlatformMap = map[string]string{
	"NES":                  "Nintendo Entertainment System",
	"Saturn":               "Sega Saturn",
	"Genesis / Mega Drive": "Sega Mega Drive/Genesis",
	"Game Boy":             "Game Boy",
	"PlayStation":          "PlayStation",
	"PlayStation 2":        "PlayStation 2",
	"PlayStation 3":        "PlayStation 3",
	"PlayStation 4":        "PlayStation 4",
	"PlayStation 5":        "PlayStation 5",
}

type PlatformMapping struct {
	CLZToIGDB map[string]int
	IGDBToCLZ map[int]string
}

var PlatformMap = PlatformMapping{
	CLZToIGDB: map[string]int{
		"Atari 2600/VCS":            59,
		"Family Computer / Famicom": 99,
		"Super Famicom":             58,
		"NES":                       18,
		"SNES":                      19,
		"Nintendo 64":               4,
		"Wii":                       5,
		"Nintendo Switch":           130,
		"Saturn":                    32,
		"Dreamcast":                 23,
		"Genesis / Mega Drive":      29,
		"Game Boy":                  33,
		"Game Boy Advance":          24,
		"Game Gear":                 35,
		"PlayStation":               7,
		"PlayStation 2":             8,
		"PlayStation 3":             9,
		"PlayStation 4":             48,
		"PlayStation 5":             167,
		"PlayStation Vita":          46,
		"PSP":                       38,
		"Xbox 360":                  12,
	},
	IGDBToCLZ: map[int]string{
		59:  "Atari 2600/VCS",
		99:  "Family Computer / Famicom",
		58:  "Super Famicom",
		18:  "NES",
		19:  "SNES",
		4:   "Nintendo 64",
		5:   "Wii",
		130: "Nintendo Switch",
		32:  "Saturn",
		23:  "Dreamcast",
		29:  "Genesis / Mega Drive",
		33:  "Game Boy",
		24:  "Game Boy Advance",
		35:  "Game Gear",
		7:   "PlayStation",
		8:   "PlayStation 2",
		9:   "PlayStation 3",
		48:  "PlayStation 4",
		167: "PlayStation 5",
		46:  "PlayStation Vita",
		38:  "PSP",
		12:  "Xbox 360",
	},
}
