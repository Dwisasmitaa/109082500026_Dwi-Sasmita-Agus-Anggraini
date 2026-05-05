package main

import "fmt"

const nProv int = 2

type NamaProv = [2]string
type PopProv = [2]int
type TumbuhProv = [2]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("==== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ====")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	idxMaks := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idxMaks] {
			idxMaks = i
		}
	}
	return idxMaks
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("==== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ====")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := int((tumbuh[i] + 1) * float64(pop[i]))
			fmt.Printf("%s %d\n", prov[i], prediksi)
		}
	}
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv

	InputData(&prov, &pop, &tumbuh)

	
	idxTercepat := ProvinsiTercepat(tumbuh)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", prov[idxTercepat])

	var namaCari string
	fmt.Scan(&namaCari)
	idxCari := IndeksProvinsi(prov, namaCari)
	if idxCari == -1 {
		fmt.Printf("Data provinsi yang dicari : %s tidak ditemukan\n", namaCari)
	} else {
		fmt.Printf("Data provinsi yang dicari : %s\n", prov[idxCari])
	}


	Prediksi(prov, pop, tumbuh)
}