package main

import "fmt"

func main() {
	fmt.Print("Bitte eine Zahl eingeben: ")
	var number int
	fmt.Scanln(&number)
	fmt.Println("Die Zahl war:", number)
}
