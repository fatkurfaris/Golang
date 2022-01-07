package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func greeting(hour int) {
	if hour < 12 {
		fmt.Println("selamat pagi")
	} else if hour < 18 {
		fmt.Println("selamat Sore")
	} else {
		fmt.Println("selamat malam")
	}

}

func main() {
	hour := 14
	greeting(hour)
	sayHello()
}
