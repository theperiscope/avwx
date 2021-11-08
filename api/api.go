package api

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

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

type MetarOptions struct {
	Stations                 []string
	StartTime                timeValue
	EndTime                  timeValue
	HoursBeforeNow           int32
	MostRecent               bool
	MostRecentForEachStation bool
	MinLat                   float64
	MaxLat                   float64
	MinLon                   float64
	MaxLon                   float64
	RadialDistance           string
	FlightPath               []string
	MinDegreeDistance        float64
	Fields                   []string
}

type TafOptions struct {
	Stations                 []string
	StartTime                timeValue
	EndTime                  timeValue
	TimeType                 string // only difference from MetarOptions
	HoursBeforeNow           int32
	MostRecent               bool
	MostRecentForEachStation bool
	MinLat                   float64
	MaxLat                   float64
	MinLon                   float64
	MaxLon                   float64
	RadialDistance           string
	FlightPath               []string
	MinDegreeDistance        float64
	Fields                   []string
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

	if len(options.Stations) > 0 {
		q.Set("stationString", strings.Join(options.Stations, " "))
	}
	if !time.Time(options.StartTime).IsZero() {
		q.Set("startTime", options.StartTime.String())
	}
	if !time.Time(options.EndTime).IsZero() {
		q.Set("endTime", options.EndTime.String())
	}
	if options.HoursBeforeNow > 0 {
		q.Set("hoursBeforeNow", strconv.FormatInt(int64(options.HoursBeforeNow), 10))
	}
	if options.MostRecent {
		q.Set("mostRecent", strconv.FormatBool(options.MostRecent))
	}
	if options.MostRecentForEachStation {
		q.Set("mostRecentForEachStation", strconv.FormatBool(options.MostRecentForEachStation))
	}
	if options.MinLat != 0 {
		q.Set("minLat", strconv.FormatFloat(options.MinLat, 'f', -1, 64))
	}
	if options.MaxLat != 0 {
		q.Set("maxLat", strconv.FormatFloat(options.MaxLat, 'f', -1, 64))
	}
	if options.MinLon != 0 {
		q.Set("minLon", strconv.FormatFloat(options.MinLon, 'f', -1, 64))
	}
	if options.MaxLon != 0 {
		q.Set("maxLon", strconv.FormatFloat(options.MaxLon, 'f', -1, 64))
	}
	if len(options.RadialDistance) > 0 {
		q.Set("radialDistance", options.RadialDistance)
	}
	if len(options.FlightPath) > 0 {
		q.Set("flightPath", strings.Join(options.FlightPath, " "))
	}
	if options.MinDegreeDistance != 0 {
		q.Set("minDegreeDistance", strconv.FormatFloat(options.MinDegreeDistance, 'f', -1, 64))
	}
	if len(options.Fields) > 0 {
		q.Set("fields", strings.Join(options.Fields, " "))
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

	if len(options.Stations) > 0 {
		q.Set("stationString", strings.Join(options.Stations, " "))
	}
	if !time.Time(options.StartTime).IsZero() {
		q.Set("startTime", options.StartTime.String())
	}
	if !time.Time(options.EndTime).IsZero() {
		q.Set("endTime", options.EndTime.String())
	}
	if len(options.TimeType) > 0 {
		q.Set("timeType", options.TimeType)
	}
	if options.HoursBeforeNow > 0 {
		q.Set("hoursBeforeNow", strconv.FormatInt(int64(options.HoursBeforeNow), 10))
	}
	if options.MostRecent {
		q.Set("mostRecent", strconv.FormatBool(options.MostRecent))
	}
	if options.MostRecentForEachStation {
		q.Set("mostRecentForEachStation", strconv.FormatBool(options.MostRecentForEachStation))
	}
	if options.MinLat != 0 {
		q.Set("minLat", strconv.FormatFloat(options.MinLat, 'f', -1, 64))
	}
	if options.MaxLat != 0 {
		q.Set("maxLat", strconv.FormatFloat(options.MaxLat, 'f', -1, 64))
	}
	if options.MinLon != 0 {
		q.Set("minLon", strconv.FormatFloat(options.MinLon, 'f', -1, 64))
	}
	if options.MaxLon != 0 {
		q.Set("maxLon", strconv.FormatFloat(options.MaxLon, 'f', -1, 64))
	}
	if len(options.RadialDistance) > 0 {
		q.Set("radialDistance", options.RadialDistance)
	}
	if len(options.FlightPath) > 0 {
		q.Set("flightPath", strings.Join(options.FlightPath, " "))
	}
	if options.MinDegreeDistance != 0 {
		q.Set("minDegreeDistance", strconv.FormatFloat(options.MinDegreeDistance, 'f', -1, 64))
	}
	if len(options.Fields) > 0 {
		q.Set("fields", strings.Join(options.Fields, " "))
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
