# GeoKML 🚀

**GeoKML** is an open-source Go library designed to extract **geohashes** from **KML (Keyhole Markup Language)** files. The library simplifies the process of converting geographic polygons into geohashes, making it ideal for applications in mapping, geospatial data processing, and location-based services.

GeoKML handles the complex tasks of parsing KML files, generating geohashes for entire polygon regions, and validating geohash containment using triangulation and point-in-polygon algorithms.

> ⚠ **Note:** This library is under continuous development and may not cover all edge cases perfectly. Community feedback and contributions are welcome as we improve its robustness and usability!

---

## 🌟 Features

- **KML Parsing**: Automatically parses KML files and extracts coordinate nodes for polygon regions.
- **Geohash Generation**: Efficiently generates geohashes that cover the entirety of the polygon area.
- **Robust Interior Point Calculation**: Uses triangulation to determine a reliable starting point for geohash generation.
- **Polygon Containment Validation**: Implements a point-in-polygon algorithm to ensure that geohashes accurately represent the region.
- **Minimal Dependencies**: Lightweight and easy to integrate into your Go projects.

---

## 📦 Installation

To install GeoKML, use:

```bash 
  go get github.com/divy9t/geokml
```

---

## 🚀 Usage Example

Using GeoKML is as simple as providing the path to your KML file and specifying the desired geohash precision. The library will handle the parsing, polygon conversion, and geohash generation for you.

```go
package main

import (
    "fmt"
    "log"

    "github.com/divy9t/geokml/pkg/utils"
)

func main() {
    // Provide the KML file path and desired geohash precision
    kmlPath := "path/to/your/file.kml"
    precision := uint(6)

    // Extract geohashes from the KML file
    geohashes, err := utils.ExtractGeohashesFromKML(kmlPath, precision)
    if err != nil {
        log.Fatalf("Error extracting geohashes: %v", err)
    }

    // Display the generated geohashes
    fmt.Println("Extracted Geohashes:")
    for gh := range geohashes {
        fmt.Println(gh)
    }
}
```

---

## 📚 Documentation

You can find the full documentation on [pkg.go.dev](https://pkg.go.dev/github.com/divy9t/geokml), where each function and type is documented using GoDoc.

---

## 🤝 Contributing

This library is under active development, and contributions are highly welcome! If you encounter any issues or have suggestions for improvement, please [open an issue](https://github.com/divy9t/geokml/issues) or submit a pull request.

### Ways to Contribute:
- Submit bug reports or feature requests.
- Fix bugs or contribute new features.
- Improve documentation or add examples.

---

## ⚖️ License

GeoKML is licensed under the [MIT License](./LICENSE). You’re free to use it in personal and commercial projects, but contributions are always appreciated!
