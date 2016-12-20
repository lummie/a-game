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
func (m *Mesh) AddVertex(vertices ...Vector) {
	for _, v := range (vertices) {
		m.Vertices = append(m.Vertices, v)
	}
}

// AddPolygon adds a new face to the object using the specified vertices' indexes
func (m *Mesh) AddPolygon(indexes ...int) *Polygon {
	p := NewPolygon(len(indexes))
	for _, v := range indexes {
		p.AddIndex(v)
	}
	m.Polygons = append(m.Polygons, p)
	return &p
}

func (m *Mesh) String() string {
	return fmt.Sprintf("vertices:%+v faces:%+v", m.Vertices, m.Polygons)
}

func (m *Mesh) Normalize() {
	for i := range m.Vertices {
		m.Vertices[i] = m.Vertices[i].Normalize()
	}
}

func (m *Mesh) getVertex(index int) Vector {
	return m.Vertices[index]
}


