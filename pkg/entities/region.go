package entities

type Region struct {
	ID       uint64    `json:"id"`
	Name     string    `json:"name"`
	Points   []Point   `json:"points"`
	Polygons []Polygon `json:"polygons"`
	Lines    []Line    `json:"lines"`
}

type RegionUpdate struct {
	Name *string `json:"name"`
}
