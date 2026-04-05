package main
import "fmt"

func tampilGanjil(n, current int) {
    if current > n {
        return
    }
	   
    if current%2 == 1 {
        fmt.Printf("%d ", current)
    }
  
    tampilGanjil(n, current+1)
}

func main() {
    var N int
    fmt.Print("Masukkan N: ")
    fmt.Scan(&N)

    tampilGanjil(N, 1)
    fmt.Println() 
}