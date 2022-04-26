package main

import "fmt"

func main() {
	first := 0
	second := 0
	tempVal := 0
	fmt.Println("1st number")
	fmt.Scanln(&first)
	fmt.Println("2nd number")
	fmt.Scanln(&second)
	if second < first {
		tempVal = second
		second = first
		first = tempVal

	}
	for i := first; i <= second; i++ {
		if i%2 == 1 {
			fmt.Print("\nThis number is odd ", i)
		}
	}
	for i := first; i <= second; i++ {
		if i%2 == 0 {
			fmt.Print("\nThis number is even ", i)
		}
	}
}
