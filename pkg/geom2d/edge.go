package geom2d

type Edge struct{
	A,B Point
}

func NewEdge(A,B Point) Edge {
	return Edge{A,B}
}

func (e Edge) ToVector() Vector{
	return VecBetweenPoints(e.A, e.B)
}

func (e Edge) MiddlePoint() Point{
	return e.A.MoveByVector(e.ToVector().TimesScalar(0.5))
}

type EdgeSet []Edge

func (ss EdgeSet) ShapeFromMidpoints() Shape {
	vs := make([]Point, len(ss))
	for i, s := range ss {
		vs[i] = s.MiddlePoint()
	}
	return Shape{vs}
}

func (e Edge) MoveByVector(v Vector) Edge {
	return Edge{e.A.MoveByVector(v), e.B.MoveByVector(v)}
}
