package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"net/http"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/gonutz/framebuffer"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

func main() {
	// flag.Parse()
	// flags := flag.Args()

	// fmt.Printf("%v\n", flags[0])

	srv := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := strings.Split(r.URL.Path, "/")

		img := image.NewRGBA(image.Rect(0, 0, 320, 240))
		draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.ZP, draw.Src)
		x, y := 0, 40
		addLabel(img, x, y, s[1])

		fb, err := framebuffer.Open("/dev/fb1")
		if err != nil {
			panic(err)
		}
		defer fb.Close()

		draw.Draw(fb, fb.Bounds(), img, image.ZP, draw.Src)
	})
	fmt.Print(srv.ListenAndServe())
}

// https://stackoverflow.com/questions/38299930/how-to-add-a-simple-text-label-to-an-image-in-go
func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	ff, _ := truetype.Parse(goregular.TTF)
	face := truetype.NewFace(ff, &truetype.Options{Size: 48})

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  point,
	}
	d.DrawString(label)
}
