package main

import "fmt"

func main() {
	fmt.Print("Bitte eine Zahl eingeben: ")
	var number1 int
	fmt.Scanln(&number1)

	fmt.Print("Bitte eine zweite Zahl eingeben: ")
	var number2 int
	fmt.Scanln(&number2)

	fmt.Println("Die Summe der beiden Zahlen ist:", number1+number2)
}
