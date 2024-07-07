package entities

type Polygon struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name" binding:"required"`
	Points [20]*Point
}
