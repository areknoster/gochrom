package geom2d

import "math"

type Point struct {
	X, Y float64
}

func (p Point) MoveByVector(vec Vector) Point {
	return Point{
		X: p.X + vec.X,
		Y: p.Y + vec.Y,
	}
}

func (p Point)VectorTo(dest Point)Vector{
	return Vector{
		X: dest.X - p.X,
		Y: dest.Y - p.Y,
	}
}

//point must be normalized
func (p Point) InvertY()Point{
	p.Y = 1-p.Y
	return p
}


var NanPoint = Point{math.NaN(), math.NaN()}