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

	candidates := []int{}
	for k := range votes {
		candidates = append(candidates, k)
	}


	sort.Slice(candidates, func(i, j int) bool {
		if votes[candidates[i]] != votes[candidates[j]] {
			return votes[candidates[i]] > votes[candidates[j]] 
		}
		return candidates[i] < candidates[j] 
	})

	ketua := candidates[0]
	wakil := candidates[1]

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}