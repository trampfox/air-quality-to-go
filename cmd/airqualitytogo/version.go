package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of AirQualityToGo",
	Long:  `All software has versions. This is AirQualityToGo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AirQualityToGo v0.1")
	},
}
