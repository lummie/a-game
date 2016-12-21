package scene_test

import (
	"testing"
	"project69/core/scene"
)

func TestNewPolygon(t *testing.T) {
	p := scene.NewPolygon(1)

	if len(p.Indexes) != 0 {
		t.Error("Expected Polygon Indexes to number zero")
	}
}


func TestPolygon_AddVertex(t *testing.T) {
	p := scene.NewPolygon(1)

	if len(p.Indexes) != 0 {
		t.Error("Expected Polygon Indexes to number zero")
	}

	p.AddIndex(10)
	if len(p.Indexes) != 1 {
		t.Error("Expected Polygon Indexes to have 1 vertex")
	}

	if p.Indexes[0] != 10 {
		t.Errorf("Expected Index[0] to be 10 got %v", p.Indexes[0])
	}

	p.AddIndex(11)
	if len(p.Indexes) != 2 {
		t.Error("Expected Polygon Indexes to have 2 vertices")
	}
	if p.Indexes[0] != 10 {
		t.Errorf("Expected Index[0] to be 10 got %v", p.Indexes[0])
	}
	if p.Indexes[1] != 11 {
		t.Errorf("Expected Index[0] to be 10 got %v", p.Indexes[0])
	}

}