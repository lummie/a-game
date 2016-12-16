package scene



// Polygon holds a list of pointers to Vectors that make up a face.
// The order of the Vertex in the Array is important and defines the draw path
type Polygon struct {
	Vertices VectorsRef
}

// NewFace creates a new Face
func NewPolygon(c int) Polygon {
	p := Polygon{
		Vertices: make(VectorsRef, 0, c),
	}
	return p
}

// AddVertex adds a new vertex to the face
func (f *Polygon) AddVertex(v *Vector) {
	f.Vertices = append(f.Vertices, v)
}

// Polygon holds a list of faces
type Polygons []Polygon

