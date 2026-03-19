package main
import "fmt"

func main() {
	var berat, kg, sisa, biayaKg, biayaSisa, total int
	fmt.Print("Masukkan berat :")
	fmt.Scan(&berat)


	kg = berat / 1000
	sisa = berat % 1000

	biayaKg = kg * 10000
	biayaSisa = 0

	if kg > 10 {
		biayaSisa = 0
	} else if sisa >= 500 {
		biayaSisa = sisa * 5
	} else {
		biayaSisa = sisa * 15
	}

	total = biayaKg + biayaSisa

	fmt.Println("===== Detail Perhitungan =====")
	fmt.Println("Detail berat :", kg, "kg +", sisa, "gram")
	fmt.Println("Detail biaya : Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp", total)
	fmt.Println()
}




