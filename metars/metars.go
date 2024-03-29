package metars

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

type Response struct {
	XMLName      xml.Name   `xml:"response" json:"-"`
	Version      string     `xml:"version,attr"`
	RequestIndex int32      `xml:"request_index"`
	Errors       []string   `xml:"errors>error"`
	Warnings     []string   `xml:"warnings>warning"`
	TimeTakenMs  int32      `xml:"time_taken_ms"`
	DataSource   DataSource `xml:"data_source"`
	Request      Request    `xml:"request"`
	Data         Data       `xml:"data"`
}

type Request struct {
	XMLName xml.Name `xml:"request" json:"-"`
	Type    string   `xml:"type,attr"`
}

type DataSource struct {
	XMLName xml.Name `xml:"data_source" json:"-"`
	Name    string   `xml:"name,attr"`
}

type Data struct {
	XMLName    xml.Name `xml:"data" json:"-"`
	NumResults int32    `xml:"num_results,attr"`
	Metars     []Metar  `xml:"METAR"`
}

type SkyCondition struct {
	XMLName        xml.Name `xml:"sky_condition" json:"-"`
	SkyCover       string   `xml:"sky_cover,attr"`
	CloudBaseFtAGL int32    `xml:"cloud_base_ft_agl,attr"`
}

type QualityControlFlags struct {
	XMLName     xml.Name `xml:"quality_control_flags" json:"-"`
	AutoStation bool     `xml:"auto_station"`
}

type Metar struct {
	XMLName                   xml.Name            `xml:"METAR" json:"-"`
	RawText                   string              `xml:"raw_text"`
	StationId                 string              `xml:"station_id"`
	ObservationTime           time.Time           `xml:"observation_time"`
	Latitude                  float64             `xml:"latitude"`
	Longitude                 float64             `xml:"longitude"`
	TempC                     float64             `xml:"temp_c"`
	DewpointC                 float64             `xml:"dewpoint_c"`
	WindDirDegrees            int32               `xml:"wind_dir_degrees"`
	WindSpeedKt               int32               `xml:"wind_speed_kt"`
	WindGustKt                int32               `xml:"wind_gust_kt"`
	VisibilityStatuteMi       float64             `xml:"visibility_statute_mi"`
	AltimInHg                 float64             `xml:"altim_in_hg"`
	SeaLevelPressureMb        float64             `xml:"sea_level_pressure_mb"`
	QualityControlFlags       QualityControlFlags `xml:"quality_control_flags"`
	WxString                  string              `xml:"wx_string"`
	SkyCondition              []SkyCondition      `xml:"sky_condition"`
	FlightCategory            string              `xml:"flight_category"`
	ThreeHrPressureTendencyMb float64             `xml:"three_hr_pressure_tendency_mb"`
	MaxTC                     float64             `xml:"maxT_c"`
	MinTC                     float64             `xml:"minT_c"`
	MaxT24hrC                 float64             `xml:"maxT24hr_c"`
	MinT24hrC                 float64             `xml:"minT24hr_c"`
	PrecipIn                  float64             `xml:"precip_in"`
	Pcp3hrIn                  float64             `xml:"pcp3hr_in"`
	Pcp6hrIn                  float64             `xml:"pcp6hr_in"`
	Pcp24hrIn                 float64             `xml:"pcp24hr_in"`
	SnowIn                    float64             `xml:"snow_in"`
	VertVisFt                 int32               `xml:"vert_vis_ft"`
	MetarType                 string              `xml:"metar_type"`
	ElevationM                float64             `xml:"elevation_m"`
}

func (r *Response) ToRawTextOnly() (s []string) {
	for _, metar := range r.Data.Metars {
		s = append(s, metar.RawText)
	}
	return
}

func (r *Response) ToJson() (s string, err error) {
	bytes, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	s = string(bytes)
	return
}

func (r *Response) ToJsonIndented() (s string, err error) {
	bytes, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}

	s = string(bytes)
	return
}
