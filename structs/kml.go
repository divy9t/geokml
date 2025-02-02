package structs

import "encoding/xml"

type KML struct {
	XMLName  xml.Name  `xml:"kml"`
	Document *Document `xml:"Document"`
}
