package utils

import (
	"github.com/divy9t/geokml/structs"
	"math"
)
import "github.com/flywave/go-earcut"
import "github.com/mmcloughlin/geohash"

// GenerateGeohashesForPolygon generates geohashes that cover the entire polygon area
func GenerateGeohashesForPolygon(polygons []structs.Polygon, precision uint) map[string]struct{} {
	geohashes := make(map[string]struct{})
	for _, polygon := range polygons {
		polygonGeohashes := generateGeohashesForSinglePolygon(&polygon, precision)
		for hash := range polygonGeohashes {
			geohashes[hash] = struct{}{}
		}
	}
	return geohashes

}

// GenerateGeohashesForSinglePolygon generates geohashes that cover the entire area of a single polygon
func generateGeohashesForSinglePolygon(polygon *structs.Polygon, precision uint) map[string]struct{} {
	expandedHashes := make(map[string]struct{})
	var queue []string

	interiorPoint := GetInteriorPointByTriangulation(polygon.Coordinates)

	initialGeohash := geohash.EncodeWithPrecision(interiorPoint.Lat, interiorPoint.Lng, precision)
	queue = append(queue, initialGeohash)

	// BFS to generate geohashes
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Skip already processed geohashes
		if _, exists := expandedHashes[current]; exists {
			continue
		}

		// Decode geohash and check containment
		bounds := geohash.BoundingBox(current)
		if contains(polygon, bounds) {
			expandedHashes[current] = struct{}{}

			// Add neighbors
			for _, neighbor := range geohash.Neighbors(current) {
				if _, exists := expandedHashes[neighbor]; !exists {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return expandedHashes
}

func triangleArea(a, b, c structs.Coordinate) float64 {
	return math.Abs((a.Lng*(b.Lat-c.Lat) + b.Lng*(c.Lat-a.Lat) + c.Lng*(a.Lat-b.Lat)) / 2)
}

// GetInteriorPointByTriangulation returns an interior point of the polygon using triangulation.
func GetInteriorPointByTriangulation(polygon []structs.Coordinate) structs.Coordinate {
	// Prepare input for earcut: convert to a flat array of [x, y] points.
	var flatCoords []float64
	for _, coord := range polygon {
		flatCoords = append(flatCoords, coord.Lng, coord.Lat) // [lng1, lat1, lng2, lat2, ...]
	}

	// Perform triangulation using earcut.
	triangles, err := earcut.Earcut(flatCoords, nil, 2) // 2 indicates [x, y] pairs
	if err != nil || len(triangles) < 3 {
		// Fallback if triangulation fails.
		return polygon[0]
	}

	var bestCentroid structs.Coordinate
	maxArea := -1.0
	found := false

	// Iterate through the triangles.
	for i := 0; i < len(triangles); i += 3 {
		// Get vertices of the triangle.
		aIndex := triangles[i] * 2
		bIndex := triangles[i+1] * 2
		cIndex := triangles[i+2] * 2
		a := structs.Coordinate{Lng: flatCoords[aIndex], Lat: flatCoords[aIndex+1]}
		b := structs.Coordinate{Lng: flatCoords[bIndex], Lat: flatCoords[bIndex+1]}
		c := structs.Coordinate{Lng: flatCoords[cIndex], Lat: flatCoords[cIndex+1]}

		// Calculate the centroid of the triangle.
		centroid := structs.Coordinate{
			Lat: (a.Lat + b.Lat + c.Lat) / 3,
			Lng: (a.Lng + b.Lng + c.Lng) / 3,
		}

		// Check if the centroid is inside the polygon.
		if PointInPolygon(centroid, polygon) {
			area := triangleArea(a, b, c)
			if area > maxArea {
				maxArea = area
				bestCentroid = centroid
				found = true
			}
		}
	}

	// If a valid centroid was found from the largest triangle, return it.
	if found {
		return bestCentroid
	}

	// Fallback: return the first vertex if no valid interior point was found.
	return polygon[0]
}
