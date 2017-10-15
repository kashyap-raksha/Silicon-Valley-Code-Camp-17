package main

import "fmt"

func main() {
	var numOne int
	var numTwo int
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&numOne)
	fmt.Print("Please enter a smaller number: ")
	fmt.Scan(&numTwo)
	arithmetic_ops(numOne, numTwo)
}

func arithmetic_ops(numOne int, numTwo int) {
	fmt.Println("Remainder:", numOne, "%", numTwo, " = ", numOne%numTwo)
	fmt.Println("Division:", numOne, "%", numTwo, "=", numOne/numTwo)
	fmt.Println("Addition:", numOne, "%", numTwo, "=", numOne+numTwo)
	fmt.Println("Subtraction:", numOne, "%", numTwo, "=", numOne-numTwo)
	fmt.Println("Multiplication:", numOne, "%", numTwo, "=", numOne/numTwo)
	return
}
