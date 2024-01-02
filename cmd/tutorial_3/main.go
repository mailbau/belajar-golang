package main

import "fmt"

func main() {
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder int = intDivision(numerator, denominator)
	fmt.Printf("The result of division is %v with remainder %v", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}
