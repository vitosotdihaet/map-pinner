package controllers

import (
	"fmt"
	"strings"

	"github.com/vitosotdihaet/map-pinner/pkg/entities"
)


func pointsToWKT(points []entities.Point) []string {
	length := len(points)
	if length == 0 {
		return make([]string, 0)
	}

	postgisPoints := make([]string, length)
	for i, point := range points {
		postgisPoints[i] = fmt.Sprintf("ST_MakePoint(%f, %f)", point.Longitude, point.Latitude)
	}

	return postgisPoints
}

func parsedWKTToPoints(coordinatePairs []string) []entities.Point {
	points := make([]entities.Point, 0)
	for _, pair := range coordinatePairs {
		coords := strings.Fields(pair)

		var point entities.Point
		fmt.Sscanf(coords[0], "%f", &point.Longitude)
		fmt.Sscanf(coords[1], "%f", &point.Latitude)
		points = append(points, point)
	}

	return points

}

func parseWKTPolygon(wkt string) []entities.Point {
	wkt = strings.TrimPrefix(wkt, "POLYGON((")
	wkt = strings.TrimSuffix(wkt, ")")
	return parsedWKTToPoints(strings.Split(wkt, ","))
}

func parseWKTGraph(wkt string) []entities.Point {
	wkt = strings.TrimPrefix(wkt, "LINESTRING(")
	wkt = strings.TrimSuffix(wkt, ")")
	return parsedWKTToPoints(strings.Split(wkt, ","))
}