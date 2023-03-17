package main

import (
	"fmt"
	"tbolimpiada_semifinals_DevOps/internal"
)

const (
	KING          = 0
	BISHOP        = 1
	ROOK          = 2
	FIGURE_NUMBER = 3
)

func main() {
	figures := make([]internal.ChessPiece, FIGURE_NUMBER)
	for i := 0; i < FIGURE_NUMBER; i++ {
		cp, err := internal.ScanChessPiece()
		if err != nil {
			panic(err)
		}
		figures[i] = cp
	}
	king := figures[KING]
	bishop := figures[BISHOP]
	rook := figures[ROOK]

	fmt.Println(internal.IsThereAThreat(king, bishop, rook))
}
