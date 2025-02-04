package main

import (
	"fmt"
	"geokml/utils"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	geohashes, err := utils.ExtractGeohashesFromKML("example.kml", 8)
	if err != nil {
		log.Fatalf("Error extracting geohashes: %v", err)
	}

	fmt.Println(len(geohashes))
}
