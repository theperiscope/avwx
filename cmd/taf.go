package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/theperiscope/avwx/api"
	"github.com/theperiscope/avwx/tafs"
)

var tafCmd = &cobra.Command{
	Use:     "taf",
	Short:   "Get TAF data",
	Long:    `Get TAF data from the Aviation Weather Center's Text Data Server.`,
	RunE:    taf,
	Args:    cobra.MinimumNArgs(0),
	Example: `   For examples and detailed description of all flags visit https://www.aviationweather.gov/dataserver/example?datatype=taf`,
}

var prettyPrint = false
var includeMetar = false

var tafOptions api.TafOptions
var tafStartTime iso8601TimeValue
var tafEndTime iso8601TimeValue

func taf(cmd *cobra.Command, args []string) error {

	var data *tafs.Response
	var err error

	if len(tafOptions.Stations) == 0 {
		return errors.New("At least one station must be specified.")
	}

	if !tafStartTime.IsZero() {
		tafOptions.StartTime.Time = tafStartTime.Time
	}

	if !tafEndTime.IsZero() {
		tafOptions.EndTime.Time = tafEndTime.Time
	}

	client := api.NewClient(api.DefaultApiEndPoint)
	data, err = client.GetTaf(tafOptions)

	if err != nil {
		return err
	}

	if len(data.Errors) > 0 {
		return errors.New("ADDS error(s): " + strings.Join(data.Errors, "\n"))
	}

	if len(data.Warnings) > 0 {
		return errors.New("ADDS warnings(s): " + strings.Join(data.Warnings, "\n"))
	}

	var result []string
	for _, taf := range data.Data.Tafs {
		result = append(result, taf.RawText)
	}

	if includeMetar {
		metarOptions := api.MetarOptions{
			Stations:                 tafOptions.Stations,
			StartTime:                tafOptions.StartTime,
			EndTime:                  tafOptions.EndTime,
			HoursBeforeNow:           tafOptions.HoursBeforeNow,
			MostRecent:               tafOptions.MostRecent,
			MostRecentForEachStation: tafOptions.MostRecentForEachStation,
			MinLat:                   tafOptions.MinLat,
			MaxLat:                   tafOptions.MaxLat,
			MinLon:                   tafOptions.MinLon,
			MaxLon:                   tafOptions.MaxLon,
			RadialDistance:           tafOptions.RadialDistance,
			FlightPath:               tafOptions.FlightPath,
			MinDegreeDistance:        tafOptions.MinDegreeDistance,
			Fields:                   tafOptions.Fields,
		}

		metarData, err := client.GetMetar(metarOptions)

		if err != nil {
			return err
		}

		if len(data.Errors) > 0 {
			for _, e := range data.Errors {
				fmt.Println("ERROR:", e)
			}
		}

		if len(data.Warnings) > 0 {
			for _, e := range data.Warnings {
				fmt.Println("WARNING:", e)
			}
		}

		var result []string
		for _, metar := range metarData.Data.Metars {
			result = append(result, metar.RawText)
		}

		if len(result) > 0 {
			fmt.Println(strings.Join(result, "\n"))
		}
	}

	if prettyPrint {
		for i := range result {
			result[i] = strings.Replace(result[i], " FM", "\n  FM", -1)
		}
	}

	if len(result) > 0 {
		fmt.Println(strings.Join(result, "\n"))
	}

	return nil
}

func init() {
	tafCmd.Flags().BoolVarP(&prettyPrint, "pretty", "p", false, "Easier to read TAF format.")
	tafCmd.Flags().BoolVarP(&includeMetar, "metar", "m", false, "Include METAR data with TAF.")

	tafCmd.Flags().StringSliceVar(&tafOptions.Stations, "stations", []string{}, "")
	tafCmd.Flags().Var(&tafStartTime, "startTime", "")
	tafCmd.Flags().Var(&tafEndTime, "endTime", "")
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
}
