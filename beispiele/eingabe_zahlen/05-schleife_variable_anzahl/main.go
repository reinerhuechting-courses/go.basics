package main

import "fmt"

func main() {
	sum := 0
	count := 5

	for i := 0; i < count; i++ {
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
