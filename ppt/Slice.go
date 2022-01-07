package main

import "fmt"

func main() {
	//LONG DECLARATION
	var event_numbers []int
	fmt.Printf("elements = %v, len = %d, cap=%d\n", event_numbers, len(event_numbers), cap(event_numbers))

	//LONG DECLARATION WITH VALUE
	var odd_numbers = []int{1, 3, 5, 7, 9}
	fmt.Printf("elements = %v, len = %d, cap=%d\n", odd_numbers, len(odd_numbers), cap(odd_numbers))

	//SHORT DECLARATION WITH VALUES
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("elements = %v, len = %d, cap=%d\n", numbers, len(numbers), cap(numbers))

	//USING MAKE FUNCTION
	var primes = make([]int, 5, 10)
	fmt.Printf("element= %v, len=%d, cap=%d\n", primes, len(primes), cap(primes))
}
