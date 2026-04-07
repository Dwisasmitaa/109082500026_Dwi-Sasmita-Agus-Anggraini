package main
import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	var hariDitanggung int
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			hariDitanggung = 3
		} else {
			hariDitanggung = jumlahHari
		}
	} else if tujuan == "mancanegara" {
		if jumlahHari > 8 { 
			hariDitanggung = 8
		} else {
			hariDitanggung = jumlahHari
		}
	}
	
	return hariDitanggung
}

func biayaPerHari(jumlahMhs int) int {
	
	biayaMakan := 2 * 35000 
	biayaPenginapan := 250000
	uangSaku := 300000 
	
	biayaPerMhs := biayaMakan + biayaPenginapan + uangSaku 
	
	return biayaPerMhs * jumlahMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) { 
	
	hari := tanggunganHari(lamaPerjalanan, tujuan)

	biayaDomestikHarian := biayaPerHari(jumlahMhs)

	if tujuan == "domestik" {
		*totalBiaya = float64(hari * biayaDomestikHarian)
	} else if tujuan == "mancanegara" {
		*totalBiaya = float64(hari) * float64(biayaDomestikHarian) * 1.5
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("Biaya perjalanan yang harus dikeluarkan tel-u : Rp. %.0f\n", biaya)
}