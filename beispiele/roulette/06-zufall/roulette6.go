package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	playerNumber := GetPlayerNumber()
	wheelNumber := SpinWheel()
	PrintResult(playerNumber, wheelNumber)
}

func SpinWheel() int {
	return rand.Intn(37)
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

func GetPlayerNumber() int {
	for {
		number := ChooseNumber()
		if NumberAccepted(number) {
			return number
		}
		fmt.Println("Das war keine gültige Zahl für's Roulette.")
		fmt.Println("Bitte probiere es noch einmal!")
	}
}

func PrintResult(playerNumber, wheelNumber int) {
	fmt.Printf("Du hast %d gesetzt, gedreht wurde %d.\n",
		playerNumber, wheelNumber)
	if playerNumber == wheelNumber {
		fmt.Println("Herzlichen Glückwunsch, du hast gewonnen!")
	} else {
		fmt.Println("Du hast leider verloren.")
	}
}
