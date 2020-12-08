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

func clamp(v, min, max float64) float64 {
    if v < min {
        return min
    }
    if v > max {
        return max
    }
    return v
}

func normRangeFF(v, min, max float64) float64 {
    return (v - min) / (max - min)
}

func NormRangeIF(v, min, max int) float64{
    if v == min{
        return 0
    }
    return float64(v - min) / float64(max - min)
}

func ClampNormRect(point, min, max geom2d.Point) geom2d.Point {
    return geom2d.Point{
        X: normRangeFF(clamp(point.X, min.X, max.X), min.X, max.X),
        Y: normRangeFF(clamp(point.Y, min.Y, max.Y), min.Y, max.Y),
    }
}
