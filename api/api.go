package api

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/relvacode/iso8601"
	"github.com/theperiscope/avwx/metars"
	"github.com/theperiscope/avwx/tafs"
)

const DefaultApiEndPoint = "https://aviationweather.gov/adds/dataserver_current/httpparam"

type Client interface {
	GetMetar(options MetarOptions) (*metars.Response, error)
	GetTaf(options TafOptions) (*tafs.Response, error)
}

type client struct {
	c           *http.Client
	ApiEndPoint string
}

// MetarOptions uses pointers to distinguish set (not nil) from unset fields (nil)
type MetarOptions struct {
	Stations                 *[]string
	StartTime                *iso8601.Time
	EndTime                  *iso8601.Time
	HoursBeforeNow           *int32
	MostRecent               *bool
	MostRecentForEachStation *bool
	MinLat                   *float64
	MaxLat                   *float64
	MinLon                   *float64
	MaxLon                   *float64
	RadialDistance           *string
	FlightPath               *[]string
	MinDegreeDistance        *float64
	Fields                   *[]string
}

// MetarOptions uses pointers to distinguish set (not nil) from unset fields (nil)
type TafOptions struct {
	Stations                 *[]string
	StartTime                *iso8601.Time
	EndTime                  *iso8601.Time
	TimeType                 *string // only difference from MetarOptions
	HoursBeforeNow           *int32
	MostRecent               *bool
	MostRecentForEachStation *bool
	MinLat                   *float64
	MaxLat                   *float64
	MinLon                   *float64
	MaxLon                   *float64
	RadialDistance           *string
	FlightPath               *[]string
	MinDegreeDistance        *float64
	Fields                   *[]string
}

func NewClient(apiEndPoint string) Client {
	c := &http.Client{}

	return &client{
		c:           c,
		ApiEndPoint: apiEndPoint,
	}
}

func (c *client) GetMetar(options MetarOptions) (*metars.Response, error) {
	u, err := url.Parse(c.ApiEndPoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("dataSource", "metars")
	q.Set("requestType", "retrieve")
	q.Set("format", "xml")

	if options.Stations != nil {
		q.Set("stationString", strings.Join(*options.Stations, " "))
	}
	if options.StartTime != nil {
		q.Set("startTime", options.StartTime.UTC().Format("2006-01-02T15:04:05"))
	}
	if options.EndTime != nil {
		q.Set("endTime", options.EndTime.UTC().Format("2006-01-02T15:04:05"))
	}
	if options.HoursBeforeNow != nil {
		q.Set("hoursBeforeNow", strconv.FormatInt(int64(*options.HoursBeforeNow), 10))
	}
	if options.MostRecent != nil {
		q.Set("mostRecent", strconv.FormatBool(*options.MostRecent))
	}
	if options.MostRecentForEachStation != nil {
		q.Set("mostRecentForEachStation", strconv.FormatBool(*options.MostRecentForEachStation))
	}
	if options.MinLat != nil {
		q.Set("minLat", strconv.FormatFloat(*options.MinLat, 'f', -1, 64))
	}
	if options.MaxLat != nil {
		q.Set("maxLat", strconv.FormatFloat(*options.MaxLat, 'f', -1, 64))
	}
	if options.MinLon != nil {
		q.Set("minLon", strconv.FormatFloat(*options.MinLon, 'f', -1, 64))
	}
	if options.MaxLon != nil {
		q.Set("maxLon", strconv.FormatFloat(*options.MaxLon, 'f', -1, 64))
	}
	if options.RadialDistance != nil {
		q.Set("radialDistance", *options.RadialDistance)
	}
	if options.FlightPath != nil {
		q.Set("flightPath", strings.Join(*options.FlightPath, " "))
	}
	if options.MinDegreeDistance != nil {
		q.Set("minDegreeDistance", strconv.FormatFloat(*options.MinDegreeDistance, 'f', -1, 64))
	}
	if options.Fields != nil {
		q.Set("fields", strings.Join(*options.Fields, " "))
	}

	u.RawQuery = q.Encode()

	httpResponse, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var r metars.Response
	err = xml.Unmarshal(data, &r)
	if err != nil {
		return nil, err

	}

	return &r, nil
}

func (c *client) GetTaf(options TafOptions) (*tafs.Response, error) {
	u, err := url.Parse(c.ApiEndPoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("dataSource", "tafs")
	q.Set("requestType", "retrieve")
	q.Set("format", "xml")

	if options.Stations != nil {
		q.Set("stationString", strings.Join(*options.Stations, " "))
	}
	if options.StartTime != nil {
		q.Set("startTime", strconv.FormatInt(options.StartTime.UTC().Unix(), 10))
	}
	if options.EndTime != nil {
		//q.Set("endTime", options.EndTime.UTC().Format("2006-01-02T15:04:05"))
		q.Set("endTime", strconv.FormatInt(options.EndTime.UTC().Unix(), 10))
	}
	if options.TimeType != nil { // only difference from GetMetar
		q.Set("timeType", *options.TimeType)
	}
	if options.HoursBeforeNow != nil {
		q.Set("hoursBeforeNow", strconv.FormatInt(int64(*options.HoursBeforeNow), 10))
	}
	if options.MostRecent != nil {
		q.Set("mostRecent", strconv.FormatBool(*options.MostRecent))
	}
	if options.MostRecentForEachStation != nil {
		q.Set("mostRecentForEachStation", strconv.FormatBool(*options.MostRecentForEachStation))
	}
	if options.MinLat != nil {
		q.Set("minLat", strconv.FormatFloat(*options.MinLat, 'f', -1, 64))
	}
	if options.MaxLat != nil {
		q.Set("maxLat", strconv.FormatFloat(*options.MaxLat, 'f', -1, 64))
	}
	if options.MinLon != nil {
		q.Set("minLon", strconv.FormatFloat(*options.MinLon, 'f', -1, 64))
	}
	if options.MaxLon != nil {
		q.Set("maxLon", strconv.FormatFloat(*options.MaxLon, 'f', -1, 64))
	}
	if options.RadialDistance != nil {
		q.Set("radialDistance", *options.RadialDistance)
	}
	if options.FlightPath != nil {
		q.Set("flightPath", strings.Join(*options.FlightPath, " "))
	}
	if options.MinDegreeDistance != nil {
		q.Set("minDegreeDistance", strconv.FormatFloat(*options.MinDegreeDistance, 'f', -1, 64))
	}
	if options.Fields != nil {
		q.Set("fields", strings.Join(*options.Fields, " "))
	}

	u.RawQuery = q.Encode()

	httpResponse, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var r tafs.Response
	err = xml.Unmarshal(data, &r)
	if err != nil {
		return nil, err

	}

	return &r, nil
}
