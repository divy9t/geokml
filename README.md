# GeoKML

GeoKML is an open source Go library for extracting geohashes from KML files. It parses KML files, extracts coordinate data, and generates geohashes that cover the polygon areas defined in the KML. The library also provides functions to verify whether a geohash’s bounding box is completely contained within a polygon.

## Features

- Parse KML files to extract coordinate nodes.
- Generate geohashes for entire polygon areas.
- Use triangulation to find an interior point for geohash generation.
- Validate polygon containment with a point-in-polygon algorithm.
- Minimal dependencies and easy to integrate into your projects.

## Installation

To install GeoKML, use `go get`:

```bash
  go get github.com/divy9t/geokml
```
