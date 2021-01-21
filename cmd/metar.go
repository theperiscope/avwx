package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"theperiscope.org/avwx/internal/pkg/metars"
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

	var data []string
	var err error

	data, err = metars.GetData(stations)

	if err != nil {
		return err
	}

	fmt.Println(strings.Join(data, "\n"))
	return nil
}
