package entities

type Graph struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name" binding:"required"`
	Points [30]*Point
}
