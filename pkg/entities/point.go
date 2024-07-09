package entities

type Point struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type PointUpdate struct {
	Name      *string  `json:"name"`
	Longitude *float64 `json:"longitude"`
	Latitude  *float64 `json:"latitude"`
}