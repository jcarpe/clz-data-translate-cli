package adapters

import (
	"main/src/domain"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestTranslateCLZ(t *testing.T) {
	// Arrange
	data, err := os.ReadFile("./test/game-data-list.xml")
	if err != nil {
		t.Errorf("error reading test data: %v", err)
	}

	input := string(data)
	expectedOutput := domain.Game{
		Boxset: false,
		Completeness: domain.Completeness{
			HasBox:    false,
			HasManual: false,
			HasGame:   true,
		},
		Condition:    "",
		DateAcquired: time.Time{},
		Developers:   []string{"Sony Interactive Studios America"},
		Edition:      "Greatest Hits",
		Format:       "CD-ROM",
		Genres:       []string{"Racing", "Sports"},
		HardwareType: "Game",
		Links: []domain.Link{
			{
				Description: "1Xtreme at Core for Games",
				URL:         "http://core.collectorz.com/games/ps1/1xtreme",
			},
			{
				Description: "1Xtreme at PriceCharting.com",
				URL:         "https://www.pricecharting.com/game/Playstation/1Xtreme",
			},
		},
		Multiplayer:        false,
		Platform:           domain.PlayStation,
		PricechartingValue: 0.00,
		Publishers:         []string{"Sony Computer Entertainment America", "And Another One"},
		Quantity:           1,
		Region:             "",
		ReleaseDate:        time.Time{},
		Series:             "",
		Title:              "1Xtreme (Greatest Hits)",
	}

	// Act
	actualOutput := TranslateCLZ(input)

	// Assert
	if len(actualOutput.Games) != 8 {
		t.Errorf("expected 8 games, got %d", len(actualOutput.Games))
	}

	if !reflect.DeepEqual(actualOutput.Games[0], expectedOutput) {
		t.Errorf("\nexpected \n%#v,\ngot \n%#v", expectedOutput, actualOutput.Games[0])
	}
}
