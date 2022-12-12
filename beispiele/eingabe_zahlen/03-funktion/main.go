package main

import "fmt"

func main() {
	number1 := ReadNumber()
	number2 := ReadNumber()
	number3 := ReadNumber()

	fmt.Println("Die Summe der Zahlen ist:", number1+number2+number3)
}

func ReadNumber() int {
	fmt.Print("Bitte eine Zahl eingeben: ")
	var number int
	fmt.Scanln(&number)
	return number
}
