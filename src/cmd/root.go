package cmd

import (
	"main/src/domain"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "CLZTranslate",
	Short: "A translation tool for CLZ game collection XML data to JSON",
	Long:  "This tool translates CLZ game collection data in XML format to JSON format with optional IGDB supplement.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	domain.LoadEnv(".env.local")
}
