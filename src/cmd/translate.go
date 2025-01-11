package cmd

import (
	"encoding/json"
	"fmt"
	"main/src/adapters"
	"os"

	"github.com/spf13/cobra"
)

var (
	seedFile      string
	writeFileName string

	translateCmd = &cobra.Command{
		Use:   "translate",
		Short: "translate game collection data to another format",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("attempt a games data translation...")

			if seedFile == "" {
				fmt.Println("seed file is required")
				return
			}

			data, err := os.ReadFile(seedFile)
			if err != nil {
				fmt.Printf("error reading test data: %v", err)
				return
			}

			translated := adapters.TranslateCLZ(string(data))

			if writeFileName != "" {
				jsonData, err := json.Marshal(translated)
				if err != nil {
					fmt.Printf("error marshalling translated data: %v", err)
					return
				}

				err = os.WriteFile(writeFileName+".json", jsonData, os.FileMode(0644))
				if err != nil {
					fmt.Printf("error writing to file: %v", err)
					return
				}
			} else {
				fmt.Println("no filename provided, skipping write to file...")
				fmt.Printf("translated JSON data: %#v \n", translated)
			}
		},
	}
)

func init() {
	translateCmd.Flags().StringVarP(&seedFile, "seedFile", "s", "", "seed data file to translate")
	translateCmd.Flags().StringVarP(&writeFileName, "writeFileName", "w", "", "filename to write JSON data to")
	rootCmd.AddCommand(translateCmd)
}
