package main

import "fmt"

const NMAX = 256

type arrInt [NMAX]int

func isiData(A *arrInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&A[i])
	}
}


func reverse(A arrInt, n int, B *arrInt) {
	for i := 0; i < n; i++ {
		B[i] = A[n-1-i]
	}
}

func isPalindrom(A arrInt, n int) bool {
	for i := 0; i < n/2; i++ {
		if A[i] != A[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var A, B arrInt
	var n int

	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	isiData(&A, n)

	
	reverse(A, n, &B)


	fmt.Print("Array A: ")
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	fmt.Print("Array Reverse: ")
	for i := 0; i < n; i++ {
		fmt.Print(B[i], " ")
	}
	fmt.Println()


	if isPalindrom(A, n) {
		fmt.Println("Array merupakan Palindrom")
	} else {
		fmt.Println("Array bukan Palindrom")
	}
}