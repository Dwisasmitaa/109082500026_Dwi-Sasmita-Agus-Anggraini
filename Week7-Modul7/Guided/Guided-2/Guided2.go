package main
import "fmt"

type kalimat string//tipe data alias

type mahasiswa struct {
	//struct 
	nama kalimat //field
	nim int
	nilai float64

}

func main () {
	var m mahasiswa

	fmt.Print("Masukkan nama mahasiswa: ")
	fmt.Scanln(&m.nama)
	fmt.Print("Masukkan nim : ")
	fmt.Scanln(&m.nim)
	fmt.Print("Masukkan nilai : ")
	fmt.Scanln(&m.nilai)

	fmt.Println("DATA MAHASISWA: ")
	fmt.Println("nama:",m.nama)
	fmt.Println("nim:",m.nim)
	fmt.Println("nilai:",m.nilai)
}