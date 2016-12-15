package scene

type Triangle struct {
	GenericObject
}

// NewTriangle returns a new triangle based on the three supplied vectors
func NewTriangle(v1,v2,v3 Vertex) *Triangle {
	t := Triangle{}
	t.AddVertex(v1,v2,v3)
	t.AddFace(0,1,2)
	return &t
}

// NewTriangleF returns a new triangle from 9 float values representing the x,y & z of each vertex
func NewTriangleF(x1,y1,z1,x2,y2,z2,x3,y3,z3 float64) *Triangle {
	return NewTriangle(Vertex{x1,y1,z1},Vertex{x2,y2,z2}, Vertex{x3,y3,z3})
}



// SubDivide returns a triangle that represents the midpoints of each segment of the triangle
func (t *Triangle) SubDivide() *Triangle {
	p1 := t.vertices[0].MidPointTo(t.vertices[1])
	p2 := t.vertices[1].MidPointTo(t.vertices[2])
	p3 := t.vertices[2].MidPointTo(t.vertices[0])
	return NewTriangle(p1,p2,p3)
}
