package board

var KnDir = [8]int8{-8, -19, -21, -12, 8, 19, 21, 12}
var RkDir = [8]int8{-1, -10, 1, 10}
var BiDir = [8]int8{-9, -11, 11, 9}
var KiDir = [8]int8{-1, -10, 1, 10, -9, -11, 11, 9}

// Side doing attacking.
func SqAttacked(pos Board, sq Square, side Color) bool {
	sq2 := int8(sq)
	assert(SqOnBoard(int(sq2)))
	assert(SideValid(side))
	// CheckBoard(&pos)

	// Pawns.
	if side == WHITE {
		if pos.pieces[sq-11] == WhitePawn || pos.pieces[sq-9] == WhitePawn {
			return true
		}
	} else {
		if pos.pieces[sq+11] == BlackPawn || pos.pieces[sq+9] == BlackPawn {
			return true
		}
	}

	// Knights.
	for index := 0; index < 8; index++ {
		piece := pos.pieces[sq2+KnDir[index]]
		if piece != OFFBOARD && IsKn(piece) && PieceColor[piece] == side {
			return true
		}
	}

	// Rooks, Queens
	for index := 0; index < 4; index++ {
		dir := RkDir[index]
		tmpSq := sq2 + dir
		piece := pos.pieces[tmpSq]
		for piece != OFFBOARD {
			if piece != EMPTY {
				if IsRQ(piece) && PieceColor[piece] == side {
					return true
				}
				break
			}
			tmpSq += dir
			piece = pos.pieces[tmpSq]
		}
	}

	// Bishop, Queen
	for index := 0; index < 4; index++ {
		dir := BiDir[index]
		tmpSq := sq2 + dir
		piece := pos.pieces[tmpSq]
		for piece != OFFBOARD {
			if piece != EMPTY {
				if IsBQ(piece) && PieceColor[piece] == side {
					return true
				}
				break
			}
			tmpSq += dir
			piece = pos.pieces[tmpSq]
		}
	}

	// King.
	for index := 0; index < 8; index++ {
		piece := pos.pieces[sq2+KiDir[index]]
		if piece != OFFBOARD && IsKi(piece) && PieceColor[piece] == side {
			return true
		}
	}

	return false
}
