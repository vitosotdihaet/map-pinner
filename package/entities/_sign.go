package entities

type SignTag uint8

const (
	None = iota
	Home
)

type Sign struct {
	tag   SignTag
	point Point
}
