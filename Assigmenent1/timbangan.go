package main
import "fmt"

const pi float64 = 3.14

func volume(r, t float64) float64 {
	return pi * r * r * t
}

func massa(r, t, p float64) float64 {
	return volume(r, t) * p
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else {
		selisih := m1 - m2
		if selisih < 0 {
			selisih = -selisih
		}
		fmt.Println("Selisih massa zat cair kiri dan massa zat cair kanan :", selisih)
	}
}

func main() {
	var r float64
	var tKiri, tKanan float64
	var mjKiri, mjKanan float64
	var massaKiri, massaKanan float64

	fmt.Print("Masukkan Jari-jari : ")
	fmt.Scan(&r)

	fmt.Print("Masukkan Tinggi zat cair tabung kiri : ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukkan Massa jenis zat cair tabung kiri : ")
	fmt.Scan(&mjKiri)

	fmt.Print("Masukkan Tinggi zat cair tabung kanan : ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukkan Massa jenis zat cair tabung kanan : ")
	fmt.Scan(&mjKanan)

	massaKiri = massa(r, tKiri, mjKiri)
	massaKanan = massa(r, tKanan, mjKanan)

	display(massaKiri, massaKanan)
}