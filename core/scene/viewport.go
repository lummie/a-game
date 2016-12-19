package scene

import (
	"os"
	"image"
	"image/color"
	"image/png"
	"log"
	"bufio"
	"fmt"
	"strings"
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
			for _, index := range p.Indexes {
				vertex := m.getVertex(index)
				wv := scene.WorldTransformation.MulPositionW(vertex)
				vv := v.Transformation.MulPositionW(wv)
				x := int(vv.X)
				y := int(vv.Y)
				if x >= 0 && x < v.Width && y >= 0 && y < v.Height {
					v.buffer[y][x] = append(v.buffer[y][x], &vv)
				} else {
					log.Println("Miss!!")
				}
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

func (v *Viewport) RenderSvg(scene *Scene, filename string, level int) {
	f, err := os.OpenFile(filename, os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	w.WriteString(fmt.Sprintf("<svg width=\"%v\" height=\"%v\" version=\"1.1\" baseProfile=\"full\" xmlns=\"http://www.w3.org/2000/svg\">", v.Width, v.Height))
	w.WriteString(fmt.Sprintf("<g transform=\"translate(0,%v) scale(1,-1)\">", v.Height))
	for _, m := range scene.Meshes {

		for _, p := range m.Polygons {
			var pl []string
			for _, index := range p.Indexes {
				ov := m.getVertex(index)
				wv := scene.WorldTransformation.MulPositionW(ov)
				vv := v.Transformation.MulPositionW(wv)
				pl = append(pl, fmt.Sprintf("%v,%v ", vv.X, vv.Y))
			}
			coords := strings.Join(pl, " ")
			w.WriteString(fmt.Sprintf("<polyline stroke=\"black\" fill=\"blue\" fill-opacity=\"0.5\" points=\"%v\" />", coords))
		}
	}
	w.WriteString("</g></svg>")
}
