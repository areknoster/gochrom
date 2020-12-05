package geom2d

import (
	"math"
)

type Shape struct {
	PointsSet
}

func NewShape(pointsSet PointsSet) Shape {
	return Shape{PointsSet: pointsSet}
}



type ShapeSet []Shape



func Middle(s Shape) Point {
	p := Point{0, 0}
	for _, vertex := range s.PointsSet {
		p.X, p.Y = p.X+vertex.X, p.Y+vertex.Y
	}
	p.X, p.Y = p.X/float64(len(s.PointsSet)), p.Y/float64(len(s.PointsSet))
	return p
}

//ClosestToShape returns index of closest shape
func (ss ShapeSet) ClosestToShape(p Point) int {
	if len(ss) == 0 {
		return -1
	}
	minNorm := math.MaxFloat64
	minIndex := -1
	for i, shape := range ss {
		norm := VecBetweenPoints(Middle(shape), p).SqNorm()
		if minNorm > norm {
			minIndex = i
			minNorm = norm
		}
	}
	return minIndex
}

func (s Shape) ToEdgeSet() EdgeSet {
	ss := make([]Edge, len(s.PointsSet))
	for i:= 0; i < len(ss) -1; i++ {
		ss[i] = NewEdge(s.PointsSet[i], s.PointsSet[i+1])
	}
	ss[len(ss) - 1] = NewEdge(s.PointsSet[len(s.PointsSet) -1], s.PointsSet[0])
	return ss
}

