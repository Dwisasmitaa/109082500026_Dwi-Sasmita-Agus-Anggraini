package main
import "fmt"

func main() {
	var klubA, klubB string
	var menangs []string

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	pertandingan := 1
	for {
		var skorA, skorB int

		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		var hasil string
		if skorA > skorB {
			hasil = klubA
		} else if skorB > skorA {
			hasil = klubB
		} else {
			hasil = "Draw"
		}

		if hasil != "Draw" {
			menangs = append(menangs, hasil)
		} else {
			menangs = append(menangs, "Draw")
		}

		fmt.Printf("Hasil %d : %s\n", pertandingan, hasil)
		pertandingan++
	}

	fmt.Println("Pertandingan selesai")
	fmt.Println()
	fmt.Println("Daftar klub yang menang:")
	for i, m := range menangs {
		if m != "Draw" {
			fmt.Printf("  Pertandingan %d : %s\n", i+1, m)
		}
	}
}