package geom3d

type Point struct{
	X,Y,Z Fl
}

func (s Point)VectorTo(dest Point)Vector{
	return Vector{
		X: dest.X - s.X,
		Y: dest.Y - s.Y,
		Z: dest.Z - s.Z,
	}
}