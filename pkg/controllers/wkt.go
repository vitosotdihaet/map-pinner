package controllers

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)


func pointsToWKT(points []entities.Point) []string {
	length := len(points)
	if length == 0 {
		return make([]string, 0)
	}

	postgisPoints := make([]string, length+1)
	for i, point := range points {
		postgisPoints[i] = fmt.Sprintf("ST_MakePoint(%f, %f)", point.Latitude, point.Longitude)
	}

	postgisPoints[length] = fmt.Sprintf("ST_MakePoint(%f, %f)", points[0].Latitude, points[0].Longitude)

	return postgisPoints
}

func parseWKT(wkt string) []entities.Point {
	// Remove the "POLYGON((" prefix and "))" suffix
	wkt = strings.TrimPrefix(wkt, "POLYGON((")
	wkt = strings.TrimSuffix(wkt, "))")

	// Split the coordinates
	coordPairs := strings.Split(wkt, ",")

	points := make([]entities.Point, 0)
	for _, pair := range coordPairs {
		coords := strings.Fields(pair)

		if len(coords) != 2 {
			logrus.Errorf("invalid wkt polygon parsing: %s", wkt)
			break
		}

		var point entities.Point
		fmt.Sscanf(coords[0], "%f", &point.Longitude)
		fmt.Sscanf(coords[1], "%f", &point.Latitude)
		points = append(points, point)
	}

	return points
}