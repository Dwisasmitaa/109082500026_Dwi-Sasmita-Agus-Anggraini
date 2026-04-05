package main

import 
	"fmt"


func cetak (x int) {
	if x > 45 {
		return
	}
	fmt.Print(x)
	cetak(x + 1)
}

func main() {
	cetak(5)

}