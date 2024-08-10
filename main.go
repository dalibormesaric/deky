package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/gonutz/framebuffer"
	"github.com/stianeikeland/go-rpio/v4"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/math/fixed"
)

type button struct {
	id   int
	down bool
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := strings.Split(r.URL.Path, "/")
		paint(s[1])
	})

	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rpio.Close()

	msg := make(chan button)

	pinId, _ := strconv.Atoi(os.Args[1])
	log.Println(pinId)
	go pin1(uint8(pinId), msg)
	// go pin1()

	go painter(msg)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func painter(msg chan button) {
	for {
		btn := <-msg
		if btn.down {
			paint(fmt.Sprintf("%d%s", btn.id, "down"))
		} else {
			paint(fmt.Sprintf("%d%s", btn.id, "up"))
		}
	}
}

func pin1(pinId uint8, msg chan button) {
	pin := rpio.Pin(pinId)
	pin.Input()
	state := ""
	for {
		time.Sleep(time.Second / 2)
		res := pin.Read()
		if res == rpio.High && state != "high" {
			state = "high"
			msg <- button{id: 1, down: false}
		}
		if res == rpio.Low && state != "low" {
			state = "low"
			msg <- button{id: 1, down: true}
		}
	}
}

func pin2() {
	pin := rpio.Pin(22)
	pin.Input()
	// pin.PullUp()
	// pin.Detect(rpio.FallEdge)
	state := ""
	for {
		// if pin.EdgeDetected() && state != "high" { // check if event occured
		// 	state = "high"
		// 	paint("high")
		// } else if state != "low" {
		// 	state = "low"
		// 	paint("low")
		// }
		time.Sleep(time.Second / 2)
		res := pin.Read()
		if res == rpio.High && state != "high" {
			state = "high"
			paint("2")
		}
		if res == rpio.Low && state != "low" {
			state = "low"
			paint("low")
		}
	}
}

func paint(text string) {
	img := image.NewRGBA(image.Rect(0, 0, 320, 240))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)
	x, y := 0, 40
	addLabel(img, x, y, text)

	fb, err := framebuffer.Open("/dev/fb1")
	if err != nil {
		panic(err)
	}
	defer fb.Close()

	draw.Draw(fb, fb.Bounds(), img, image.Point{}, draw.Src)
}

// https://stackoverflow.com/questions/38299930/how-to-add-a-simple-text-label-to-an-image-in-go
func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	ff, _ := truetype.Parse(gomono.TTF)
	face := truetype.NewFace(ff, &truetype.Options{Size: 48})

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  point,
	}
	d.DrawString(label)
}
