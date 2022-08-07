package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jsadowyj0/CIS167_FinalProject/ansi"
)

func clear() {
	// clears screen with ANSI escapes
	fmt.Print("\033[H\033[2J")
}

func isValidInput(input string) bool {
	if _, err := strconv.Atoi(input); err == nil {
		parsedInput, _ := strconv.Atoi(input)
		return parsedInput >= 0 && parsedInput <= 8
	}
	return false
}

func isValidBoolean(input string) bool {
	_, err := strconv.ParseBool(input)
	return err == nil
}

func main() {
	var input string
	gameOver := false
	board := Board{}
	board.New()
	board.Print()
	fmt.Print(ansi.Bold("Press enter to start..."))
	fmt.Scanln(&input)
	for !gameOver {
		var msg string
		// X's turn
		board.PlayerTurn()
		gameOver, msg = board.CheckGameOver()
		if !gameOver {
			// O's turn
			board.Print()
			fmt.Print(ansi.Bold(ansi.Green("Thinking...")))
			time.Sleep(time.Second / 3)
			board.PlaceO(generateMaximizerMove(board))
			gameOver, msg = board.CheckGameOver()
		}

		if gameOver {
			input = ""
			for !isValidBoolean(input) {
				board.Print()
				fmt.Println(board.GenerateMessage(msg))
				fmt.Printf("Play again? [%s/%s]: ", ansi.Red("0"), ansi.Green("1"))
				fmt.Scanln(&input)
			}
			choice, _ := strconv.ParseBool(input)
			gameOver = !choice
			board.New()
		}
	}
}
