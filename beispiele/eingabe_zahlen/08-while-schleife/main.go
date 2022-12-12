package main

import "fmt"

func main() {
	fmt.Println("Die Summe der Zahlen ist:", ReadAndAdd())
}

func ReadAndAdd() int {
	sum := 0
	number := 0
	for number != 38 {
		number = ReadNumber()
		sum += number
	}
	return sum
}

func ReadNumber() int {
	fmt.Print("Bitte eine Zahl eingeben (Beenden mit 38): ")
	var number int
	fmt.Scanln(&number)
	return number
}
