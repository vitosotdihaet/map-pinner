package entities

type None struct{}

func (n *None) GetType() MarkerType {
	return NoneType
}
