package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/golang/freetype/truetype"
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
		x, y := 0, 40
		addLabel(img, x, y, s[1])
		buf := new(bytes.Buffer)
		png.Encode(buf, img)

		file, _ := os.Create("img.png")
		io.Copy(file, buf)

		execute()
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

// https://stackoverflow.com/questions/6182369/exec-a-shell-command-in-go
func execute() {
	cmd := exec.Command("fbi", "-T", "2", "-d", "/dev/fb1", "-noverbose", "-a", "img.png")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
