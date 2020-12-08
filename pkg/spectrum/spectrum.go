package spectrum

import "math"

const SpectrumRange = 300

type Data struct {
    Lambdas [SpectrumRange + 1]float64 // represents wavelengths from 400 to 700
}

//NewData initializes new Lambdas data with sine-wave distribution
func NewData() *Data {
    data := &Data{}
    for i := range data.Lambdas {
        v := float64(i) * math.Pi * 4 / 300
        data.Lambdas[i] = (math.Sin(v) + 1) / 2
    }
    return data
}

func (d *Data) ToXYZ() XYZ {

    r, g, b := 0.0, 0.0, 0.0
    for i, dp := range d.Lambdas {
        r += MatchingDataPoints[i].X * dp
        g += MatchingDataPoints[i].Y * dp
        b += MatchingDataPoints[i].Z * dp
    }
    return XYZ{r / SpectrumRange, g / SpectrumRange, b / SpectrumRange}
}
