package main

import "fmt"

func main() {
	sentence := "Hello"

	for i := 0; i < len(sentence); i++ {
		fmt.Println(string(sentence[i]))
	}
}
