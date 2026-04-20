package main
import "fmt"


type Mahasiswa struct {
	nama     string
	nim      string
	eprt     int
	semester int
	ipk      float64
}

func inputData(data *[1000]Mahasiswa, n *int) {
	i := 0
	for {
		fmt.Print("Nama : ")
		fmt.Scan(&data[i].nama)

		if data[i].nama == "none" {
			break
		}

		fmt.Print("Nim: ")
		fmt.Scan(&data[i].nim)

		fmt.Print("EPRT: ")
		fmt.Scan(&data[i].eprt)

		fmt.Print("Semester: ")
		fmt.Scan(&data[i].semester)

		fmt.Print("IPK: ")
		fmt.Scan(&data[i].ipk)

		i++
	}
	*n = i
}


func maxEprt(data [1000]Mahasiswa, n int) int {
	max := data[0].eprt
	for i := 1; i < n; i++ {
		if data[i].eprt > max {
			max = data[i].eprt
		}
	}
	return max
}


func minIPK(data [1000]Mahasiswa, n int) float64 {
	min := data[0].ipk
	for i := 1; i < n; i++ {
		if data[i].ipk < min {
			min = data[i].ipk
		}
	}
	return min
}

func rataSemester(data [1000]Mahasiswa, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += data[i].semester
	}
	return float64(total) / float64(n)
}

func main() {
	var data [1000]Mahasiswa
	var n int

	inputData(&data, &n)

	fmt.Println("EPRT Tertinggi:", maxEprt(data, n))
	fmt.Println("IPK Terendah:", minIPK(data, n))
	fmt.Println("Rata-rata Semester:", rataSemester(data, n))
}