package entities

import (
	"encoding/json"
)


const PolygonMaxPoints = 20

type Polygon struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Points [PolygonMaxPoints]*Point
	Length int
}

func (polygon *Polygon) UnmarshalJSON(data []byte) error {
	var rawPoints []map[string]interface{}
	if err := json.Unmarshal(data, &rawPoints); err != nil {
		return err
	}

	for i, rawPoint := range rawPoints {
		point := &Point{}

		if id, ok := rawPoint["id"].(float64); ok {
			point.ID = uint64(id)
		}
		if name, ok := rawPoint["name"].(string); ok {
			point.Name = name
		}
		if latitude, ok := rawPoint["latitude"].(float64); ok {
			point.Latitude = latitude
		}
		if longitude, ok := rawPoint["longitude"].(float64); ok {
			point.Longitude = longitude
		}

		if i < PolygonMaxPoints {
			polygon.Points[i] = point
			polygon.Length = i + 1
		} else {
			return nil
		}
	}

	return nil
}