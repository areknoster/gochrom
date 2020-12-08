package raster

type Pixel struct {
    X, Y int
}

func (p Pixel) MoveByVec(v Pixel) Pixel {
    return Pixel{
        X: p.X + v.X,
        Y: p.Y + v.Y,
    }
}
