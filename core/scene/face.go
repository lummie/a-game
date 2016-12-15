package scene



// Face holds a list of pointers to Vectors that make up a face.
// The order of the Vertex in the Array is important and defines the draw path
type Face struct {
	Vertices VerticesRef
}

func NewFace(c int) Face{
	p := Face{
		Vertices: make(VerticesRef,0,c),
	}
	return p
}

func (f *Face) AddVertex(v *Vertex) {
	f.Vertices = append(f.Vertices,v)
}




// Faces holds a list of faces
type Faces []Face

