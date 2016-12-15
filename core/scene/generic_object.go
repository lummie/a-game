package scene

import "fmt"

type GenericObject struct {
	vertices Vertices
	faces Faces
}


// NewGenericObject returns an new object allocated with the specified capacities for vertices and faces
func NewGenericObject(vertexCapacity, faceCapacity int) *GenericObject{
	g := GenericObject{
		vertices : make(Vertices,0,vertexCapacity),
		faces : make(Faces,0,faceCapacity),
	}
	return &g
}

// AddVertex adds one or more vertices to the object
func (g *GenericObject) AddVertex(vertices ...Vertex) {
	for _,v := range(vertices) {
		g.vertices = append(g.vertices,v)
	}
}

// AddFace adds a new face to the object using the specified vertices' indexes
func (g *GenericObject) AddFace(indexes ...int) {
	f := NewFace(len(indexes))
	for i := range indexes {
		f.AddVertex(&g.vertices[i])
	}
	g.faces = append(g.faces,f)
}

func (g *GenericObject) Faces() *Faces {
	return &g.faces
}

func (g *GenericObject) Vertices() *Vertices {
	return &g.vertices
}

func (g *GenericObject) String() string {
	return fmt.Sprintf("vertices:%+v faces:%+v", g.vertices, g.faces)
}

func (g *GenericObject) Normalize(){
	for i := range g.vertices {
		g.vertices[i] = g.vertices[i].Normalize()
	}
}
