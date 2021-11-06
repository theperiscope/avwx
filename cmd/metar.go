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

	var err error

	options := api.MetarOptions{}
	elems := reflect.ValueOf(&options).Elem()

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

				// expect be non Ptrs in the structs
				if ff.Type().Kind() == reflect.Ptr {
					continue
				}

				t := ff.Type()

				if t.String() == "iso8601.Time" {
					vv, err := iso8601.ParseString(v)
					if err == nil {
						vvv := iso8601.Time{Time: vv}
						ff.Set(reflect.ValueOf(vvv))
					}
					continue
				}

				if t.Kind() == reflect.Slice {
					if t.Elem().Kind() == reflect.String {
						vv := strings.Split(v, " ")
						ff.Set(reflect.ValueOf(vv))
					}
				} else if t.Kind() == reflect.String {
					ff.Set(reflect.ValueOf(&v))
				} else if t.Kind() == reflect.Float64 {
					vv, err := strconv.ParseFloat(v, 64)
					if err == nil {
						ff.Set(reflect.ValueOf(vv))
					}
				} else if t.Kind() == reflect.Int32 {
					vv, err := strconv.ParseInt(v, 10, 32)
					vvv := int32(vv)
					if err == nil {
						ff.Set(reflect.ValueOf(vvv))
					}
				} else if t.Kind() == reflect.Bool {
					vv, err := strconv.ParseBool(v)
					if err == nil {
						ff.Set(reflect.ValueOf(vv))
					}
				}
			}
		}
	}

	client := api.NewClient(api.DefaultApiEndPoint)
	data, err := client.GetMetar(options)

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
	for _, metar := range data.Data.Metars {
		result = append(result, metar.RawText)
	}

	if err != nil {
		return err
	}

	if len(result) > 0 {
		fmt.Println(strings.Join(result, "\n"))
	}

	return nil
}
