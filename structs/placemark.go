package structs

type Placemark struct {
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
	Polygon     *Polygon `xml:"Polygon"`
}
