package scene


type Scene struct {
	Objects Objects

}

// NewScene creates a new scene
func NewScene(c int) *Scene {
	s := Scene{
		Objects:make([]Object,0,c),
	}
	return &s
}

func (s *Scene) AddObject(o Object) {
	s.Objects = append(s.Objects, o)
}

