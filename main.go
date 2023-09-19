package main

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"log"
	"math/rand"
)

func getColor(r int, g int, b int, a int) color.RGBA {
	return color.RGBA{
		R: uint8(r % 256),
		G: uint8(g % 256),
		B: uint8(b % 256),
		A: uint8(a % 256),
	}
}

func getColorRandom() color.Color {
	return getColor(
		rand.Intn(8)*32,
		rand.Intn(8)*32,
		rand.Intn(8)*32,
		255)
}

func main() {
	const (
		imgWidth  = 10240
		imgHeight = 10240
	)

	img := imaging.New(imgWidth, imgHeight, image.Black)

	d := map[int]color.Color{}

	for px := 0; px < imgWidth; px++ {
		log.Println(float64(px) / float64(imgWidth) * 100.0)
		for py := 0; py < imgHeight; py++ {
			sx := float64(px)/float64(imgWidth)*4.0 - 2.0
			sy := float64(py)/float64(imgHeight)*4.0 - 2.0
			vx := sx
			vy := sy

			itMax := 200

			for it := 0; it < itMax; it++ {
				if vx*vx+vy*vy > 4 {
					if _, ok := d[it]; !ok {
						d[it] = getColorRandom()
					}

					img.Set(px, py, d[it])
					
					break
				}

				newVx := vx*vx - vy*vy + sx
				newVy := 2*vx*vy + sy
				vx, vy = newVx, newVy
			}
		}
	}

	log.Println(d)

	_ = imaging.Save(img, "output.png")
}
