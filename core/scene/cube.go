package scene

type Cube struct {
	GenericObject
}

func NewCube(w,h,d float64) *Cube {
	c := Cube{}
	c.AddVertex(
		Vertex{-w,-h,-d},
		Vertex{-w,+h,-d},
		Vertex{+w,+h,-d},
		Vertex{+w,-h,-d},
		Vertex{-w,-h,+d},
		Vertex{-w,+h,+d},
		Vertex{+w,+h,+d},
		Vertex{+w,-h,+d},
	)
	c.AddFace(0,1,2,3)
	c.AddFace(7,6,5,4)
	c.AddFace(4,3,1,0)
	c.AddFace(3,2,6,7)
	c.AddFace(1,5,6,2)
	c.AddFace(4,0,3,7)
	return &c
}