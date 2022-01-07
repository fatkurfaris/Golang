package main

import "fmt"

func multiplication(a, b int) (mul int) {
	mul = a * b
	return
}

func main() {
	m := multiplication(5, 5)
	fmt.Println("5x5 = ", m)
}
