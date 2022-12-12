package main

import "fmt"

func main() {
	sum := ReadAndAdd(5)
	fmt.Println("Die Summe der Zahlen ist:", sum)
}

func ReadAndAdd(count int) int {
	sum := 0
	for i := 0; i < count; i++ {
		number := ReadNumber()
		sum += number
	}
	return sum
}

func ReadNumber() int {
	fmt.Print("Bitte eine Zahl eingeben: ")
	var number int
	fmt.Scanln(&number)
	return number
}
