package main

import "fmt"

func main() {
	var today int = 4

	switch today {
	case 1:
		fmt.Println("today is monday")
	case 2:
		fmt.Println("Today is Tuesday")
	case 3:
		fmt.Println("Today is Wednesday")
	case 4:
		fmt.Println("Today is Thursday")
	case 5:
		fmt.Println("Today is Friday")
	case 6:
		fmt.Println("Today is Saturday")
	case 7:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Invalid Date")
	}
}
