package main

import "fmt"

type geometry struct {
	area, perimeter int
}

type rectangle struct {
	length, width int
	color         string
	property      geometry
}

func isiData(persegi *rectangle) {
	fmt.Scan(&persegi.length, &persegi.width, &persegi.color)
}

func hitung(persegi *rectangle) {
	persegi.property.area = persegi.length * persegi.width
	persegi.property.perimeter = 2 * (persegi.length + persegi.width)
}

func main() {
	var persegi rectangle

	isiData(&persegi)
	hitung(&persegi)

	fmt.Println(persegi.property.area)
	fmt.Println(persegi.property.perimeter)
}
