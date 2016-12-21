package scene

import (
	"testing"
)

func TestGeoLocationFromVector(t *testing.T) {
	geo := GoeLocationFromVector(Vector{0.5, 0.5, 0.7071})
	if RoundInt(geo.lat) != 45 {
		t.Errorf("Expected lat: %v got %v", 45, int(geo.lat))
	}
}

func TestGeoLocationFromVectorAndBack(t *testing.T) {
	geo := GoeLocationFromVector(Vector{0.5, 0.5, 0.7071})
	if RoundInt(geo.lat) != 45 {
		t.Errorf("Expected lat: %v got %v", 45, int(geo.lat))
	}
	v := geo.Vector().Round(4)
	if v.X != 0.5 && v.Y != 0.5 && v.Z != 0.7071 {
		t.Errorf("Expected %v got %v ", Vector{0.5, 0.5, 0.7071}, v)
	}
}
