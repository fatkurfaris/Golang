package main

import "fmt"

func main() {
	// var myAge int = 17

	// if myAge == 5 {
	// 	fmt.Println("You too young")
	// }
	// if myAge == 17 {
	// 	fmt.Println("so sweet")
	// }

	// if myAge >= 17 && myAge <= 30 {
	// 	fmt.Println("my age is between 17 and 30")
	// }

	// if dadAge := 9; dadAge < myAge {
	// 	fmt.Println(dadAge)
	// }

	// number := 5

	// if number%2 == 0 {
	// 	fmt.Println("Genap")
	// } else {
	// 	fmt.Println("Ganjil")
	// }

	hour := 15
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("selamat Sore")
	} else {
		fmt.Println("selamat Malam")
	}
}
