package utils

import (
	"geokml/structs"
	"github.com/mmcloughlin/geohash"
)

func contains(p *structs.Polygon, bounds geohash.Box) bool {
	for _, point := range []structs.Coordinate{
		{bounds.MinLat, bounds.MinLng},
		{bounds.MinLat, bounds.MaxLng},
		{bounds.MaxLat, bounds.MinLng},
		{bounds.MaxLat, bounds.MaxLng},
	} {
		if !PointInPolygon(point, p.Coordinates) {
			return false
		}
	}
	return true
}

func PointInPolygon(point structs.Coordinate, polygon []structs.Coordinate) bool {
	n := len(polygon)
	inside := false

	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		if ((polygon[i].Lat > point.Lat) != (polygon[j].Lat > point.Lat)) &&
			(point.Lng < (polygon[j].Lng-polygon[i].Lng)*(point.Lat-polygon[i].Lat)/(polygon[j].Lat-polygon[i].Lat)+polygon[i].Lng) {
			inside = !inside
		}
	}

	return inside
}
