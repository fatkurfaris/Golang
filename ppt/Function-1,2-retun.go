package main

import (
	"fmt"
	"math"
)

//SINGLE RETURN VALUE
func calculateSquare(side int) int {
	return side * side
}

func calculateCircle(diameter float64) (float64, float64) {
	var keliling = math.Pi * math.Pow(diameter/2, 2)
	var luas = math.Pi * diameter
	return keliling, luas
}

func main() {
	var side = 5
	wide := calculateSquare(side)
	fmt.Printf("Lias Persegi empat : %d \n\n", wide)

	var diameter float64 = 15
	keliling, luas := calculateCircle(diameter)
	fmt.Printf("Keliling Lingkaran : %.2f \n", keliling)
	fmt.Printf("Luas Lingkaran : %.2f \n", luas)
}
