package scene



// Polygon holds a list of pointers to Vectors that make up a face.
// The order of the Vertex in the Array is important and defines the draw path
type Polygon struct {
	Indexes []int
}

// NewFace creates a new Face
func NewPolygon(capacity int) Polygon {
	p := Polygon{
		Indexes: make([]int, 0, capacity),
	}
	return p
}

// AddVertex adds a new vertex to the face
func (p *Polygon) AddIndex(i int) {
	p.Indexes = append(p.Indexes, i)
}

// Polygon holds a list of faces
type Polygons []Polygon


