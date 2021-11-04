package tafs

import (
	"encoding/xml"

	"github.com/relvacode/iso8601"
)

type Response struct {
	XMLName      xml.Name   `xml:"response"`
	Version      string     `xml:"version,attr"`
	RequestIndex int32      `xml:"request_index"`
	Errors       []Error    `xml:"errors>error"`
	Warnings     []Warning  `xml:"warnings>warning"`
	TimeTakenMs  int32      `xml:"time_taken_ms"`
	DataSource   DataSource `xml:"data_source"`
	Request      Request    `xml:"request"`
	Data         Data       `xml:"data"`
}

type Error struct {
	XMLName xml.Name `xml:"error"`
	Error   string   `xml:",innerXml"`
}

type Warning struct {
	XMLName xml.Name `xml:"warning"`
	Warning string   `xml:",innerXml"`
}

type Request struct {
	Type string `xml:"type,attr"`
}

type DataSource struct {
	Name string `xml:"name,attr"`
}

type Data struct {
	XMLName    xml.Name `xml:"data"`
	NumResults int32    `xml:"num_results,attr"`
	Tafs       []Taf    `xml:"TAF"`
}

type SkyCondition struct {
	XMLName        xml.Name `xml:"sky_condition"`
	SkyCover       string   `xml:"sky_cover,attr"`
	CloudBaseFtAGL int32    `xml:"cloud_base_ft_agl,attr"`
	CloudType      string   `xml:"cloud_type,attr"`
}

type Forecast struct {
	XMLName             xml.Name     `xml:"forecast"`
	FcstTimeFrom        iso8601.Time `xml:"fcst_time_from"`
	FcstTimeTo          iso8601.Time `xml:"fcst_time_to"`
	ChangeIndicator     string       `xml:"change_indicator"`
	TimeBecoming        iso8601.Time `xml:"time_becoming"`
	Probability         int32        `xml:"probability"`
	WindDirDegrees      int32        `xml:"wind_dir_degrees"`
	WindSpeedKt         int32        `xml:"wind_speed_kt"`
	WindGustKt          int32        `xml:"wind_gust_kt"`
	WindShearHgtFtAgl   int32        `xml:"wind_shear_hgt_ft_agl"`
	WindShearDirDegrees int32        `xml:"wind_shear_dir_degrees"`
	WindShearSpeedKt    float64      `xml:"wind_shear_speed_kt"`
	VisibilityStatuteMi float64      `xml:"visibility_statute_mi"`
	AltimInHg           int32        `xml:"altim_in_hg"`
	VertVisFt           string       `xml:"vert_vis_ft"`
	WxString            string       `xml:"wx_string"`
	NotDecoded          string       `xml:"not_decoded"`
	SkyCondition        string       `xml:"sky_condition"`
	TurbulenceCondition string       `xml:"turbulence_condition"`
	IcingCondition      string       `xml:"icing_condition"`
	Temperature         []int32      `xml:"temperature"`
	ValidTime           string       `xml:"valid_time"`
	SfcTempC            float64      `xml:"sfc_temp_c"`
	MaxTempC            float64      `xml:"max_temp_c"`
	MinTempC            float64      `xml:"min_temp_c"`
}

type Taf struct {
	XMLName       xml.Name     `xml:"TAF"`
	RawText       string       `xml:"raw_text"`
	StationId     string       `xml:"station_id"`
	IssueTime     iso8601.Time `xml:"issue_time"`
	BulletinTime  iso8601.Time `xml:"bulletin_time"`
	ValidTimeFrom iso8601.Time `xml:"valid_time_from"`
	ValidTimeTo   iso8601.Time `xml:"valid_time_to"`
	Remarks       string       `xml:"remarks"`
	Latitude      float64      `xml:"latitude"`
	Longitude     float64      `xml:"longitude"`
	ElevationM    float64      `xml:"elevation_m"`
	Forecast      []Forecast   `xml:"forecast"`
}
