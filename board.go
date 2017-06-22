package main

import (
	"fmt"
	"strconv"
)

// Status definitions
const (
	StatusInProgress = "I"
	StatusXWin       = "X"
	StatusOWin       = "O"
	StatusDraw       = "D"
)

// Board definitions
const (
	BoardLength = 3
)

// Board store the tic-tac-toe board info
type Board struct {
	grids [BoardLength][BoardLength]string
}

func (b *Board) newBoard() {
	for y := 0; y < BoardLength; y++ {
		for x := 0; x < BoardLength; x++ {
			b.grids[x][y] = strconv.Itoa(y*BoardLength + x + 1)
		}
	}
}

// Place a mark on the board, at particular position
func (b *Board) Place(player string, posX int, posY int) bool {
	if b.Status() != StatusInProgress || b.grids[posX][posY] == "X" || b.grids[posX][posY] == "O" {
		return false
	}

	b.grids[posX][posY] = player
	return true
}

// Status return the current state of the board
func (b *Board) Status() string {
	// Check if player win
	if b.isWin("X") {
		return StatusXWin
	}

	if b.isWin("O") {
		return StatusOWin
	}

	// Check if still empty grids
	for y := 0; y < BoardLength; y++ {
		for x := 0; x < BoardLength; x++ {
			if b.grids[x][y] != "X" && b.grids[x][y] != "O" {
				return StatusInProgress
			}
		}
	}

	// Otherwise it's a draw
	return StatusDraw
}

func (b *Board) isWin(player string) bool {
	for y := 0; y < BoardLength; y++ {
		isHorizontalWin := true
		for x := 0; x < BoardLength; x++ {
			if b.grids[x][y] != player {
				isHorizontalWin = false
				break
			}
		}

		if isHorizontalWin {
			return true
		}
	}

	for x := 0; x < BoardLength; x++ {
		isVerticalWin := true
		for y := 0; y < BoardLength; y++ {
			if b.grids[x][y] != player {
				isVerticalWin = false
				break
			}
		}

		if isVerticalWin {
			return true
		}
	}

	isDiagonalWin := true
	isReverseDiagonalWin := true
	for i := 0; i < BoardLength; i++ {
		if b.grids[i][i] != player {
			isDiagonalWin = false
		}
		if b.grids[BoardLength-i-1][i] != player {
			isReverseDiagonalWin = false
		}
	}

	if isDiagonalWin || isReverseDiagonalWin {
		return true
	}

	return false
}

// Display redner the board
func (b *Board) Display() {
	for y := 0; y < BoardLength; y++ {
		for x := 0; x < BoardLength; x++ {
			fmt.Print(b.grids[x][y])
		}
		fmt.Println()
	}
}
