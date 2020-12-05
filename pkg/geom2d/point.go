package geom2d

import "math"

type Point struct {
	X, Y float64
}

func (v Point) MoveByVector(vec Vector) Point {
	return Point{
		X: v.X + vec.X,
		Y: v.Y + vec.Y,
	}
}

func (source Point)VectorTo(dest Point)Vector{
	return Vector{
		X: dest.X - source.X,
		Y: dest.Y - source.Y,
	}
}

var NanPoint = Point{math.NaN(), math.NaN()}