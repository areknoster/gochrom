package normde

import "math"

//DenormalizeFloat maps normalized[0,1] float value to [min, max] int value
func DenormalizeFloat(v float64, min, max int) int{
    rang := float64(max - min)
    return min + int(math.Round(v * rang))
}
