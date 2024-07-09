package entities

type Graph struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Points [30]*Point
}
