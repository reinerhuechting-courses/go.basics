package main

import "fmt"

func main() {
	playerNumber := ChooseNumber()
	if !NumberAccepted(playerNumber) {
		fmt.Println("Das war keine gültige Zahl für's Roulette.")
		fmt.Println("Bitte verlasse jetzt das Casino!")
		return
	}

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

func NumberAccepted(number int) bool {
	return number >= 0 && number <= 36
}

func ChooseNumber() int {
	fmt.Print("Bitte wähle eine Zahl: ")
	var number int
	fmt.Scanln(&number)
	return number
}
