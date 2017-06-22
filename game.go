package main

import (
	"fmt"
)

// Game is responsible for handling the game related logics
type Game struct {
	board           Board
	currentPlayer   string
	currentNumRound int
}

// NewGame initialize a tic-tac-toe game
func (g *Game) NewGame() {
	g.board.newBoard()
	g.currentPlayer = "O"
	g.currentNumRound = 0
}

// SwitchPlayer change turn for player X and O
func (g *Game) SwitchPlayer() {
	switch g.currentPlayer {
	case "X":
		g.currentPlayer = "O"
	case "O":
		g.currentPlayer = "X"
	}

	g.currentNumRound++
}

// Run the game
func (g *Game) Run() {
	g.NewGame()
	for g.board.Status() == StatusInProgress {
		g.SwitchPlayer()
		fmt.Println()
		fmt.Printf("Round %d, player %q:\n", g.currentNumRound, g.currentPlayer)

		for isValid := false; !isValid; {
			g.board.Display()

			var posCode int
			fmt.Scan(&posCode)
			posX := (posCode - 1) % 3
			posY := (posCode - 1) / 3

			isValid = g.board.Place(g.currentPlayer, posX, posY)
			if !isValid {
				fmt.Println("Not valid, try again")
			}
		}

		if g.currentNumRound >= 9 {
			break
		}
	}

	fmt.Println()
	fmt.Println("Final result")
	fmt.Println("========================")

	switch g.board.Status() {
	case StatusDraw:
		fmt.Println("It is a draw!")
	case StatusXWin:
		fmt.Println("X Win!")
	case StatusOWin:
		fmt.Println("O Win!")
	}

	g.board.Display()
}

func main() {
	fmt.Println("github.com/gaimuk/go-ttt")
	fmt.Println("========================")

	g := Game{}
	g.NewGame()
	g.Run()
}
