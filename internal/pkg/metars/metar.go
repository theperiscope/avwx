package metars

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const apiEndPoint = "https://aviationweather.gov/adds/dataserver_current/httpparam"

type xmlResponse struct {
	XMLName xml.Name `xml:"response"`
	Data    xmlData  `xml:"data"`
}

type xmlData struct {
	XMLName xml.Name   `xml:"data"`
	Metars  []xmlMetar `xml:"METAR"`
}

type xmlMetar struct {
	XMLName xml.Name `xml:"METAR"`
	RawText string   `xml:"raw_text"`
}

// GetData returns METAR data from Aviation Weather Center's Text Data Server.
// Sample stations: KORD PH* @ny ~us
func GetData(stations []string) ([]string, error) {

	u, err := url.Parse(apiEndPoint)
	if err != nil {
		return nil, err
	}

	s := strings.Join(stations, " ")

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

	var r xmlResponse
	err = xml.Unmarshal(data, &r)
	if err != nil {
		return nil, err

	}

	var result []string

	for _, metar := range r.Data.Metars {
		result = append(result, metar.RawText)
	}

	return result, nil
}
