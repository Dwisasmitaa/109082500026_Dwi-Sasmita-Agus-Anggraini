package main
import "fmt"

func LuasTabung(r, t float64)float64 {
	var hasil float64
	hasil = 2 * 3.14 * r * (r + t ) 
	return hasil

	
}

func main (){
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Print(a, b)

	fmt.Print(LuasTabung(a, b))

}