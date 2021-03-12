package main

import "golang.org/x/tour/pic"
import "math"

const maxiter = 1000

func Pic(dx, dy int) (img [][]uint8) {
	img = make([][]uint8, dy)
	for i := range img {
		img[i] = make([]uint8, dx)
	}

	fw := float64(dx)
	fh := float64(dy)

	ecx := fw / 2
	ecy := fh / 4

	rx := math.Pow(fw/5, 4)
	ry := math.Pow(fh/10, 4)

	for y := range img {
		for x := range img[y] {
			img[y][x] = 255

			fy := float64(y)
			fx := float64(x)

			if mandelbrot(-0.75+1.3*fy/fh, 1.0-2.0*fx/fw) {
				img[y][x] = 0
			}

			if math.Pow(fx-ecx, 4)/rx+math.Pow(fy-ecy, 4)/ry < 1 {
				img[y][x] = uint8(255 - (fx+fy)/2)
			}
		}
	}

	return
}

func mandelbrot(a, b float64) bool {
	x := 0.0
	y := 0.0

	for i := 0; i <= maxiter && x*x+y*y < 4; i++ {
		tx := x
		x = x*x - y*y + a
		y = 2*tx*y + b
	}

	return x*x+y*y < 4
}

func main() {
	pic.Show(Pic)
}
