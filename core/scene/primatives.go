package scene

func NewCube(w, h, d float64) *Mesh {
	c := Mesh{}
	c.AddVertex(
		Vector{-w, -h, -d},
		Vector{-w, +h, -d},
		Vector{+w, +h, -d},
		Vector{+w, -h, -d},
		Vector{-w, -h, +d},
		Vector{-w, +h, +d},
		Vector{+w, +h, +d},
		Vector{+w, -h, +d},
	)
	c.AddPolygon(0, 1, 2, 3)
	c.AddPolygon(7, 6, 5, 4)
	c.AddPolygon(4, 3, 1, 0)
	c.AddPolygon(3, 2, 6, 7)
	c.AddPolygon(1, 5, 6, 2)
	c.AddPolygon(4, 0, 3, 7)
	return &c
}

// NewTriangle returns a new triangle based on the three supplied vectors
func NewTriangle(v1, v2, v3 Vector) *Mesh {
	t := Mesh{}
	t.AddVertex(v1, v2, v3)
	t.AddPolygon(0, 1, 2)
	return &t
}

// NewTriangleF returns a new triangle from 9 float values representing the x,y & z of each vertex
func NewTriangleF(x1, y1, z1, x2, y2, z2, x3, y3, z3 float64) *Mesh {
	return NewTriangle(Vector{x1, y1, z1}, Vector{x2, y2, z2}, Vector{x3, y3, z3})
}
