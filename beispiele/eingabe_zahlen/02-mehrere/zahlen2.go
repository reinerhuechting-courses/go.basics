package main

import "fmt"

func main() {
	fmt.Print("Bitte eine Zahl eingeben: ")
	var number1 int
	fmt.Scanln(&number1)

	fmt.Print("Bitte eine Zahl eingeben: ")
	var number2 int
	fmt.Scanln(&number2)

	fmt.Print("Bitte eine Zahl eingeben: ")
	var number3 int
	fmt.Scanln(&number3)

	fmt.Println("Die Summe der Zahlen ist:", number1+number2+number3)
}
