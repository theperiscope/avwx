package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theperiscope/avwx/api"
	"github.com/theperiscope/avwx/tafs"
)

var tafCmd = &cobra.Command{
	Use:   "taf <station(s)>",
	Short: "Get TAF data",
	Long:  `Get TAF data from the Aviation Weather Center's Text Data Server.`,
	RunE:  taf,
	Args:  cobra.MinimumNArgs(1),
	Example: `  - single ICAO station:
    avwx taf KORD

  - multiple ICAO stations:
    avwx taf KORD PHOG

  - partial ICAO station name:
    avwx taf PH*

  - state/province: (use two-letter U.S. state or two-letter Canadian province)
    avwx taf @il

  - country: (use two-letter country abbreviation)
    avwx taf ~au

  - mix and match:
    avwx taf KORD CY* ~au @hi`,
}

var prettyPrint = false
var includeMetar = false

func taf(cmd *cobra.Command, args []string) error {

	stations := args

	var data []tafs.Taf
	var err error

	hoursBeforeNow := int32(9)
	mostRecentForEachStation := true

	tafOptions := api.TafOptions{
		Stations:                 &stations,
		HoursBeforeNow:           &hoursBeforeNow,
		MostRecentForEachStation: &mostRecentForEachStation,
	}

	client := api.NewClient(api.DefaultApiEndPoint)

	data, err = client.GetTaf(tafOptions)

	if err != nil {
		return err
	}

	var result []string
	for _, taf := range data {
		result = append(result, taf.RawText)
	}

	if includeMetar {
		metarOptions := api.MetarOptions{
			Stations:                 &stations,
			HoursBeforeNow:           &hoursBeforeNow,
			MostRecentForEachStation: &mostRecentForEachStation,
		}

		metarData, err := client.GetMetar(metarOptions)
		if err != nil {
			return err
		}

		var result []string
		for _, metar := range metarData {
			result = append(result, metar.RawText)
		}

		fmt.Println(strings.Join(result, "\n"))
	}

	if prettyPrint {
		for i := range result {
			result[i] = strings.Replace(result[i], " FM", "\n  FM", -1)
		}
	}

	fmt.Println(strings.Join(result, "\n"))
	return nil
}

func init() {
	tafCmd.Flags().BoolVarP(&prettyPrint, "pretty", "p", false, "Easier to read TAF format.")
	tafCmd.Flags().BoolVarP(&includeMetar, "metar", "m", false, "Include METAR data with TAF.")
}
