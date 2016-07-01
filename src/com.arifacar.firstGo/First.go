package main

import "fmt"

var name = "Arif Acar" // private
var School = "Marmara University"  // public

func multiply(x int, y int) int {
	return x * y;
}

func multipleReturn(x, y, z int) (int, int) {
	a := x + y;
	var b = x + z;
	return a, b
}

func main() {
	fmt.Println("Merhaba Go ve " + name + "\n")
	fmt.Println("School " + School + "\n")

	fmt.Println("Sonuç => ", multiply(2, 5))

	var a, b = multipleReturn(2, 3, 4)

	fmt.Println("Multiple Return Result => ", a, b)

	var empty int;
	fmt.Println("Empty Result => ", empty)

}