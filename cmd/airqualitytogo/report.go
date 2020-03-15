package cmd

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
	scraper "github.com/trampfox/air-quality-to-go/pkg/scraper/cittametropolitana"
)

func init() {
	rootCmd.AddCommand(reportCmd)
}

var reportCmd = &cobra.Command{
	Use:   "report [date YYYYmmaa format]",
	Short: "Download air quality data for the specified date",
	Long:  `Download air quality data for the specified date`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("report date is required")
		}
		if len(args) > 1 {
			return errors.New("too many arguments")
		}
		if isValidDate(args[0]) {
			return nil
		}
		return errors.New("invalid date")
	},
	Run: func(cmd *cobra.Command, args []string) {
		rs, err := scraper.ReportScraper(args[0])
		if err != nil {
			panic(err)
		}

		data := rs.GetStringData()
		fmt.Println(data)
	},
}

func isValidDate(date string) bool {
	validD := regexp.MustCompile(`20([0-9]){2}([0123]{1}[0-9]{1})([0123]{1}[0-9]{1})$`)
	return validD.MatchString(date)
}
