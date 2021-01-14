package tafs

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
	XMLName xml.Name `xml:"data"`
	Tafs    []xmlTaf `xml:"TAF"`
}

type xmlTaf struct {
	XMLName xml.Name `xml:"TAF"`
	RawText string   `xml:"raw_text"`
}

// GetData returns TAF data from Aviation Weather Center's Text Data Server.
// Sample stations: KORD PH* @ny ~us
func GetData(stations []string) ([]string, error) {

	u, err := url.Parse(apiEndPoint)
	if err != nil {
		return nil, err
	}

	s := strings.Join(stations, " ")

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

	var r xmlResponse
	err = xml.Unmarshal(data, &r)
	if err != nil {
		return nil, err

	}

	var result []string

	for _, taf := range r.Data.Tafs {
		result = append(result, taf.RawText)
	}

	return result, nil
}
