package structs

type Polygon struct {
	OuterBoundaryIs OuterBoundaryIs `xml:"outerBoundaryIs"`
	Coordinates     []Coordinate
}
