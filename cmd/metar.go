package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theperiscope/avwx/api"
)

var metarCmd = &cobra.Command{
	Use:     "metar",
	Short:   "Get METAR data",
	Long:    `Get METAR data from the Aviation Weather Center's Text Data Server.`,
	RunE:    metar,
	Args:    cobra.MinimumNArgs(0),
	Example: `   For examples and detailed description of all flags visit https://www.aviationweather.gov/dataserver/example?datatype=metar`,
}

var metarOptions api.MetarOptions

func metar(cmd *cobra.Command, args []string) (err error) {

	if len(metarOptions.Stations) == 0 {
		return errors.New("At least one station must be specified.")
	}

	client := api.NewClient(api.DefaultApiEndPoint)
	data, err := client.GetMetar(metarOptions)

	if err != nil {
		return
	}

	if len(data.Errors) > 0 {
		return errors.New("ADDS error(s): " + strings.Join(data.Errors, "\n"))
	}

	if len(data.Warnings) > 0 {
		return errors.New("ADDS warnings(s): " + strings.Join(data.Warnings, "\n"))
	}

	var result []string
	for _, metar := range data.Data.Metars {
		result = append(result, metar.RawText)
	}

	if err != nil {
		return
	}

	if len(result) > 0 {
		fmt.Println(strings.Join(result, "\n"))
	}

	return nil
}

func init() {
	metarCmd.Flags().SortFlags = false

	metarCmd.Flags().StringSliceVar(&metarOptions.Stations, "stations", []string{}, "required")
	metarCmd.Flags().Var(&metarOptions.StartTime, "startTime", "")
	metarCmd.Flags().Var(&metarOptions.EndTime, "endTime", "")
	metarCmd.Flags().Int32Var(&metarOptions.HoursBeforeNow, "hoursBeforeNow", 6, "")
	metarCmd.Flags().BoolVar(&metarOptions.MostRecent, "mostRecent", false, "")
	metarCmd.Flags().BoolVar(&metarOptions.MostRecentForEachStation, "mostRecentForEachStation", true, "")
	metarCmd.Flags().Float64Var(&metarOptions.MinLat, "minLat", 0, "")
	metarCmd.Flags().Float64Var(&metarOptions.MaxLat, "maxLat", 0, "")
	metarCmd.Flags().Float64Var(&metarOptions.MinLon, "minLon", 0, "")
	metarCmd.Flags().Float64Var(&metarOptions.MaxLon, "maxLon", 0, "")
	metarCmd.Flags().StringVar(&metarOptions.RadialDistance, "radialDistance", "", "")
	metarCmd.Flags().StringSliceVar(&metarOptions.FlightPath, "flightPath", []string{}, "")
	metarCmd.Flags().Float64Var(&metarOptions.MinDegreeDistance, "minDegreeDistance", 0, "")
	metarCmd.Flags().StringSliceVar(&metarOptions.Fields, "fields", []string{}, "")
}
