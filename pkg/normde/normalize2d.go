package normde

import (
	"github.com/areknoster/gochrom/pkg/geom2d"
	"github.com/areknoster/gochrom/pkg/raster"
)

func NormPoint2D(point raster.Pixel, w, h int) geom2d.Point {
	return geom2d.Point{
		X: float64(point.X) / float64(w),
		Y: float64(point.Y) / float64(h),
	}
}

func NormVector2D(vec raster.Pixel, w, h int) geom2d.Vector {
	return geom2d.Vector{
		X: float64(vec.X) / float64(w),
		Y: float64(vec.Y) / float64(h),
	}
}

func DenormPoint2D(p geom2d.Point, w, h int) raster.Pixel {
	return raster.Pixel{int(float64(w) * p.X), int(float64(h) * p.Y)}
}
