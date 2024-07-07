package entities

type Point struct {
	// Name       [255]byte `json:"name" binding:"required"`
	ID         uint64  `json:"id"`
	Longtitude float64 `json:"longtitude" binding:"required"`
	Lattitude  float64 `json:"lattitude" binding:"required"`
}
