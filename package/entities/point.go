package entities

type Point struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Lattitude float64 `json:"lattitude" binding:"required"`
}
