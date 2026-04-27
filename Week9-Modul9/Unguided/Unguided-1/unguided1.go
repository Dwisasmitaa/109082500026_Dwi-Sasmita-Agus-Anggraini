package main
import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat  Titik
	radius int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func posisiTitik(c1, c2 Lingkaran, p Titik) string {
	dalam1 := didalam(c1, p)
	dalam2 := didalam(c2, p)

	switch {
	case dalam1 && dalam2:
		return "Titik di dalam lingkaran 1 dan 2"
	case dalam1:
		return "Titik di dalam lingkaran 1"
	case dalam2:
		return "Titik di dalam lingkaran 2"
	default:
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {
	var c1, c2 Lingkaran
	var p Titik

	// Input lingkaran 1: pusat (cx1, cy1) dan radius r1
	fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.radius)
	// Input lingkaran 2: pusat (cx2, cy2) dan radius r2
	fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.radius)
	// Input titik sembarang
	fmt.Scan(&p.x, &p.y)

	fmt.Println(posisiTitik(c1, c2, p))
}