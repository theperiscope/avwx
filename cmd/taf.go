package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"theperiscope.org/avwx/internal/pkg/tafs"
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

func taf(cmd *cobra.Command, args []string) error {

	stations := args

	var data []string
	var err error

	data, err = tafs.GetData(stations)

	if err != nil {
		return err
	}

	if prettyPrint {
		for i := range data {
			data[i] = strings.Replace(data[i], " FM", "\n  FM", -1)
		}
	}

	fmt.Println(strings.Join(data, "\n"))
	return nil
}

func init() {
	tafCmd.Flags().BoolVarP(&prettyPrint, "pretty", "p", false, "Easier to read TAF format.")
}
