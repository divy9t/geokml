package utils

import (
	"geokml/structs"
	"strconv"
	"strings"
)

func ExtractGeohashesFromKML(kmlPath string, precision uint) (map[string]struct{}, error) {
	// Extract coordinates nodes
	nodes, err := getNodesFromKmlFile(kmlPath)
	if err != nil {
		return nil, err
	}

	// Create polygons from parsed coordinates
	var polygons []structs.Polygon
	for _, node := range nodes {
		coords := parseCoordinates(node.InnerText())
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

// parseCoordinates parses the KML coordinates string into a slice of Coordinates
func parseCoordinates(coordString string) []structs.Coordinate {
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
