package main

import "fmt"

func penjumlahan(a, b, c int) int {
	var hasil int
	hasil = a + b + c
	return hasil

}

func main() {
	fmt.Println(penjumlahan(10, 5, 15))
}