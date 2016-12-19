package scene_test

import (
	"testing"
	"project69/core/scene"
)

func TestNewPolygon(t *testing.T) {
	p := scene.NewPolygon(1)

	if len(p.Vertices) != 0 {
		t.Error("Expected Polygon Vertices to number zero")
	}
}


func TestPolygon_AddVertex(t *testing.T) {
	p := scene.NewPolygon(1)

	if len(p.Vertices) != 0 {
		t.Error("Expected Polygon Vertices to number zero")
	}

	p.AddVertex(&scene.Vector{1.0,2.0,3.0})
	if len(p.Vertices) != 1 {
		t.Error("Expected Polygon Vertices to have 1 vertex")
	}
	if p.Vertices[0].X != 1.0 ||
		p.Vertices[0].Y != 2.0 ||
		p.Vertices[0].Z != 3.0 {
		t.Errorf("Expected vertices zero to be 1,2,3 and got %v", p.Vertices[0])
	}

	p.AddVertex(&scene.Vector{4.0,5.0,6.0})
	if len(p.Vertices) != 2 {
		t.Error("Expected Polygon Vertices to have 2 vertices")
	}
	if p.Vertices[0].X != 1.0 ||
		p.Vertices[0].Y != 2.0 ||
		p.Vertices[0].Z != 3.0 {
		t.Errorf("Expected vertices zero to be 1,2,3 and got %v", p.Vertices[0])
	}
	if p.Vertices[1].X != 4.0 ||
		p.Vertices[1].Y != 5.0 ||
		p.Vertices[1].Z != 6.0 {
		t.Errorf("EXpected vertices ONE to be 4,5,6 and got %v", p.Vertices[1])
	}

}