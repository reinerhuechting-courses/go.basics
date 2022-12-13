package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	balance := 20

	for balance != 0 {
		fmt.Printf("Dein Kontostand: %d\n", balance)
		playerNumber := GetPlayerNumber()
		playerBet := GetPlayerBet()
		wheelNumber := SpinWheel()
		balance += EvaluateRound(playerNumber, wheelNumber, playerBet)
		fmt.Println()
	}
}

func SpinWheel() int {
	return rand.Intn(37)
}

func NumberAccepted(number int) bool {
	return number >= 0 && number <= 36
}

func AmountValid(amount int) bool {
	return amount > 0
}

func ReadNumber(message string) int {
	fmt.Printf("%s ", message)
	var number int
	fmt.Scanln(&number)
	return number
}

func GetPlayerNumber() int {
	for {
		number := ReadNumber("Bitte wähle eine Zahl:")
		if NumberAccepted(number) {
			return number
		}
		fmt.Println("Das war keine gültige Zahl für's Roulette.")
		fmt.Println("Bitte probiere es noch einmal!")
	}
}

func GetPlayerBet() int {
	for {
		amount := ReadNumber("Wie viel möchtest du setzen?")
		if AmountValid(amount) {
			return amount
		}
		fmt.Println("Du kannst keinen negativen Betrag setzen.")
		fmt.Println("Bitte probiere es noch einmal!")
	}
}

func EvaluateRound(playerNumber, wheelNumber, playerBet int) int {
	fmt.Printf("Du hast auf %d gesetzt, gedreht wurde %d.\n", playerNumber, wheelNumber)

	if playerNumber == wheelNumber {
		fmt.Println("Herzlichen Glückwunsch, du hast gewonnen!")
		return playerBet * 35
	}
	fmt.Println("Du hast leider verloren.")
	return -playerBet
}
