package main

import "fmt"

const BMAX = 1000000

var data [BMAX]int

func isArray() int {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
	return n
}

func posisi(n, k int) int {
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		if data[mid] == k {
			return mid
		} else if data[mid] < k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	var n, k int

	fmt.Scan(&n, &k)


	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	pos := posisi(n, k)

	if pos == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(pos)
	}
}