package main

import (
	"fmt"
	"sort"
)

func main() {
	votes := make(map[int]int)
	totalMasuk := 0
	totalSah := 0

	for {
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}

		totalMasuk++

		if n >= 1 && n <= 20 {
			totalSah++
			votes[n]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", totalSah)

	// Ambil semua calon yang mendapat suara, lalu urutkan
	candidates := []int{}
	for k := range votes {
		candidates = append(candidates, k)
	}
	sort.Ints(candidates)

	for _, c := range candidates {
		fmt.Printf("%d: %d\n", c, votes[c])
	}
}