package main

import "fmt"

func akar(x float64) float64 {
	var i int
	var z float64

	z = x / 2

	for i = 0; i < 100; i++ {
		z = z - (z*z-x)/(2*z)
	}

	return z
}

func jarak(x1, y1, x2, y2 float64) float64 {
	var dx, dy float64

	dx = x1 - x2
	dy = y1 - y2

	return akar((dx * dx) + (dy * dy))
}

func main() {
	var x1, y1, x2, y2 float64

	fmt.Println("Input x1, y2, x2, y2 :")
	fmt.Scan(&x1, &y1, &x2, &y2)

	fmt.Printf("%.2f\n", jarak(x1, y1, x2, y2))
}
