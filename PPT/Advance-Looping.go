package main

import "fmt"

func main() {
	N := 5

	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
