package utils

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"os"
)

func getNodesFromKmlFile(kmlPath string) ([]*xmlquery.Node, error) {
	file, err := os.Open(kmlPath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	doc, err := xmlquery.Parse(file)
	if err != nil {
		return nil, err
	}

	nodes := xmlquery.Find(doc, "//coordinates")
	if len(nodes) == 0 {
		return nil, fmt.Errorf("no coordinates found in KML")
	}

	return nodes, nil

}

func getPolygonFromNodes(nodes []*xmlquery.Node) []string {
	var polygons []string
	for _, node := range nodes {
		polygons = append(polygons, node.InnerText())
	}
	return polygons

}
