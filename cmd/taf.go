package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theperiscope/avwx/api"
)

var tafCmd = &cobra.Command{
	Use:     "taf",
	Short:   "Get TAF data",
	Long:    `Get TAF data from the Aviation Weather Center's Text Data Server.`,
	RunE:    taf,
	Args:    cobra.MinimumNArgs(0),
	Example: `   For examples and detailed description of all flags visit https://www.aviationweather.gov/dataserver/example?datatype=taf`,
}

var tafOptions api.TafOptions
var tafOutputFormat = api.NewEnumValue([]string{"json", "json-pretty", "rawtextonly", "rawtextonly-pretty"}, "rawtextonly-pretty")

func taf(cmd *cobra.Command, args []string) (err error) {

	client := api.NewClient(api.DefaultApiEndPoint)
	data, err := client.GetTaf(tafOptions)

	if err != nil {
		return
	}

	switch tafOutputFormat.String() {
	case "json":
		s, e := data.ToJson()
		if e != nil {
			return e
		}
		fmt.Println(s)
	case "json-pretty":
		s, e := data.ToJsonIndented()
		if e != nil {
			return e
		}
		fmt.Println(s)
	case "rawtextonly":
	case "rawtextonly-pretty":
		if len(data.Errors) > 0 {
			return errors.New("ADDS error(s): " + strings.Join(data.Errors, "\n"))
		}

		if len(data.Warnings) > 0 {
			return errors.New("ADDS warnings(s): " + strings.Join(data.Warnings, "\n"))
		}

		if tafOutputFormat.String() == "rawtextonly-pretty" {
			fmt.Println(strings.Replace(strings.Join(data.ToRawTextOnly(), "\n"), " FM", "\n  FM", -1))
		} else {
			fmt.Println(strings.Join(data.ToRawTextOnly(), "\n"))
		}

	default:
		err = fmt.Errorf("invalid output format '%s'", tafOutputFormat)
		return
	}

	return
}

func init() {
	tafCmd.Flags().SortFlags = false

	tafCmd.Flags().StringSliceVar(&tafOptions.Stations, "stations", []string{}, "")
	tafCmd.MarkFlagRequired("stations")
	tafCmd.Flags().Var(&tafOptions.StartTime, "startTime", "")
	tafCmd.Flags().Var(&tafOptions.EndTime, "endTime", "")
	tafCmd.Flags().StringVar(&tafOptions.TimeType, "timeType", "", "")
	tafCmd.Flags().Int32Var(&tafOptions.HoursBeforeNow, "hoursBeforeNow", 6, "")
	tafCmd.Flags().BoolVar(&tafOptions.MostRecent, "mostRecent", false, "")
	tafCmd.Flags().BoolVar(&tafOptions.MostRecentForEachStation, "mostRecentForEachStation", true, "")
	tafCmd.Flags().Float64Var(&tafOptions.MinLat, "minLat", 0, "")
	tafCmd.Flags().Float64Var(&tafOptions.MaxLat, "maxLat", 0, "")
	tafCmd.Flags().Float64Var(&tafOptions.MinLon, "minLon", 0, "")
	tafCmd.Flags().Float64Var(&tafOptions.MaxLon, "maxLon", 0, "")
	tafCmd.Flags().StringVar(&tafOptions.RadialDistance, "radialDistance", "", "")
	tafCmd.Flags().StringSliceVar(&tafOptions.FlightPath, "flightPath", []string{}, "")
	tafCmd.Flags().Float64Var(&tafOptions.MinDegreeDistance, "minDegreeDistance", 0, "")
	tafCmd.Flags().StringSliceVar(&tafOptions.Fields, "fields", []string{}, "")

	tafCmd.Flags().Var(tafOutputFormat, "output", "")
}
