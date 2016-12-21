package scene

import "math"

type GeoLocation struct {
	lat  float64
	long float64
}

const (
	DegToRad = math.Pi / 180
	RadToDeg = 180 / math.Pi
)

func GoeLocationFromVector(v Vector) GeoLocation {

	return GeoLocation{
		lat : math.Atan2(v.Z, math.Sqrt((v.X * v.X) + (v.Y * v.Y))) * RadToDeg,
		long : math.Atan2(v.Y, v.X) * RadToDeg,
	}
}

func (g *GeoLocation) Vector() Vector {
	lat := g.lat * DegToRad
	long := g.long * DegToRad
	return Vector{math.Cos(lat) * math.Cos(long), math.Cos(lat) * math.Sin(long), math.Sin(lat)}
}