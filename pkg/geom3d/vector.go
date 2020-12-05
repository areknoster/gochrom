package geom3d

import (
	"math"
)

type Vector struct {
	X, Y, Z Fl
}

var WrongVector = Vector{X: math.NaN(), Y: math.NaN(), Z: math.NaN()}

func (a Vector) Add(b Vector) Vector{
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func (a Vector) Substract(b Vector) Vector{
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func (v Vector) EuclidNormSq() Fl {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector) Length() Fl {
	return math.Sqrt(v.EuclidNormSq())
}

func (v Vector) ToVersor() Vector {
	l := v.Length()
	return Vector{
		v.X / l,
		v.Y / l,
		v.Z / l,
	}
}

func (v Vector)TimesScalar(sc Fl) Vector{
	return Vector{
		v.X * sc,
		v.Y * sc,
		v.Z * sc,
	}
}

func (a Vector) Dot(b Vector) Fl{
	return a.X * b.X + a.Y * b.Y + a.Z * b.Z
}

func (a Vector) MultiplyByElem(b Vector) Vector{
	return Vector{
		X: a.X * b.X,
		Y: a.Y * b.Y,
		Z: a.Z * b.Z,
	}
}

func (a Vector) Cross(b Vector) Vector {
	return Vector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}


func (v Vector) Normalize() Vector{
	max := -math.MaxFloat64
	for _, val := range []float64{v.X, v.Y, v.Z}{
		if val > max{
			max = val
		}
	}

	return Vector{
		X: v.X/max,
		Y: v.Y/max,
		Z: v.Z/max,
	}
}

func VectorBetweenPoints(a,b Point) Vector{
	return Vector{
		X: b.X - a.X,
		Y: b.Y - a.Y,
		Z: b.Z - a.Z,
	}
}

