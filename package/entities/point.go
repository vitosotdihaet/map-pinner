package entities

type Point struct {
	ID         uint64  `json:"id"`
	Longtitude float64 `json:"longtitude"`
	Lattitude  float64 `json:"lattitude"`
}
