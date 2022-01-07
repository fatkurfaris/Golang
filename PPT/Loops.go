package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	//DO WHILE

	sum := 1

	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}
