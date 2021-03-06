package scene

import (
	"math"
)

func NewCube(w, h, d float64) *Mesh {
	m := NewMesh(8, 6)
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
	return m
}

// NewTriangle returns a new triangle based on the three supplied vectors
func NewTriangle(v1, v2, v3 Vector) *Mesh {
	m := NewMesh(3, 1)
	m.AddVertex(v1, v2, v3)
	m.AddPolygon(0, 1, 2)
	return m
}

// NewTriangleF returns a new triangle from 9 float values representing the x,y & z of each vertex
func NewTriangleF(x1, y1, z1, x2, y2, z2, x3, y3, z3 float64) *Mesh {
	return NewTriangle(Vector{x1, y1, z1}, Vector{x2, y2, z2}, Vector{x3, y3, z3})
}

func NewIcosphere(detailLevel int) *Mesh {
	t := (1.0 + math.Sqrt(5.0)) / 2.0
	m := NewMesh(12, 20)
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

	onSphere := func(v Vector) Vector {
		l := math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
		return Vector{v.X / l, v.Y / l, v.Z / l}
	}

	currentMesh := m
	for i := 0; i < detailLevel; i++ {
		detailMesh := NewMesh(6, 4)
		o := 0 // offset
		for _, p := range (currentMesh.Polygons) {
			i1 := p.Indexes[0]
			i2 := p.Indexes[1]
			i3 := p.Indexes[2]
			a := currentMesh.getVertex(i1).MidPointTo(currentMesh.getVertex(i2))
			b := currentMesh.getVertex(i2).MidPointTo(currentMesh.getVertex(i3))
			c := currentMesh.getVertex(i3).MidPointTo(currentMesh.getVertex(i1))
			// map to sphere
			detailMesh.AddVertex(onSphere(currentMesh.getVertex(i1)), onSphere(currentMesh.getVertex(i2)),
				onSphere(currentMesh.getVertex(i3)), onSphere(a), onSphere(b), onSphere(c))
			detailMesh.AddPolygon(o + 0, o + 3, o + 5)
			detailMesh.AddPolygon(o + 1, o + 4, o + 3)
			detailMesh.AddPolygon(o + 2, o + 5, o + 4)
			detailMesh.AddPolygon(o + 3, o + 4, o + 5)
			o = o + 6
		}
		currentMesh = detailMesh
	}

	// copy last mesh to the result mes
	resultMesh := NewMesh(len(currentMesh.Vertices), len(currentMesh.Polygons))
	for _, v := range (currentMesh.Vertices) {
		resultMesh.AddVertex(v)
	}
	for _, p := range (currentMesh.Polygons) {
		resultMesh.AddPolygon(p.Indexes...)
	}

	return resultMesh
}