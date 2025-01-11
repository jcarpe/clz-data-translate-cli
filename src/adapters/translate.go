package adapters

import (
	"encoding/xml"
	"log"
	"main/src/domain"
	"time"
)

type clzXML struct {
	XMLName                    xml.Name `xml:"game"`
	PricechartingURL           string   `xml:"pricechartingurl"`
	PricechartingLoose         float64  `xml:"pricechartingloose"`
	PricechartingCIB           float64  `xml:"pricechartingcib"`
	PricechartingNew           float64  `xml:"pricechartingnew"`
	Platform									 Naming 	`xml:"platform"`
	CompletenessNum            string   `xml:"completenessnum"`
	Completeness               string   `xml:"completeness"`
	Condition									 string   `xml:"condition"`
	LastModified               Date     `xml:"lastmodified>date"`
	Quantity									 int      `xml:"quantity"`
	Language									 string   `xml:"language"`
	Publishers								 []Naming `xml:"publishers>publisher"`
	Developers								 []Naming `xml:"developers>developer"`
	Genres										 []Naming `xml:"genres>genre"`
	DateAdded                  Date     `xml:"dateadded>date"`
	ReleaseDate								 Date     `xml:"releasedate>date"`
	GameHardwareType           Naming   `xml:"gameshardware"`
	ThumbFilePath              string   `xml:"thumbfilepath"`
	BPGameID                   int      `xml:"bpgameid"`
	Region										 Naming   `xml:"region"`
	BPMediaID                  int      `xml:"bpmediaid"`
	CLZPlatformID              int      `xml:"clzplatformid"`
	BPGameLastReceivedRevision int      `xml:"bpgamelastreceivedrevision"`
	BPMediaLastReceivedRevision int     `xml:"bpmedialastreceivedrevision"`
	Multiplayer								 string   `xml:"multiplayer"`
	Format										 Naming	 	`xml:"format"`
	StorageDevice              string   `xml:"storagedevice"`
	SubmissionDate             string   `xml:"submissiondate"`
	Tags                       string   `xml:"tags"`
	TitleFirstLetter           Naming   `xml:"titlefirstletter"`
	Title									 		 string   `xml:"title"`
	Edition                    Naming   `xml:"edition"`
	Boxset                     string   `xml:"boxset"`
	HasBox                     string 	`xml:"hasbox"`
	HasManual                  string 	`xml:"hasmanual"`
	Links										   []Link		`xml:"links>link"`
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
	URLType		  string `xml:"urltype"`
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
      URL: link.URL,
    })
  }

  return domainLinks
}

func TranslateCLZ(input string) domain.Game {
	var clzData clzXML

	err := xml.Unmarshal([]byte(input), &clzData)
	if err != nil {
		log.Fatalf("error unmarshalling xml: %v", err)
	}

	gameInstance := domain.Game{
		Boxset: clzData.Boxset == "true",
		Completeness: domain.Completeness{
			HasBox: clzData.HasBox == "true",
			HasManual: clzData.HasManual == "true",
			HasGame: clzData.Quantity > 0,
		},
		Condition: clzData.Condition,
		DateAcquired: clzData.DateAdded.Value,
		Developers: extractDisplayNames(clzData.Developers),
		Edition: clzData.Edition.DisplayName,
		Format: clzData.Format.DisplayName,
		Genres: extractDisplayNames(clzData.Genres),
		HardwareType: clzData.GameHardwareType.DisplayName,
    Links: extractLinks(clzData.Links),
		Multiplayer: clzData.Multiplayer == "true",
		Platform: domain.Platform(clzData.Platform.DisplayName),
		PricechartingValue: 0.00,
		Publishers: extractDisplayNames(clzData.Publishers),
		Quantity: clzData.Quantity,
		Region: clzData.Region.DisplayName,
		ReleaseDate: clzData.ReleaseDate.Value,
    Series: "",
		Title: clzData.Title,
	}

	return gameInstance
}