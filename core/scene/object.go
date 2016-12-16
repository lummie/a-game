package scene


type Object interface {
	AddVertex(vertices ...Vector)
	AddPolygon(indexes ...int)
	Normalize()
}

type Objects []Object


