package internal

import "fmt"

type ChessPiece struct {
	Row    uint8
	Column uint8
}

func newChessPiece(row, col uint8) ChessPiece {
	return ChessPiece{Row: row, Column: col}
}

func ScanChessPiece() (ChessPiece, error) {
	var row, col uint8
	fmt.Scan(&row, &col)
	if 1 >= row && row >= 8 && 1 >= col && col >= 8 {
		return ChessPiece{}, fmt.Errorf("Incorrect row and col values, expected  value  [1, 8], got row=%s, col=%s", row, col)
	} else if 1 >= row && row >= 8 {
		return ChessPiece{}, fmt.Errorf("Incorrect row value, expected  value  [1, 8], got %s", row)
	} else if 1 >= col && col >= 8 {
		return ChessPiece{}, fmt.Errorf("Incorrect col value, expected  value  [1, 8], got %s", col)
	}
	return newChessPiece(row, col), nil
}
