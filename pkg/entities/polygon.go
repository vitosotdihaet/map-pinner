package entities

import (
	"encoding/json"
)


type Polygon struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Points []Point `json:"points"`
}

func (polygon *Polygon) UnmarshalJSON(data []byte) error {
	var raw struct {
		Name   string `json:"name"`
		Points []json.RawMessage `json:"points"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	polygon.Name = raw.Name
	polygon.Points = make([]Point, len(raw.Points))

	for i, rawPoint := range raw.Points {
		var point Point
		if err := json.Unmarshal(rawPoint, &point); err != nil {
			return err
		}
		polygon.Points[i] = point
	}

	return nil
}

type PolygonUpdate struct {
	Name   *string `json:"name"`
	Points *[]Point `json:"points"`
}