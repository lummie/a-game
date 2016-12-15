package scene_test

import (
	"testing"
	"tessallation/core/scene"
)

func TestNewTriangle(t *testing.T) {
	tr := scene.NewTriangleF(1.0,1.0,1.0,2.0,2.0,2.0,3.0,3.0,3.0)
	if !tr.Vertices()[0].ApproxEqual(scene.Vertex{1.0,1.0,1.0}) {
		t.Error("Expected Vertex 1 to be correct")
	}
	if !tr.Vertices()[1].ApproxEqual(scene.Vertex{2.0,2.0,2.0}) {
		t.Error("Expected Vertex 2 to be correct")
	}
	if !tr.Vertices()[2].ApproxEqual(scene.Vertex{3.0,3.0,3.0}) {
		t.Error("Expected Vertex 3 to be correct")
	}
}

func TestNewTriangleV(t *testing.T) {
	tr1 := scene.Vertex{1.0,1.0,1.0}
	tr2 := scene.Vertex{2.0,2.0,2.0}
	tr3 := scene.Vertex{3.0,3.0,3.0}
	tr := scene.NewTriangle(tr1,tr2,tr3)
	if !tr.Vertices()[0].ApproxEqual(tr1) {
		t.Error("Expected Vertex 1 to be correct")
	}
	if !tr.Vertices()[1].ApproxEqual(tr2) {
		t.Error("Expected Vertex 2 to be correct")
	}
	if !tr.Vertices()[2].ApproxEqual(tr3) {
		t.Error("Expected Vertex 3 to be correct")
	}
}

func TestTriangle_SubDivide(t *testing.T) {
	tr := scene.NewTriangleF(1.0,1.0,1.0,2.0,2.0,2.0,3.0,3.0,3.0)
	sd := tr.SubDivide()
	e1 := scene.Vertex{1.5,1.5,1.5}
	e2 := scene.Vertex{2.5,2.5,2.5}
	e3 := scene.Vertex{2.0,2.0,2.0}

	if !sd.Vertices()[0].ApproxEqual(e1) {
		t.Error("Expected Vertex 1 %v to be %v", sd.Vertices()[0], e1)
	}
	if !sd.Vertices()[1].ApproxEqual(e2) {
		t.Error("Expected Vertex 2 %v to be %v", sd.Vertices()[1], e2)
	}
	if !sd.Vertices()[2].ApproxEqual(e3) {
		t.Error("Expected Vertex 3 %v to be %v", sd.Vertices()[2], e3)
	}
}