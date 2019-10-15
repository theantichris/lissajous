package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.RGBA{0x00, 0x00, 0x00, 0x00}, color.RGBA{0x00, 0xff, 0x00, 0x00}}

const (
	blackIndex = 0
	greenIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles         = 5     // number of complete x oscillator revolutions
		res            = 0.001 // angular resolution
		size           = 100   // image canvas covers [-size..+size]
		numberOfFrames = 64    // number of frames in the animation
		delay          = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	animation := gif.GIF{LoopCount: numberOfFrames}
	phase := 0.0 // phase difference

	for i := 0; i < numberOfFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}

		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, img)
	}

	gif.EncodeAll(out, &animation) // NOTE: ignoring encoding errors
}
