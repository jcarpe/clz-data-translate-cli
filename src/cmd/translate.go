package cmd

import (
	"encoding/json"
	"fmt"
	"main/src/adapters/write"
	clz_translate "main/src/domain/clz-translation"
	"os"

	"github.com/spf13/cobra"
)

var (
	seedFile       string
	writeFileName  string
	igdbSupplement bool

	translateCmd = &cobra.Command{
		Use:   "translate",
		Short: "Translate provided CLZ game collection data in XML format to JSON",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("attempt a games data translation...")

			if seedFile == "" {
				fmt.Println("seed file is required")
				return
			}

			data, err := os.ReadFile(seedFile)
			if err != nil {
				fmt.Printf("error reading CLZ data: %v", err)
				return
			}

			translated := clz_translate.TranslateCLZ(string(data), igdbSupplement)

			if writeFileName != "" {
				jsonData, marshalErr := json.Marshal(translated)
				if marshalErr != nil {
					fmt.Printf("error marshalling translated data JSON: %v", err)
					return
				}

				writeErr := write.WriteFile(jsonData, writeFileName+".json")
				if writeErr != nil {
					fmt.Printf("error writing to file: %v", err)
					return
				}
				fmt.Printf("translated JSON data written to file: %s.json \n", writeFileName)
			} else {
				fmt.Println("no filename provided, skipping write to file...")
				fmt.Printf("translated JSON data: %#v \n", translated)
			}
		},
	}
)

func init() {
	translateCmd.Flags().StringVarP(&seedFile, "seedFile", "s", "", "seed data file to translate (CLZ collection XML export)")
	translateCmd.Flags().StringVarP(&writeFileName, "writeFileName", "w", "", "filename to write JSON data to")
	translateCmd.Flags().BoolVarP(&igdbSupplement, "igdbSupplement", "i", false, "whether to supplement data with IGDB data")
	rootCmd.AddCommand(translateCmd)
}
