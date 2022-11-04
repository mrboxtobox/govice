package board

import "fmt"

const NOMOVE = 0

func HASH_PCE(pos *Board, pce Piece, sq int) {
	pos.posKey ^= PieceKeys[pce][sq]
}

func HASH_CA(pos *Board) {
	pos.posKey ^= CastleKeys[pos.castlePerm]
}

func HASH_SIDE(pos *Board) {
	pos.posKey ^= SideKey
}

func HASH_EP(pos *Board) {
	pos.posKey ^= PieceKeys[EMPTY][pos.enPas]
}

// Do bitwise and with this array and the square from and to
var CastlePerm = [120]uint8{
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 13, 15, 15, 15, 12, 15, 15, 14, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 7, 15, 15, 15, 3, 15, 15, 11, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
}

func ClearPiece(sq int, pos *Board) {
	pce := pos.pieces[sq]
	col := PieceColor[pce]
	t_pceNum := -1

	assert(SqOnBoard(sq))
	assert(PieceValid(int(pce)))

	HASH_PCE(pos, pce, sq)

	pos.pieces[sq] = EMPTY
	pos.material[col] -= PieceValue[pce]

	if BigPiece[pce] {
		pos.bigPieceCounts[col]--

		if MajorPiece[pce] {
			pos.majorPieceCounts[col]--
		} else {
			pos.minorPieceCounts[col]--
		}
	} else {
		ClearBit(&pos.Pawns[col], int(SQ64(int8(sq))))
		ClearBit(&pos.Pawns[Both], int(SQ64(int8(sq))))
	}

	// Swap with that position for the square.
	for index := 0; index < int(int(pos.pceNum[pce])); index++ {
		if pos.pieceList[pce][index] == sq {
			t_pceNum = index
			break
		}
	}
	assert(t_pceNum != -1)
	pos.pceNum[pce]--
	pos.pieceList[pce][t_pceNum] = pos.pieceList[pce][int(pos.pceNum[pce])]
}

func AddPiece(sq int, pos *Board, pce Piece) {
	var col = PieceColor[pce]

	assert(PieceValid(int(pce)))
	assert(SqOnBoard(int(sq)))

	HASH_PCE(pos, pce, int(sq))

	pos.pieces[sq] = pce

	if BigPiece[pce] {
		pos.bigPieceCounts[col]++
		if MajorPiece[pce] {
			pos.majorPieceCounts[col]++
		} else {
			pos.minorPieceCounts[col]++
		}
	} else {
		SetBit(&pos.Pawns[col], int(SQ64(int8(sq))))
		SetBit(&pos.Pawns[Both], int(SQ64(int8(sq))))
	}

	pos.material[col] += PieceValue[pce]
	pos.pieceList[pce][int(pos.pceNum[pce])] = int(sq)

	pos.pceNum[pce]++

}

func MovePiece(from, to int, pos *Board) {
	index := 0
	pce := pos.pieces[from]
	col := PieceColor[pce]
	// Could only want in debug mode.
	t_PieceNum := false

	assert(SqOnBoard(from))
	assert(SqOnBoard(to))

	HASH_PCE(pos, pce, from)
	pos.pieces[from] = EMPTY

	HASH_PCE(pos, pce, to)
	pos.pieces[to] = pce

	if !BigPiece[pce] {
		ClearBit(&pos.Pawns[col], int(SQ64(int8(from))))
		ClearBit(&pos.Pawns[Both], int(SQ64(int8(from))))
		SetBit(&pos.Pawns[col], int(SQ64(int8(to))))
		SetBit(&pos.Pawns[Both], int(SQ64(int8(to))))
	}

	for index = 0; index < int(pos.pceNum[pce]); index++ {
		if pos.pieceList[pce][index] == from {
			pos.pieceList[pce][index] = to
			t_PieceNum = true
			break
		}
	}
	assert(t_PieceNum)
}

func MakeMove(pos *Board, move int) bool {
	from := FromSQ(move)
	to := ToSQ(move)
	side := pos.Side
	captured := Captured(move)
	prPce := Promoted(move)

	CheckBoard(pos)
	assert(SqOnBoard(from))
	assert(SqOnBoard(to))
	assert(SideValid(side))
	assert(PieceValid(int(pos.pieces[from])))

	pos.history[pos.hisPly].positionKey = pos.posKey

	if (move & MFLAGEP) != 0 {
		if side == WHITE {
			ClearPiece(to-10, pos)
		} else {
			ClearPiece(to+10, pos)
		}
	} else if (move & MFLAGCA) != 0 {
		switch Square(to) {
		case C1:
			MovePiece(int(A1), int(D1), pos)
		case C8:
			MovePiece(int(A8), int(D8), pos)
		case G1:
			MovePiece(int(H1), int(F1), pos)
		case G8:
			MovePiece(int(H8), int(F8), pos)
		default:
			PrintBoard(*pos)
			println("Move: ", PrMove(move))
			fmt.Printf("%v %v \n", FilesBrd[to], RanksBoard[to])
			assert(false)
		}
	}

	// Hash out En Passant and Castling as they will be included by the new move.
	if pos.enPas != NO_SQ {
		HASH_EP(pos)
	}
	// Hash out state.
	HASH_CA(pos)

	pos.history[pos.hisPly].move = move
	pos.history[pos.hisPly].fiftyMoves = pos.fiftyMove
	pos.history[pos.hisPly].enPas = pos.enPas
	pos.history[pos.hisPly].castlePerm = pos.castlePerm

	pos.castlePerm &= CastlePerm[from]
	pos.castlePerm &= CastlePerm[to]
	pos.enPas = NO_SQ

	// Hash back in.
	HASH_CA(pos)

	pos.fiftyMove++

	if captured != int(EMPTY) {
		assert(PieceValid(captured))
		ClearPiece(to, pos)
		pos.fiftyMove = 0
	}

	pos.hisPly++
	pos.ply++

	if PiecePawn[pos.pieces[from]] {
		pos.fiftyMove = 0
		if (move & MFLAGPS) != 0 {
			if side == WHITE {
				pos.enPas = Square(from + 10)
				assert(RanksBoard[pos.enPas] == int(Rank3))
			} else {
				pos.enPas = Square(from - 10)
				assert(RanksBoard[pos.enPas] == int(Rank6))
			}
			// Hash in en passant after applying.
			HASH_EP(pos)
		}
	}

	MovePiece(from, to, pos)

	if prPce != int(EMPTY) {
		assert(PieceValid(prPce) && !PiecePawn[prPce])
		ClearPiece(to, pos)
		AddPiece(to, pos, Piece(prPce))
	}

	if PieceKing[pos.pieces[to]] {
		pos.KingSq[pos.Side] = Square(to)
	}

	pos.Side ^= 1
	HASH_SIDE(pos)

	CheckBoard(pos)

	// Check if king in check.
	if SqAttacked(*pos, pos.KingSq[side], pos.Side) {
		TakeMove(pos)
		return false
	}

	return true
}

// Reverse of apply move.
func TakeMove(pos *Board) {
	move := pos.history[pos.hisPly-1].move
	from := FromSQ(move)
	to := ToSQ(move)
	captured := Captured(move)

	CheckBoard(pos)

	pos.hisPly--
	pos.ply--

	assert(SqOnBoard(from))
	assert(SqOnBoard(to))

	if pos.enPas != NO_SQ {
		HASH_EP(pos)
	}
	HASH_CA(pos)

	// Hash out old state.
	pos.castlePerm = pos.history[pos.hisPly].castlePerm
	pos.fiftyMove = pos.history[pos.hisPly].fiftyMoves
	pos.enPas = pos.history[pos.hisPly].enPas

	if pos.enPas != NO_SQ {
		HASH_EP(pos)
	}
	HASH_CA(pos)

	pos.Side ^= 1
	HASH_SIDE(pos)

	if (MFLAGEP & move) != 0 {
		if pos.Side == WHITE {
			AddPiece(to-10, pos, BlackPawn)
		} else {
			AddPiece(to+10, pos, WhitePawn)
		}
	} else if (MFLAGCA & move) != 0 {
		switch Square(to) {
		case C1:
			MovePiece(int(D1), int(A1), pos)
		case C8:
			MovePiece(int(D8), int(A8), pos)
		case G1:
			MovePiece(int(F1), int(H1), pos)
		case G8:
			MovePiece(int(F8), int(H8), pos)
		default:
			assert(false)
		}
	}

	MovePiece(to, from, pos)

	if PieceKing[pos.pieces[from]] {
		pos.KingSq[pos.Side] = Square(from)
	}

	if captured != int(EMPTY) {
		assert(PieceValid(captured))
		AddPiece(to, pos, Piece(captured))
	}

	if Promoted(move) != int(EMPTY) {
		assert(PieceValid(Promoted(move)) && !PiecePawn[Promoted(move)])
		ClearPiece(from, pos)
		var pce Piece
		if PieceColor[Promoted(move)] == WHITE {
			pce = WhitePawn
		} else {
			pce = BlackPawn
		}
		AddPiece(from, pos, pce)
	}
	CheckBoard(pos)
}

func MakeNullMove(pos *Board) {
	CheckBoard(pos)
	assert(!SqAttacked(*pos, pos.KingSq[pos.Side], pos.Side^1))

	pos.ply++
	pos.history[pos.hisPly].positionKey = pos.posKey

	if pos.enPas != NO_SQ {
		HASH_EP(pos)
	}

	pos.history[pos.hisPly].move = NOMOVE
	pos.history[pos.hisPly].fiftyMoves = pos.fiftyMove
	pos.history[pos.hisPly].enPas = pos.enPas
	pos.history[pos.hisPly].castlePerm = pos.castlePerm
	pos.enPas = NO_SQ

	pos.Side ^= 1
	pos.hisPly++
	HASH_SIDE(pos)

	CheckBoard(pos)
}

func TakeNullMove(pos *Board) {
	CheckBoard(pos)

	pos.hisPly--
	pos.ply--

	pos.castlePerm = pos.history[pos.hisPly].castlePerm
	pos.fiftyMove = pos.history[pos.hisPly].fiftyMoves
	pos.enPas = pos.history[pos.hisPly].enPas

	if pos.enPas != NO_SQ {
		HASH_EP(pos)
	}
	pos.Side ^= 1
	HASH_SIDE(pos)

	CheckBoard(pos)
}
