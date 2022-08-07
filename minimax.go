package main

import (
	"math"
)

// fmt.Printf("------------------------------------\n")
// copy.Print()
// fmt.Printf("max: x:%09b o:%09b score:%f depth: %d\n", copy.xBits, copy.oBits, bestMove, depth)

// fmt.Printf("------------------------------------\n")
// copy.Print()
// fmt.Printf("min: x:%09b o:%09b score:%f depth: %d\n", copy.xBits, copy.oBits, bestMove, depth)

func minimax(board Board, depth int, isMaximizing bool) float64 {
	copy := board.Copy()
	gameOver, msg := copy.CheckGameOver()
	if gameOver {
		if msg == "O" {
			return 1
		} else if msg == "X" {
			return -1
		} else {
			return 0
		}
	}
	if isMaximizing {
		bestMove := math.Inf(-1)
		for i := range copy.cells {
			if copy.isOpen(i) {
				copy.PlaceO(i)
				bestMove = math.Max(bestMove, float64(minimax(copy, depth+1, false)))
				copy.ResetCell(i)
			}
		}
		return bestMove
	} else {
		bestMove := math.Inf(1)
		for i := range copy.cells {
			if copy.isOpen(i) {
				copy.PlaceX(i)
				bestMove = math.Min(bestMove, float64(minimax(copy, depth+1, true)))
				copy.ResetCell(i)
			}
		}
		return bestMove
	}
}

func generateMaximizerMove(board Board) int {
	bestScore := math.Inf(-1)
	bestMove := 0
	copy := board.Copy()
	for i := range copy.cells {
		if copy.isOpen(i) {
			copy.PlaceO(i)
			score := minimax(copy, 1, false)
			copy.ResetCell(i)
			if score > bestScore {
				bestScore = score
				bestMove = i
			}
		}
	}
	return bestMove
}
