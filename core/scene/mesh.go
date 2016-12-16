package scene

import "fmt"

type Mesh struct {
	Vertices            Vectors
	Polygons            Polygons
	WorldTransformation Matrix
}

type Meshes []*Mesh

// NewGenericObject returns an new object allocated with the specified capacities for vertices and faces
func NewMesh(vertexCapacity, polygonCapacity int) *Mesh {
	g := Mesh{
		Vertices : make(Vectors, 0, vertexCapacity),
		Polygons : make(Polygons, 0, polygonCapacity),
	}
	return &g
}

// AddVertex adds one or more vertices to the object
func (g *Mesh) AddVertex(vertices ...Vector) {
	for _, v := range (vertices) {
		g.Vertices = append(g.Vertices, v)
	}
}

// AddFace adds a new face to the object using the specified vertices' indexes
func (g *Mesh) AddPolygon(indexes ...int) {
	f := NewPolygon(len(indexes))
	for i := range indexes {
		f.AddVertex(&g.Vertices[i])
	}
	g.Polygons = append(g.Polygons, f)
}

func (g *Mesh) String() string {
	return fmt.Sprintf("vertices:%+v faces:%+v", g.Vertices, g.Polygons)
}

func (g *Mesh) Normalize() {
	for i := range g.Vertices {
		g.Vertices[i] = g.Vertices[i].Normalize()
	}
}
