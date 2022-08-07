package main

import (
	"fmt"
	"strconv"

	"github.com/jsadowyj0/CIS167_FinalProject/ansi"
)

type Board struct {
	cells []string
	xBits uint
	oBits uint
}

func (b *Board) New() *Board {
	b.cells = []string{
		ansi.Faint("0"), ansi.Faint("1"), ansi.Faint("2"),
		ansi.Faint("3"), ansi.Faint("4"), ansi.Faint("5"),
		ansi.Faint("6"), ansi.Faint("7"), ansi.Faint("8")}
	b.xBits = 0
	b.oBits = 0
	return b
}

func (b Board) Copy() Board {
	newBoard := Board{}
	newBoard.cells = append(newBoard.cells, b.cells...)
	newBoard.xBits = b.xBits
	newBoard.oBits = b.oBits
	return newBoard
}

func (b Board) Print() {
	clear()
	fmt.Printf(" %s-%s-%s\n", ansi.Red("Tic"), ansi.Green("Tac"), ansi.Blue("Toe"))
	fmt.Println("+-----------+")
	fmt.Printf("| %s | %s | %s |\n", b.cells[0], b.cells[1], b.cells[2])
	fmt.Println("|-----------|")
	fmt.Printf("| %s | %s | %s |\n", b.cells[3], b.cells[4], b.cells[5])
	fmt.Println("|-----------|")
	fmt.Printf("| %s | %s | %s |\n", b.cells[6], b.cells[7], b.cells[8])
	fmt.Println("+-----------+")
}

func (b *Board) isOpen(index int) bool {
	// bitwise wizardry
	return (b.xBits|b.oBits)&(1<<index) == 0
}

func (b *Board) PlaceX(index int) bool {
	if b.isOpen(index) {
		b.cells[index] = ansi.Red("X")
		b.xBits += 1 << index
		return true
	}
	return false
}

func (b *Board) PlaceO(index int) bool {
	if b.isOpen(index) {
		b.cells[index] = ansi.Green("O")
		b.oBits += 1 << index
		return true
	}
	return false
}

// Tic-Tac-Toe
// +-----------+
// | X | O | X |
// |-----------|
// | X | O | 5 |
// |-----------|
// | X | O | O |
// +-----------+

// +-----------+
// | 2^0 | 2^1 | 2^2 |
// |-----------|
// | 2^3 | 2^4|  2^5 |
// |-----------|
// | 2^6 | 2^7 | 2^8 |
// +-----------+
// x: 0b001001101
// o: 0b110010010
// m: 0b111101111

func (b *Board) ResetCell(index int) {
	b.cells[index] = ansi.Faint(strconv.Itoa(index))
	// just some sneaky bitmasking
	b.oBits &= ^(1 << index)
	b.xBits &= ^(1 << index)
}

func (b *Board) PlayerTurn() {
	var input string
	validInput := false
	for !validInput {
		clear()
		b.Print()
		fmt.Print(ansi.Bold(ansi.Red("[0-8]: ")))
		fmt.Scanln(&input)
		validInput = isValidInput(input)
		if validInput {
			validInput = false
			parsedInput, _ := strconv.Atoi(input)
			validInput = b.PlaceX(parsedInput)
		}
	}
}

func (b Board) GenerateMessage(msg string) string {
	if msg == "X" {
		return ansi.Bold(ansi.Red("X wins!"))
	} else if msg == "O" {
		return ansi.Bold(ansi.Green("O wins!"))
	} else {
		return ansi.Bold(ansi.Yellow("Draw!"))
	}
}

func (b *Board) CheckGameOver() (bool, string) {
	wins := []uint{
		// Rows
		0b000000111,
		0b000111000,
		0b111000000,
		// Columns
		0b100100100,
		0b010010010,
		0b001001001,
		// Diagonals
		0b100010001,
		0b001010100}

	for _, win := range wins {
		if b.xBits&win == win {
			return true, "X"
		} else if b.oBits&win == win {
			return true, "O"
		} else if b.xBits|b.oBits == 0b111111111 {
			return true, ""
		}
	}
	return false, ""
}
