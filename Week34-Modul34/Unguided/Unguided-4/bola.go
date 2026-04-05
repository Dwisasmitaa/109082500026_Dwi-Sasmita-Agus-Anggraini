package main

import "fmt"

type tabGol []int

func inputData(t *tabGol, n *int) {
	var gol int

	*n = 0
	*t = make(tabGol, 0)

	for {
		fmt.Scan(&gol)
		if gol < 0 {
			break
		}
		*t = append(*t, gol)
		*n++
	}
}

func rataan(t tabGol, n int) float64 {
	var i, total int

	total = 0

	if n == 0 {
		return 0.0
	}

	for i = 0; i < n; i++ {
		total += t[i]
	}
	return float64(total) / float64(n)
}

func main() {
	var t1, t2 tabGol
	var n1, n2 int
	var r1, r2 float64

	inputData(&t1, &n1)
	inputData(&t2, &n2)

	r1 = rataan(t1, n1)
	r2 = rataan(t2, n2)

	fmt.Printf("%.2f %.2f\n", r1, r2)
}
