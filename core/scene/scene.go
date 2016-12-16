package scene


type Scene struct {
	Meshes              Meshes
	WorldTransformation Matrix
}

// NewScene creates a new scene
func NewScene(c int) *Scene {
	s := Scene{
		Meshes:make(Meshes, 0, c),
		WorldTransformation: NewScale(Vector{1.0, 1.0, 1.0}),
	}
	return &s
}

func (s *Scene) AddMesh(m *Mesh) {
	s.Meshes = append(s.Meshes, m)
}

