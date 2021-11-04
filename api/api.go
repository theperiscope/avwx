package api

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/theperiscope/avwx/metars"
	"github.com/theperiscope/avwx/tafs"
)

const DefaultApiEndPoint = "https://aviationweather.gov/adds/dataserver_current/httpparam"

type Client struct {
	c           *http.Client
	ApiEndPoint string
}

func NewClient(apiEndPoint string) *Client {
	c := &http.Client{}

	return &Client{
		c:           c,
		ApiEndPoint: apiEndPoint,
	}
}

func (c *Client) GetMetar(stations []string) ([]metars.Metar, error) {
	s := strings.Join(stations, " ")

	u, err := url.Parse(c.ApiEndPoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("dataSource", "metars")
	q.Set("requestType", "retrieve")
	q.Set("format", "xml")
	q.Set("hoursBeforeNow", "3")
	q.Set("mostRecentForEachStation", "true")
	q.Set("stationString", s)
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

	return r.Data.Metars, nil
}

func (c *Client) GetTaf(stations []string) ([]tafs.Taf, error) {
	s := strings.Join(stations, " ")

	u, err := url.Parse(c.ApiEndPoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("dataSource", "tafs")
	q.Set("requestType", "retrieve")
	q.Set("format", "xml")
	q.Set("hoursBeforeNow", "9")
	q.Set("mostRecentForEachStation", "true")
	q.Set("stationString", s)
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

	return r.Data.Tafs, nil
}
