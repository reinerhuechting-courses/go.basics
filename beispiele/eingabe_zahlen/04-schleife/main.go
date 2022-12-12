package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i <= 3; i++ {
		number := ReadNumber()
		sum += number
	}

	fmt.Println("Die Summe der Zahlen ist:", sum)
}

func ReadNumber() int {
	fmt.Print("Bitte eine Zahl eingeben: ")
	var number int
	fmt.Scanln(&number)
	return number
}
