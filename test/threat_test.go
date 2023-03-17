package test

import (
	"tbolimpiada_semifinals_DevOps/internal"
	"testing"
)

type BishopTC struct {
	king   internal.ChessPiece
	bishop internal.ChessPiece
	want   bool
}

// bishopTC набор тестов для Слон (ходит по диагонали)
var bishopTC = []BishopTC{
	{internal.ChessPiece{1, 1}, internal.ChessPiece{2, 2}, true},
	{internal.ChessPiece{1, 1}, internal.ChessPiece{8, 8}, true},
	{internal.ChessPiece{8, 8}, internal.ChessPiece{1, 1}, true},
	{internal.ChessPiece{1, 3}, internal.ChessPiece{3, 5}, true},
	{internal.ChessPiece{4, 4}, internal.ChessPiece{8, 8}, true},
	{internal.ChessPiece{2, 4}, internal.ChessPiece{6, 3}, false},
}

func TestThreatFromBlackBishop(t *testing.T) {
	for i, tc := range bishopTC {
		got := internal.ThreatFromBlackBishop(tc.king, tc.bishop)
		if got != tc.want {
			t.Errorf("TC#%d got %t, wanted %t", i+1, got, tc.want)
		}
	}
}

type RookTC struct {
	king internal.ChessPiece
	rook internal.ChessPiece
	want bool
}

// rookTC набор тестов для Ладьи (ходит по горизонтали и вертикали)
var rookTC = []RookTC{
	{internal.ChessPiece{1, 1}, internal.ChessPiece{1, 5}, true},
	{internal.ChessPiece{6, 4}, internal.ChessPiece{4, 3}, false},
	{internal.ChessPiece{1, 1}, internal.ChessPiece{2, 5}, false},
	{internal.ChessPiece{8, 8}, internal.ChessPiece{2, 5}, false},
	{internal.ChessPiece{4, 4}, internal.ChessPiece{1, 5}, false},
}

func TestThreatFromBlackRook(t *testing.T) {
	for i, tc := range rookTC {
		got := internal.ThreatFromBlackRook(tc.king, tc.rook)
		if got != tc.want {
			t.Errorf("TC#%d got %t, wanted %t", i+1, got, tc.want)
		}
	}
}
