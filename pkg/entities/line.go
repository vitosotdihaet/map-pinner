package entities

type Line Polygon

func (p *Line) GetType() MarkerType {
	return LineType
}

type LineUpdate PolygonUpdate

func (p *LineUpdate) GetType() MarkerType {
	return LineType
}
