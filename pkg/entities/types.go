package entities

import (
	"fmt"
)

type MarkerType string

// GetType implements Marker.
func (m MarkerType) GetType() MarkerType {
	return m
}

const (
	PointType   MarkerType = "point"
	PolygonType MarkerType = "polygon"
	LineType    MarkerType = "line"
	NoneType    MarkerType = "none"
)

func TypeFromString(s string) (MarkerType, error) {
	switch s {
	case string(PointType):
		return PointType, nil
	case string(PolygonType):
		return PolygonType, nil
	case string(LineType):
		return LineType, nil
	default:
		return "", fmt.Errorf("invalid marker type %s", s)
	}
}
