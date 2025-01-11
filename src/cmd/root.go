package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "CLZTranslate",
	Short: "A simple data translation tool for vdeo game collection data from CLZ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(translateCmd)
}
