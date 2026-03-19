package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(1)
}

func main() {
	var input int
	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&input)

	if input > 0 && input < 1000000 {
		cetakDeret(input)
	} else {
		fmt.Println("Input harus antara 1 sampai 999.999")
	}
}