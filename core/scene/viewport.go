package scene

import (
	"os"
	"image"
	"image/color"
	"image/png"
	"log"
)

type Viewport struct {
	Width          int
	Height         int
	Transformation Matrix
	buffer         [][]VectorsRef
}

func NewViewport(w, h int) *Viewport {
	v := Viewport{
		Width : w,
		Height : h,
	}

	v.buffer = make([][]VectorsRef, h)
	for i := range (v.buffer) {
		v.buffer[i] = make([]VectorsRef, w)
	}
	v.CalcTransformation()
	return &v
}

func (v *Viewport) CalcTransformation() {
	tv := Vector{float64(v.Width) / 2.0, float64(v.Height) / 2.0, 0.0}
	v.Transformation = NewIdentity().Translate(tv)
}

func (v *Viewport) Rasterize(scene *Scene) {
	for _, m := range scene.Meshes {
		for _, p := range m.Polygons {
			for _, ov := range p.Vertices {
				wv := scene.WorldTransformation.MulPositionW(*ov)
				vv := v.Transformation.MulPositionW(wv)
				x := int(vv.X)
				y := int(vv.Y)
				if x >= 0 && x < v.Width && y >= 0 && y < v.Height {
					v.buffer[y][x] = append(v.buffer[y][x], &vv)
				}
				log.Println(vv)
			}
		}
	}
}

func (v *Viewport) RenderPng(filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img := image.NewRGBA(image.Rect(0, 0, v.Width, v.Height))
	for y := 0; y < img.Rect.Max.Y; y++ {
		for x := 0; x < img.Rect.Max.X; x++ {
			c := color.RGBA{255, 255, 255, 255}
			if (len(v.buffer[y][x]) > 0) {
				c = color.RGBA{0, 0, 0, 255}
			}

			img.Set(x, y, c)
		}
	}
	if err = png.Encode(f, img); err != nil {
		panic(err)
	}
}