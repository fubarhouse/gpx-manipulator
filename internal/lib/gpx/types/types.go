package types

import (
	"encoding/xml"
	"time"
)

type FixType string

const (
	FixNone FixType = "none"
	Fix2D   FixType = "2d"
	Fix3D   FixType = "3d"
	FixDGPS FixType = "dgps"
	FixPPS  FixType = "pps"
)

type Link struct {
	Href string `xml:"href,attr"`
	Text string `xml:"text,omitempty"`
	Type string `xml:"type,omitempty"`
}

type RawExtensions struct {
	InnerXML []byte `xml:",innerxml"`
}

type GPX struct {
	XMLName xml.Name   `xml:"gpx"`
	Creator string     `xml:"creator,attr"`
	Attr    []xml.Attr `xml:",any,attr"`

	Metadata   *Metadata      `xml:"metadata,omitempty"`
	Waypoints  []Wpt          `xml:"wpt,omitempty"`
	Routes     []Rte          `xml:"rte,omitempty"`
	Tracks     []Track        `xml:"trk,omitempty"`
	Extensions *RawExtensions `xml:"extensions,omitempty"`
}

type Metadata struct {
	Name       *string        `xml:"name,omitempty"`
	Desc       *string        `xml:"desc,omitempty"`
	Author     *Person        `xml:"author,omitempty"`
	Link       []Link         `xml:"link,omitempty"`
	Time       *time.Time     `xml:"time,omitempty"`
	Keywords   *string        `xml:"keywords,omitempty"`
	Extensions *RawExtensions `xml:"extensions,omitempty"`
}

type Person struct {
	Name  *string `xml:"name,omitempty"`
	Email *Email  `xml:"email,omitempty"`
	Link  []Link  `xml:"link,omitempty"`
}

type Email struct {
	ID     string `xml:"id,attr"`
	Domain string `xml:"domain,attr"`
}

type Wpt struct {
	XMLName xml.Name `xml:"wpt"`

	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`

	Elevation     *float64       `xml:"ele,omitempty"`
	Time          *time.Time     `xml:"time,omitempty"`
	MagVar        *float64       `xml:"magvar,omitempty"`
	GeoIDHeight   *float64       `xml:"geoidheight,omitempty"`
	Name          *string        `xml:"name,omitempty"`
	Cmt           *string        `xml:"cmt,omitempty"`
	Desc          *string        `xml:"desc,omitempty"`
	Src           *string        `xml:"src,omitempty"`
	Link          []Link         `xml:"link,omitempty"`
	Sym           *string        `xml:"sym,omitempty"`
	Type          *string        `xml:"type,omitempty"`
	Fix           *FixType       `xml:"fix,omitempty"`
	Sat           *uint          `xml:"sat,omitempty"`
	HDOP          *float64       `xml:"hdop,omitempty"`
	VDOP          *float64       `xml:"vdop,omitempty"`
	PDOP          *float64       `xml:"pdop,omitempty"`
	AgeOfDGPSData *float64       `xml:"ageofdgpsdata,omitempty"`
	DGPSID        *uint          `xml:"dgpsid,omitempty"`
	Extensions    *RawExtensions `xml:"extensions,omitempty"`
}

type Rte struct {
	XMLName xml.Name `xml:"rte"`

	Name       *string        `xml:"name,omitempty"`
	Cmt        *string        `xml:"cmt,omitempty"`
	Desc       *string        `xml:"desc,omitempty"`
	Src        *string        `xml:"src,omitempty"`
	Link       []Link         `xml:"link,omitempty"`
	Number     *uint          `xml:"number,omitempty"`
	Type       *string        `xml:"type,omitempty"`
	Extensions *RawExtensions `xml:"extensions,omitempty"`

	Rtepts []Rtept `xml:"rtept"`
}

type Rtept struct {
	XMLName xml.Name `xml:"rtept"`

	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`

	Elevation     *float64       `xml:"ele,omitempty"`
	Time          *time.Time     `xml:"time,omitempty"`
	MagVar        *float64       `xml:"magvar,omitempty"`
	GeoIDHeight   *float64       `xml:"geoidheight,omitempty"`
	Name          *string        `xml:"name,omitempty"`
	Cmt           *string        `xml:"cmt,omitempty"`
	Desc          *string        `xml:"desc,omitempty"`
	Src           *string        `xml:"src,omitempty"`
	Link          []Link         `xml:"link,omitempty"`
	Sym           *string        `xml:"sym,omitempty"`
	Type          *string        `xml:"type,omitempty"`
	Fix           *FixType       `xml:"fix,omitempty"`
	Sat           *uint          `xml:"sat,omitempty"`
	HDOP          *float64       `xml:"hdop,omitempty"`
	VDOP          *float64       `xml:"vdop,omitempty"`
	PDOP          *float64       `xml:"pdop,omitempty"`
	AgeOfDGPSData *float64       `xml:"ageofdgpsdata,omitempty"`
	DGPSID        *uint          `xml:"dgpsid,omitempty"`
	Extensions    *RawExtensions `xml:"extensions,omitempty"`
}

type Track struct {
	XMLName xml.Name `xml:"trk"`

	Name       string         `xml:"name,omitempty"`
	Extensions *RawExtensions `xml:"extensions,omitempty"`
	Segs       []TrkSeg       `xml:"trkseg"`
}

type TrkSeg struct {
	Points []TrkPt `xml:"trkpt"`
}

type TrkPt struct {
	XMLName xml.Name `xml:"trkpt"`

	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`

	Elevation *float64 `xml:"ele,omitempty"`
	Time      *string  `xml:"time,omitempty"`

	MagVar        *float64       `xml:"magvar,omitempty"`
	GeoIDHeight   *float64       `xml:"geoidheight,omitempty"`
	Name          *string        `xml:"name,omitempty"`
	Cmt           *string        `xml:"cmt,omitempty"`
	Desc          *string        `xml:"desc,omitempty"`
	Src           *string        `xml:"src,omitempty"`
	Link          []Link         `xml:"link,omitempty"`
	Sym           *string        `xml:"sym,omitempty"`
	Type          *string        `xml:"type,omitempty"`
	Fix           *FixType       `xml:"fix,omitempty"`
	Sat           *uint          `xml:"sat,omitempty"`
	HDOP          *float64       `xml:"hdop,omitempty"`
	VDOP          *float64       `xml:"vdop,omitempty"`
	PDOP          *float64       `xml:"pdop,omitempty"`
	AgeOfDGPSData *float64       `xml:"ageofdgpsdata,omitempty"`
	DGPSID        *uint          `xml:"dgpsid,omitempty"`
	Extensions    *RawExtensions `xml:"extensions,omitempty"`
}
