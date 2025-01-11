package adapters

import (
	"encoding/xml"
	"log"
	"main/src/domain"
	"time"
)

type clzXMLList struct {
	GameList []clzXML `xml:"gamelist>game"`
}

type clzXML struct {
	XMLName                     xml.Name `xml:"game"`
	PricechartingURL            string   `xml:"pricechartingurl"`
	PricechartingLoose          float64  `xml:"pricechartingloose"`
	PricechartingCIB            float64  `xml:"pricechartingcib"`
	PricechartingNew            float64  `xml:"pricechartingnew"`
	Platform                    Naming   `xml:"platform"`
	CompletenessNum             string   `xml:"completenessnum"`
	Completeness                string   `xml:"completeness"`
	Condition                   string   `xml:"condition"`
	LastModified                Date     `xml:"lastmodified>date"`
	Quantity                    int      `xml:"quantity"`
	Language                    string   `xml:"language"`
	Publishers                  []Naming `xml:"publishers>publisher"`
	Developers                  []Naming `xml:"developers>developer"`
	Genres                      []Naming `xml:"genres>genre"`
	DateAdded                   Date     `xml:"dateadded>date"`
	ReleaseDate                 Date     `xml:"releasedate>date"`
	GameHardwareType            Naming   `xml:"gameshardware"`
	ThumbFilePath               string   `xml:"thumbfilepath"`
	BPGameID                    int      `xml:"bpgameid"`
	Region                      Naming   `xml:"region"`
	BPMediaID                   int      `xml:"bpmediaid"`
	CLZPlatformID               int      `xml:"clzplatformid"`
	BPGameLastReceivedRevision  int      `xml:"bpgamelastreceivedrevision"`
	BPMediaLastReceivedRevision int      `xml:"bpmedialastreceivedrevision"`
	Multiplayer                 string   `xml:"multiplayer"`
	Format                      Naming   `xml:"format"`
	StorageDevice               string   `xml:"storagedevice"`
	SubmissionDate              string   `xml:"submissiondate"`
	Tags                        string   `xml:"tags"`
	TitleFirstLetter            Naming   `xml:"titlefirstletter"`
	Title                       string   `xml:"title"`
	Edition                     Naming   `xml:"edition"`
	Boxset                      string   `xml:"boxset"`
	HasBox                      string   `xml:"hasbox"`
	HasManual                   string   `xml:"hasmanual"`
	Links                       []Link   `xml:"links>link"`
}

type Date struct {
	Value time.Time
}

type Naming struct {
	DisplayName string `xml:"displayname"`
	SortName    string `xml:"sortname"`
}

type Link struct {
	Description string `xml:"description"`
	URL         string `xml:"url"`
	URLType     string `xml:"urltype"`
}

func extractDisplayNames(namings []Naming) []string {
	var names []string

	for _, naming := range namings {
		names = append(names, naming.DisplayName)
	}

	return names
}

func extractLinks(links []Link) []domain.Link {
	var domainLinks []domain.Link

	for _, link := range links {
		domainLinks = append(domainLinks, domain.Link{
			Description: link.Description,
			URL:         link.URL,
		})
	}

	return domainLinks
}

func TranslateCLZ(input string) domain.GameCollection {
	var clzData clzXMLList

	err := xml.Unmarshal([]byte(input), &clzData)
	if err != nil {
		log.Fatalf("error unmarshalling xml: %v", err)
	}

	gameCollection := domain.GameCollection{
		Games: []domain.Game{},
	}

	for _, game := range clzData.GameList {
		gameCollection.Games = append(gameCollection.Games, domain.Game{
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
			PricechartingValue: 0.00,
			Publishers:         extractDisplayNames(game.Publishers),
			Quantity:           game.Quantity,
			Region:             game.Region.DisplayName,
			ReleaseDate:        game.ReleaseDate.Value,
			Series:             "",
			Title:              game.Title,
		})
	}

	return gameCollection
}
