package scene

import "math"

func NewCube(w, h, d float64) *Mesh {
	m := Mesh{}
	m.AddVertex(
		Vector{-w, -h, -d},
		Vector{-w, +h, -d},
		Vector{+w, +h, -d},
		Vector{+w, -h, -d},
		Vector{-w, -h, +d},
		Vector{-w, +h, +d},
		Vector{+w, +h, +d},
		Vector{+w, -h, +d},
	)
	m.AddPolygon(0, 1, 2, 3)
	m.AddPolygon(7, 6, 5, 4)
	m.AddPolygon(4, 3, 1, 0)
	m.AddPolygon(3, 2, 6, 7)
	m.AddPolygon(1, 5, 6, 2)
	m.AddPolygon(4, 0, 3, 7)
	return &m
}

// NewTriangle returns a new triangle based on the three supplied vectors
func NewTriangle(v1, v2, v3 Vector) *Mesh {
	m := Mesh{}
	m.AddVertex(v1, v2, v3)
	m.AddPolygon(0, 1, 2)
	return &m
}

// NewTriangleF returns a new triangle from 9 float values representing the x,y & z of each vertex
func NewTriangleF(x1, y1, z1, x2, y2, z2, x3, y3, z3 float64) *Mesh {
	return NewTriangle(Vector{x1, y1, z1}, Vector{x2, y2, z2}, Vector{x3, y3, z3})
}

func NewIcosphere() *Mesh {
	t := (1.0 + math.Sqrt(5.0)) / 2.0
	m := Mesh{}
	m.AddVertex(Vector{-1, t, 0})
	m.AddVertex(Vector{1, t, 0})
	m.AddVertex(Vector{-1, -t, 0})
	m.AddVertex(Vector{1, -t, 0})

	m.AddVertex(Vector{0, -1, t})
	m.AddVertex(Vector{0, 1, t})
	m.AddVertex(Vector{0, -1, -t})
	m.AddVertex(Vector{0, 1, -t})

	m.AddVertex(Vector{t, 0, -1})
	m.AddVertex(Vector{t, 0, 1})
	m.AddVertex(Vector{-t, 0, -1})
	m.AddVertex(Vector{-t, 0, 1})

	// 5 faces around point 0
	m.AddPolygon(0, 11, 5)
	m.AddPolygon(0, 5, 1)
	m.AddPolygon(0, 1, 7)
	m.AddPolygon(0, 7, 10)
	m.AddPolygon(0, 10, 11)

	// 5 adjacent faces
	m.AddPolygon(1, 5, 9)
	m.AddPolygon(5, 11, 4)
	m.AddPolygon(11, 10, 2)
	m.AddPolygon(10, 7, 6)
	m.AddPolygon(7, 1, 8)

	// 5 faces around point 3
	m.AddPolygon(3, 9, 4)
	m.AddPolygon(3, 4, 2)
	m.AddPolygon(3, 2, 6)
	m.AddPolygon(3, 6, 8)
	m.AddPolygon(3, 8, 9)

	// 5 adjacent faces
	m.AddPolygon(4, 9, 5)
	m.AddPolygon(2, 4, 11)
	m.AddPolygon(6, 2, 10)
	m.AddPolygon(8, 6, 7)
	m.AddPolygon(9, 8, 1)

	return &m
}