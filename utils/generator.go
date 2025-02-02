package utils

import "geokml/structs"
import "github.com/flywave/go-earcut"
import "github.com/mmcloughlin/geohash"

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

// generateGeohashesForPolygon generates geohashes that cover the entire polygon area
func generateGeohashesForSinglePolygon(polygon *structs.Polygon, precision uint) map[string]struct{} {
	expandedHashes := make(map[string]struct{})
	var queue []string

	interiorPoint := getInteriorPointByTriangulation(polygon.Coordinates)

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

// getInteriorPointByTriangulation finds a valid point inside the polygon using triangulation
func getInteriorPointByTriangulation(polygon []structs.Coordinate) structs.Coordinate {
	// Prepare input for earcut: convert to a flat array of [x, y] points
	var flatCoords []float64
	for _, coord := range polygon {
		flatCoords = append(flatCoords, coord.Lng, coord.Lat) // [lng1, lat1, lng2, lat2, ...]
	}

	// Perform triangulation using earcut
	triangles, err := earcut.Earcut(flatCoords, nil, 2) // 2 indicates [x, y] pairs
	if err != nil {
		return polygon[0] // Fallback to the first vertex
	}
	// Iterate through the triangles and find the centroid of each
	for i := 0; i < len(triangles); i += 3 {
		// Get vertices of the triangle
		aIndex, bIndex, cIndex := triangles[i]*2, triangles[i+1]*2, triangles[i+2]*2
		a := structs.Coordinate{Lng: flatCoords[aIndex], Lat: flatCoords[aIndex+1]}
		b := structs.Coordinate{Lng: flatCoords[bIndex], Lat: flatCoords[bIndex+1]}
		c := structs.Coordinate{Lng: flatCoords[cIndex], Lat: flatCoords[cIndex+1]}

		// Calculate the centroid of the triangle
		centroidLat := (a.Lat + b.Lat + c.Lat) / 3
		centroidLng := (a.Lng + b.Lng + c.Lng) / 3
		centroid := structs.Coordinate{Lat: centroidLat, Lng: centroidLng}

		// Check if the centroid is inside the polygon
		if pointInPolygon(centroid, polygon) {
			return centroid // Found a valid interior point
		}
	}

	// Fallback to the first vertex if no valid point is found
	return polygon[0]
}
