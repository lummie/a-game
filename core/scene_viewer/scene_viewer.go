package main

import "github.com/lummie/a-game/core/scene"

func main(){
	s := scene.NewScene(1)
	vp := scene.NewViewport(400, 400)

	is := scene.NewIcosphere(4)
	s.AddMesh(is)

	//vp.Rasterize(s)
	//vp.RenderPng("test.png")
	vp.RenderSvg(s, "test.svg")
}
