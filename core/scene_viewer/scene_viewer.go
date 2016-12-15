package main

import (
	"tessallation/core/scene"
	"fmt"
)

func main(){
	s := scene.NewScene(1)

	g01 := scene.NewGenericObject(3,1)
	g01.AddVertex(
		scene.Vertex{1.0,1.0,1.0},
		scene.Vertex{2.0,2.0,2.0},
		scene.Vertex{3.0,3.0,3.0},
	)
	g01.AddFace(0,1,2)
	s.AddObject(g01)
	s.AddObject(scene.NewTriangleF(1.0,1.0,1.0,2.0,2.0,2.0,3.0,3.0,3.0))
	s.AddObject(scene.NewCube(10,10,10))

	fmt.Printf("%+v\n",s.Objects)
	s.Objects[0].Normalize()
	fmt.Printf("%+v\n",s.Objects)

}
