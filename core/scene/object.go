package scene


type Object interface {
	Faces() *Faces
	Normalize()
}

type Objects []Object


