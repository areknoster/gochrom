package geom2d

import "math"

type PointsSet []Point

func (ps PointsSet) ClosestToVertex(p Point) int {
	index := -1
	closestSqDist := math.MaxFloat64
	for i, vertex := range ps {
		if dist := SqDistBetweenPoints(vertex, p); dist < closestSqDist {
			closestSqDist = dist
			index = i
		}
	}
	return index
}
