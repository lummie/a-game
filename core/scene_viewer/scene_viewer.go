package main

import (
	"project69/core/scene"
)

func main(){
	s := scene.NewScene(1)
	vp := scene.NewViewport(400, 400)

	g01 := scene.NewMesh(3, 1)
	g01.AddVertex(
		scene.Vector{1.0, 1.0, 1.0},
		scene.Vector{3.0, 1.0, 2.0},
		scene.Vector{3.0, 3.0, 3.0},
	)
	g01.AddPolygon(0, 1, 2)

	//	s.AddMesh(g01)
	//s.AddMesh(scene.NewTriangleF(5.0,5.0,1.0,8.0,8.0,8.0,10.0,10.0,10.0))
	s.AddMesh(scene.NewCube(10, 10, 10))

	vp.Rasterize(s)
	vp.RenderPng("test.png")
}
