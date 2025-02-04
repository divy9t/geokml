package utils

import (
	"github.com/divy9t/geokml/structs"
	"strconv"
	"strings"
)

// ExtractGeohashesFromKML parses the given KML file at kmlPath,
// extracts coordinate data from polygons, and generates geohashes
// with the specified precision.
//
// kmlPath: The path to the KML file.
// precision: The desired geohash precision.
// Returns: A set of geohashes covering the KML-defined regions, or an error if parsing fails.
func ExtractGeohashesFromKML(kmlPath string, precision uint) (map[string]struct{}, error) {
	// Extract coordinates nodes
	nodes, err := getNodesFromKmlFile(kmlPath)
	if err != nil {
		return nil, err
	}

	// Create polygons from parsed coordinates
	var polygons []structs.Polygon
	for _, node := range nodes {
		coords := ParseCoordinates(node.InnerText())
		if len(coords) < 3 {
			continue
		}
		polygon := structs.Polygon{Coordinates: coords}
		polygons = append(polygons, polygon)
	}
	// Generate geohashes for all polygons
	geohashes := GenerateGeohashesForPolygon(polygons, precision)

	return geohashes, nil
}

// ParseCoordinates parses the KML coordinates string into a slice of Coordinates
func ParseCoordinates(coordString string) []structs.Coordinate {
	var coordinates []structs.Coordinate
	// Split by spaces (coordinate pairs)
	pairs := strings.Fields(coordString)
	for _, pair := range pairs {
		// Split by commas (lon,lat[,alt])
		values := strings.Split(pair, ",")
		if len(values) < 2 {
			continue
		}
		lng, err1 := strconv.ParseFloat(values[0], 64)
		lat, err2 := strconv.ParseFloat(values[1], 64)
		if err1 == nil && err2 == nil {
			coordinates = append(coordinates, structs.Coordinate{Lat: lat, Lng: lng})
		}
	}
	return coordinates
}
