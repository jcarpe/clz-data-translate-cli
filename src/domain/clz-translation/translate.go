package clz_translate

import (
	"encoding/xml"
	"fmt"
	"log"
	"main/src/adapters/igdb"
	"main/src/domain"
	"os"
	"strconv"
	"time"
)

type clzXMLList struct {
	GameList []clzXML `xml:"gamelist>game"`
}

type clzXML struct {
	XMLName                     xml.Name    `xml:"game"`
	PricechartingURL            string      `xml:"pricechartingurl"`
	PricechartingLoose          float64     `xml:"pricechartingloose"`
	PricechartingCIB            float64     `xml:"pricechartingcib"`
	PricechartingNew            float64     `xml:"pricechartingnew"`
	PricechartingValue          float64     `xml:"pricechartingvalue"`
	Platform                    namingDef   `xml:"platform"`
	CompletenessNum             string      `xml:"completenessnum"`
	Completeness                string      `xml:"completeness"`
	Condition                   string      `xml:"condition"`
	LastModified                dateDef     `xml:"lastmodified>date"`
	Quantity                    int         `xml:"quantity"`
	Language                    string      `xml:"language"`
	Publishers                  []namingDef `xml:"publishers>publisher"`
	Developers                  []namingDef `xml:"developers>developer"`
	Genres                      []namingDef `xml:"genres>genre"`
	DateAdded                   dateDef     `xml:"dateadded>date"`
	ReleaseDate                 dateDef     `xml:"releasedate>date"`
	GameHardwareType            namingDef   `xml:"gameshardware"`
	ThumbFilePath               string      `xml:"thumbfilepath"`
	BPGameID                    int         `xml:"bpgameid"`
	Region                      namingDef   `xml:"region"`
	BPMediaID                   int         `xml:"bpmediaid"`
	CLZPlatformID               int         `xml:"clzplatformid"`
	BPGameLastReceivedRevision  int         `xml:"bpgamelastreceivedrevision"`
	BPMediaLastReceivedRevision int         `xml:"bpmedialastreceivedrevision"`
	Multiplayer                 string      `xml:"multiplayer"`
	Format                      namingDef   `xml:"format"`
	StorageDevice               string      `xml:"storagedevice"`
	SubmissionDate              string      `xml:"submissiondate"`
	Tags                        string      `xml:"tags"`
	TitleFirstLetter            namingDef   `xml:"titlefirstletter"`
	Title                       string      `xml:"title"`
	Edition                     namingDef   `xml:"edition"`
	Boxset                      string      `xml:"boxset"`
	HasBox                      string      `xml:"hasbox"`
	HasManual                   string      `xml:"hasmanual"`
	Links                       []linkDef   `xml:"links>link"`
}

type dateDef struct {
	Value time.Time
}

type namingDef struct {
	DisplayName string `xml:"displayname"`
	SortName    string `xml:"sortname"`
}

type linkDef struct {
	Description string `xml:"description"`
	URL         string `xml:"url"`
	URLType     string `xml:"urltype"`
}

func extractDisplayNames(namings []namingDef) []string {
	var names []string

	for _, naming := range namings {
		names = append(names, naming.DisplayName)
	}

	return names
}

func extractLinks(links []linkDef) []domain.Link {
	var domainLinks []domain.Link

	for _, link := range links {
		domainLinks = append(domainLinks, domain.Link{
			Description: link.Description,
			URL:         link.URL,
		})
	}

	return domainLinks
}

func retrieveIGDBSupplement(game domain.Game, igdbAdapter *igdb.IGDBAdapter) igdb.IGDBGameData {

	igdbGameData := igdbAdapter.GetGameData(game.IGDB_ID)
	if igdbGameData.ID == 0 {
		fmt.Println("No game data found in IGDB for CLZ game:", game.Title)
		return igdb.IGDBGameData{}
	}

	rateLimitStr := os.Getenv("IGDB_API_RATE_LIMIT")
	rateLimit, err := strconv.Atoi(rateLimitStr)
	if err != nil {
		fmt.Printf("Invalid IGDB_API_RATE_LIMIT value: %v\n", err)
		rateLimit = 0 // Default to 0 second if parsing fails
	}
	sleepTime := time.Duration(rateLimit) * time.Second
	time.Sleep(sleepTime)

	return igdbGameData
}

func translateGamesDataToDomain(clzXMLData string) []domain.Game {
	var (
		clzData clzXMLList
	)

	err := xml.Unmarshal([]byte(clzXMLData), &clzData)
	if err != nil {
		log.Fatalf("error unmarshalling xml: %v", err)
	}

	gameCollection := domain.GameCollection{
		Games: []domain.Game{},
	}

	for _, game := range clzData.GameList {
		newGame := domain.Game{
			Boxset: game.Boxset == "true",
			Completeness: domain.Completeness{
				HasBox:    game.HasBox == "true",
				HasManual: game.HasManual == "true",
				HasGame:   game.Quantity > 0,
			},
			Condition:          game.Condition,
			DateAcquired:       game.DateAdded.Value,
			Developers:         extractDisplayNames(game.Developers),
			Edition:            game.Edition.DisplayName,
			Format:             game.Format.DisplayName,
			Genres:             extractDisplayNames(game.Genres),
			HardwareType:       game.GameHardwareType.DisplayName,
			Links:              extractLinks(game.Links),
			Multiplayer:        game.Multiplayer == "true",
			Platform:           domain.Platform(game.Platform.DisplayName),
			PricechartingValue: game.PricechartingValue,
			Publishers:         extractDisplayNames(game.Publishers),
			Quantity:           game.Quantity,
			Region:             game.Region.DisplayName,
			ReleaseDate:        game.ReleaseDate.Value,
			Series:             "",
			Title:              game.Title,
		}

		gameCollection.Games = append(gameCollection.Games, newGame)
	}

	return gameCollection.Games
}

// TranslateCLZ translates a CLZ XML input string into a domain.GameCollection.
// It unmarshals the XML input into a clzXMLList structure and then iterates
// through the list of games to populate a domain.GameCollection with the
// relevant game data.
//
// Parameters:
//   - input: A string containing the CLZ XML data.
//
// Returns:
//   - domain.GameCollection: A collection of games translated from the CLZ XML data.
//   - igdbSupplement: A boolean indicating whether to supplement the data with IGDB data.
//
// The function will log a fatal error if the XML unmarshalling fails.
func TranslateCLZ(input string, igdbSupplement bool) domain.GameCollection {
	var (
		igdbAdapter *igdb.IGDBAdapter = nil
	)

	gameCollection := translateGamesDataToDomain(input)

	if igdbSupplement {
		igdbAdapter = igdb.NewIGDBAdapter(igdb.IGDBAdapterInit{
			AuthBaseUrl:      os.Getenv("IGDB_AUTH_BASE_URL"),
			AuthUrlPath:      os.Getenv("IGDB_AUTH_PATH"),
			AuthClientId:     os.Getenv("IGDB_CLIENT_ID"),
			AuthClientSecret: os.Getenv("IGDB_CLIENT_SECRET"),
			IGDBBaseUrl:      os.Getenv("IGDB_BASE_URL"),
		})

		// perform fuzzy find for all games in order to get IGDB_ID
		gameCollectionWithIgdbIds := igdbAdapter.FuzzyFindGamesList(gameCollection)

		for i, game := range gameCollectionWithIgdbIds {
			if game.HardwareType == "Game" {
				igdbData := retrieveIGDBSupplement(game, igdbAdapter)

				gameCollection[i].FirstReleaseDate = time.Unix(int64(igdbData.First_release_date), 0)
				gameCollection[i].Storyline = igdbData.Storyline
				gameCollection[i].Summary = igdbData.Summary
				gameCollection[i].Cover = domain.Cover{
					ID:    igdbData.Cover.ID,
					Width: igdbData.Cover.Width,
					URL:   igdbData.Cover.URL,
				}
			}
		}

		gameCollection = gameCollectionWithIgdbIds
	}

	return domain.GameCollection{
		Games: gameCollection,
	}
}
