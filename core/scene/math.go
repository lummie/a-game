package scene

import "math"

func RoundInt(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func RoundUp(f float64, places int) (newVal float64) {
	output := math.Pow(10, float64(places))
	return float64(RoundInt(f * output)) / output
}

