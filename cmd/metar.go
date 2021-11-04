package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theperiscope/avwx/api"
	"github.com/theperiscope/avwx/metars"
)

var metarCmd = &cobra.Command{
	Use:   "metar <station(s)>",
	Short: "Get METAR data",
	Long:  `Get METAR data from the Aviation Weather Center's Text Data Server.`,
	RunE:  metar,
	Args:  cobra.MinimumNArgs(1),
	Example: `  - single ICAO station:
    avwx metar KORD

  - multiple ICAO stations:
    avwx metar KORD PHOG

  - partial ICAO station name:
    avwx metar PH*

  - state/province: (use two-letter U.S. state or two-letter Canadian province)
    avwx metar @il

  - country: (use two-letter country abbreviation)
    avwx metar ~au

  - mix and match:
    avwx metar KORD CY* ~au @hi`,
}

func metar(cmd *cobra.Command, args []string) error {

	stations := args

	var data []metars.Metar
	var err error

	client := api.NewClient(api.DefaultApiEndPoint)

	hoursBeforeNow := int32(3)
	mostRecentForEachStation := true

	options := api.MetarOptions{
		Stations:                 &stations,
		HoursBeforeNow:           &hoursBeforeNow,
		MostRecentForEachStation: &mostRecentForEachStation,
	}

	data, err = client.GetMetar(options)

	var result []string
	for _, metar := range data {
		result = append(result, metar.RawText)
	}

	if err != nil {
		return err
	}

	fmt.Println(strings.Join(result, "\n"))
	return nil
}
