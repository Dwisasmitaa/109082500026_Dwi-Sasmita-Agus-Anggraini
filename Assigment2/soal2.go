package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa = [51]mahasiswa


func cariNilaiPertama(arr arrayMahasiswa, N int, nim string) int {
	for i := 0; i < N; i++ {
		if arr[i].NIM == nim {
			return arr[i].nilai 
		}
	}
	return -1 
}
func cariNilaiTerbesar(arr arrayMahasiswa, N int, nim string) int {
	maks := -1
	for i := 0; i < N; i++ {
		if arr[i].NIM == nim {
			if arr[i].nilai > maks {
				maks = arr[i].nilai
			}
		}
	}
	return maks
}

func main() {
	var arr arrayMahasiswa
	var N int

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&arr[i].NIM, &arr[i].nama, &arr[i].nilai)
	}

	var nimCari string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	nilaiPertama := cariNilaiPertama(arr, N, nimCari)
	nilaiTerbesar := cariNilaiTerbesar(arr, N, nimCari)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimCari, nilaiPertama)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimCari, nilaiTerbesar)
}