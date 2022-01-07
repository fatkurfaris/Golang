package main

import "fmt"

func main() {
	primes := [5]int{2, 3, 5}
	// fmt.Println(primes)

	for index := 0; index < len(primes); index++ {
		fmt.Println(primes[index])
	}
}
