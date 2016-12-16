package scene_test

import (
	"testing"
	"project69/core/scene"
)

func TestNewTriangle(t *testing.T) {
	tr := scene.NewTriangleF(1.0, 1.0, 1.0, 2.0, 2.0, 2.0, 3.0, 3.0, 3.0)
	if !tr.Vertices[0].ApproxEqual(scene.Vector{1.0, 1.0, 1.0}) {
		t.Error("Expected Vertex 1 to be correct")
	}
	if !tr.Vertices[1].ApproxEqual(scene.Vector{2.0, 2.0, 2.0}) {
		t.Error("Expected Vertex 2 to be correct")
	}
	if !tr.Vertices[2].ApproxEqual(scene.Vector{3.0, 3.0, 3.0}) {
		t.Error("Expected Vertex 3 to be correct")
	}
}

func TestNewTriangleV(t *testing.T) {
	tr1 := scene.Vector{1.0, 1.0, 1.0}
	tr2 := scene.Vector{2.0, 2.0, 2.0}
	tr3 := scene.Vector{3.0, 3.0, 3.0}
	tr := scene.NewTriangle(tr1, tr2, tr3)
	if !tr.Vertices[0].ApproxEqual(tr1) {
		t.Error("Expected Vertex 1 to be correct")
	}
	if !tr.Vertices[1].ApproxEqual(tr2) {
		t.Error("Expected Vertex 2 to be correct")
	}
	if !tr.Vertices[2].ApproxEqual(tr3) {
		t.Error("Expected Vertex 3 to be correct")
	}
}

