package main

import (
	"project69/core/scene"
	"fmt"
)

func main(){
	s := scene.NewScene(1)
	vp := scene.NewViewport(400, 400)

	is := scene.NewIcosphere(4)
	fmt.Printf("%v", is)
	s.AddMesh(is)

	//vp.Rasterize(s)
	//vp.RenderPng("test.png")
	vp.RenderSvg(s, "test.svg", 1)
}
