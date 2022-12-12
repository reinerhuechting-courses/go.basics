package main

import "fmt"

func main() {
	fmt.Println("Die Summe der Zahlen ist:", ReadAndAdd(5))
}

func ReadAndAdd(count int) int {
	sum := 0
	for i := 0; i < count; i++ {
		sum += ReadNumber()
	}
	return sum
}

func ReadNumber() int {
	fmt.Print("Bitte eine Zahl eingeben: ")
	var number int
	fmt.Scanln(&number)
	return number
}
