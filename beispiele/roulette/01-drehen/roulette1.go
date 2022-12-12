package main

import "fmt"

func main() {
	playerNumber := ChooseNumber()
	wheelNumber := SpinWheel()

	if playerNumber == wheelNumber {
		fmt.Println("Herzlichen Glückwunsch, du hast gewonnen!")
	} else {
		fmt.Println("Du hast leider verloren.")
	}
}

func SpinWheel() int {
	return 15
}

func ChooseNumber() int {
	fmt.Print("Bitte wähle eine Zahl: ")
	var number int
	fmt.Scanln(&number)
	return number
}
