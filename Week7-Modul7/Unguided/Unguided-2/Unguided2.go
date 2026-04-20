package main
import "fmt"

type kata string//tipe data alias
type angka int
type buku struct {
	//struct 
	JudulBuku kata //field
	Namapenulis kata //field
	Penerbit kata //field
	Tahunterbit angka
	Jumlahhalaman angka
}

func main (){
	var b buku
	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukkan judul buku : ")
	fmt.Scanln(&b.JudulBuku)
	fmt.Print("Masukkan nama penulis : ")
	fmt.Scanln(&b.Namapenulis)
	fmt.Print("Masukkan penerbit buku : ")
	fmt.Scanln(&b.Penerbit)
	fmt.Print("Masukkan tahun terbit : ")
	fmt.Scanln(&b.Tahunterbit)
	fmt.Print("Masukkan jumlah halaman : ")
	fmt.Scanln(&b.Jumlahhalaman)


	fmt.Println("=== DATA BUKU ===")
	fmt.Println("judul buku:",b.JudulBuku)
	fmt.Println("nama penulis:",b.Namapenulis)
	fmt.Println("penerbit buku:",b.Penerbit)
	fmt.Println("tahun terbit:",b.Tahunterbit)
	fmt.Println("jumlah halaman:",b.Jumlahhalaman)
}
  


