package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	scraper "github.com/trampfox/air-quality-to-go/pkg/scraper/cittametropolitana"
)

func init() {
	rootCmd.AddCommand(ipqaCmd)
}

var ipqaCmd = &cobra.Command{
	Use:   "ipqa",
	Short: "Retrieve IPQA data for today and tomorrow",
	Long:  `Retrieve from the cittametropolitana website the IPQA data for today and tomorrow`,
	Run: func(cmd *cobra.Command, args []string) {
		cis, err := scraper.CollyIPQAScraper()
		if err != nil {
			panic(err)
		}
		data := cis.GetStringData()

		fmt.Println(data)
	},
}
