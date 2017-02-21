package scene_test

import (
	"testing"
	"github.com/lummie/a-game/core/scene"
)

func TestVertex_ApproxEqual(t *testing.T) {
	a := scene.Vector{2.0, 2.0, 2.0}
	b := scene.Vector{2.0, 2.0, 2.0}
	if !a.ApproxEqual(b) {
		t.Errorf("Expected %v to equal %v :", a, b)
	}

	c := scene.Vector{1.0, 2.0, 2.0}
	if a.ApproxEqual(c) {
		t.Errorf("Expected %v to NOT equal %v :", a, c)
	}
}

func TestVertex_Length(t *testing.T) {
	v := scene.Vector{1.0, 0.0, 0.0}
	l := v.Length()
	if l != 1.0 {
		t.Errorf("Vertex Length incorrect %v", l)
	}

	v = scene.Vector{1.0, 2.0, 3.0}
	l = v.Length()
	if l != 3.7416573867739413 {
		t.Errorf("Vertex Length incorrect %v", l)
	}
}

func TestVertex_Add(t *testing.T) {
	a := scene.Vector{2.0, 2.0, 2.0}
	b := scene.Vector{2.0, 2.0, 2.0}
	e := scene.Vector{4.0, 4.0, 4.0}
	l := a.Add(b)
	if !l.ApproxEqual(e) {
		t.Errorf("Expected %v not %v :", e, l)
	}
}

func TestVertex_Subtract(t *testing.T) {
	a := scene.Vector{2.0, 2.0, 2.0}
	b := scene.Vector{2.0, 2.0, 2.0}
	e := scene.Vector{0.0, 0.0, 0.0}
	l := a.Subtract(b)
	if !l.ApproxEqual(e) {
		t.Errorf("Expected %v not %v :", e, l)
	}
}

func TestVertex_Multiply(t *testing.T) {
	a := scene.Vector{2.0, 2.0, 2.0}
	b := scene.Vector{2.0, 2.0, 2.0}
	e := scene.Vector{4.0, 4.0, 4.0}
	l := a.Multiply(b)
	if !l.ApproxEqual(e) {
		t.Errorf("Expected %v not %v :", e, l)
	}
}

func TestVertex_Divide(t *testing.T) {
	a := scene.Vector{2.0, 2.0, 2.0}
	b := scene.Vector{2.0, 4.0, 1.0}
	e := scene.Vector{1.0, 0.5, 2.0}
	l := a.Divide(b)
	if !l.ApproxEqual(e) {
		t.Errorf("Expected %v not %v :", e, l)
	}
}

func TestVertex_MidPointTo(t *testing.T) {
	a := scene.Vector{0.0, 0.0, 0.0}
	b := scene.Vector{3.0, 3.0, 3.0}
	e := scene.Vector{1.5, 1.5, 1.5}
	l := a.MidPointTo(b)
	if !l.ApproxEqual(e) {
		t.Errorf("Expected %v not %v :", e, l)
	}
}

func TestVertex_AddScalar(t *testing.T) {
	a := scene.Vector{2.0, 2.0, 2.0}
	b := 1.1
	e := scene.Vector{3.1, 3.1, 3.1}
	l := a.AddScalar(b)
	if !l.ApproxEqual(e) {
		t.Errorf("Expected %v not %v :", e, l)
	}
}

func TestVertex_SubtractScalar(t *testing.T) {
	a := scene.Vector{2.0, 2.0, 2.0}
	b := 1.1
	e := scene.Vector{0.9, 0.9, 0.9}
	l := a.SubtractScalar(b)
	if !l.ApproxEqual(e) {
		t.Errorf("Expected %v not %v :", e, l)
	}
}

func TestVertex_MultiplyScalar(t *testing.T) {
	a := scene.Vector{2.0, 2.0, 2.0}
	b := 1.1
	e := scene.Vector{2.2, 2.2, 2.2}
	l := a.MultiplyScalar(b)
	if !l.ApproxEqual(e) {
		t.Errorf("Expected %v not %v :", e, l)
	}
}

func TestVertex_DivideScalar(t *testing.T) {
	a := scene.Vector{2.0, 2.0, 2.0}
	b := 2.0
	e := scene.Vector{1.0, 1.0, 1.0}
	l := a.DivideScalar(b)
	if !l.ApproxEqual(e) {
		t.Errorf("Expected %v not %v :", e, l)
	}
}

func TestVertex_CrossProduct(t *testing.T) {
	a := scene.Vector{0.0, 0.0, 1.0}
	b := scene.Vector{1.0, 0.0, 0.0}
	e := scene.Vector{0.0, 1.0, 0.0}
	l := a.CrossProduct(b)
	if !l.ApproxEqual(e) {
		t.Errorf("Expected %v not %v :", e, l)
	}
}

func TestVertex_DotProduct(t *testing.T) {
	{
		// positive
		a := scene.Vector{2.0, 0.0, 1.0}
		b := scene.Vector{2.0, 3.5, 0.0}
		e := 4.0
		l := a.DotProduct(b)
		if l != e {
			t.Errorf("Expected %v not %v :", e, l)
		}
	}
	{
		// negative
		a := scene.Vector{2.0, 2.0, 2.0}
		b := scene.Vector{-2.5, -2.5, -2.5}
		e := -15.0
		l := a.DotProduct(b)
		if l != e {
			t.Errorf("Expected %v not %v :", e, l)
		}
	}
	{
		// No change
		a := scene.Vector{2.0, 2.0, 2.0}
		b := scene.Vector{0.0, 0.0, 0.0}
		e := 0.0
		l := a.DotProduct(b)
		if l != e {
			t.Errorf("Expected %v not %v :", e, l)
		}
	}
}

func TestVertex_LengthSquared(t *testing.T) {
	a := scene.Vector{2.0, 2.0, 2.0}
	e := 12.0
	l := a.LengthSquared()
	if l != e {
		t.Errorf("Expected %v not %v :", e, l)
	}
}

func TestVertex_Normalize(t *testing.T) {
	a := scene.Vector{3.0, 1.0, 2.0}
	e := scene.Vector{0.8017837257372732, 0.2672612419124244, 0.5345224838248488}
	l := a.Normalize()
	if !l.ApproxEqual(e) {
		t.Errorf("Expected %v not %v :", e, l)
	}

}
