package entities

type Point struct {
	ID         uint64  `json:"id"`
	Name       string  `json:"name" binding:"required"`
	Longtitude float64 `json:"longtitude" binding:"required"`
	Lattitude  float64 `json:"lattitude" binding:"required"`
}
