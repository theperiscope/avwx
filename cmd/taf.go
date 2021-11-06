package cmd

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/relvacode/iso8601"
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

	var data *tafs.Response
	var err error

	tafOptions := api.TafOptions{}
	elems := reflect.ValueOf(&tafOptions).Elem()

	argsMap := make(map[string]string)
	flag.Parse()
	currentArg := ""
	for _, f := range flag.Args() {
		if strings.HasPrefix(f, "--") && elems.FieldByName(strings.TrimPrefix(f, "--")).IsValid() {
			currentArg = f
			continue
		}

		if len(currentArg) > 0 && len(f) > 0 {
			argsMap[strings.TrimPrefix(currentArg, "--")] = f
		}
	}

	for k, v := range argsMap {
		ff := elems.FieldByName(k)

		if ff.IsValid() {
			if ff.CanSet() {

				// expect be all Ptrs in the structs
				if ff.Type().Kind() != reflect.Ptr {
					continue
				}

				ptrType := ff.Type().Elem()

				if ptrType.String() == "iso8601.Time" {
					vv, err := iso8601.ParseString(v)
					if err == nil {
						vvv := iso8601.Time{Time: vv}
						ff.Set(reflect.ValueOf(&vvv))
					}
					continue
				}

				if ptrType.Kind() == reflect.Slice {
					if ptrType.Elem().Kind() == reflect.String {
						vv := strings.Split(v, " ")
						ff.Set(reflect.ValueOf(&vv))
					}
				} else if ptrType.Kind() == reflect.String {
					ff.Set(reflect.ValueOf(&v))
				} else if ptrType.Kind() == reflect.Float64 {
					vv, err := strconv.ParseFloat(v, 64)
					if err == nil {
						ff.Set(reflect.ValueOf(&vv))
					}
				} else if ptrType.Kind() == reflect.Int32 {
					vv, err := strconv.ParseInt(v, 10, 32)
					vvv := int32(vv)
					if err == nil {
						ff.Set(reflect.ValueOf(&vvv))
					}
				} else if ptrType.Kind() == reflect.Bool {
					vv, err := strconv.ParseBool(v)
					if err == nil {
						ff.Set(reflect.ValueOf(&vv))
					}
				}
			}
		}
	}

	client := api.NewClient(api.DefaultApiEndPoint)

	data, err = client.GetTaf(tafOptions)

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
}
