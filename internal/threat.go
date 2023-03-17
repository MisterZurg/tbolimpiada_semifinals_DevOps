package internal

func IsThereAThreat(king, bishop, rook ChessPiece) string {
	if ThreatFromBlackBishop(king, bishop) {
		return "Шах от слона"
	} else if ThreatFromBlackRook(king, rook) {
		return "Шах от ладьи"
	}
	return "Нет шаха"
}

// ThreatFromBlackBishop - returns whether there is a threat to the white king from the black bishop
func ThreatFromBlackBishop(king, bishop ChessPiece) bool {
	// Шахматный слон ходит по диагонали.
	// Даны две различные клетки шахматной доски, определите, может ли слон попасть с первой клетки на вторую одним ходом.
	if (king.Row-bishop.Row == king.Column-bishop.Column) || (bishop.Row-king.Row == bishop.Column-king.Column) {
		return true
	}
	return false
}

// ThreatFromBlackRook - returns whether there is a threat to the white king from the black bishop
func ThreatFromBlackRook(king, rook ChessPiece) bool {
	// Шахматная ладья ходит по горизонтали или вертикали.
	// Даны две различные клетки шахматной доски, определите, может ли ладья попасть с первой клетки на вторую одним ходом.
	if (king.Row == rook.Row && king.Column != rook.Column) || (king.Row != rook.Row && king.Column == rook.Column) {
		return true
	}
	return false
}
