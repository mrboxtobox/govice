package board

import (
	"math"
)

const PawnIsolated = -10

var PawnPassed = [8]int{0, 5, 10, 20, 35, 60, 100, 200}

const RookOpenFile = 10
const RookSemiOpenFile = 5
const QueenOpenFile = 5
const QueenSemiOpenFile = 3
const BishopPair = 30

var PawnTable = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	10, 10, 0, -10, -10, 0, 10, 10,
	5, 0, 0, 5, 5, 0, 0, 5,
	0, 0, 10, 20, 20, 10, 0, 0,
	5, 5, 5, 10, 10, 5, 5, 5,
	10, 10, 10, 20, 20, 10, 10, 10,
	20, 20, 20, 30, 30, 20, 20, 20,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var KnightTable = [64]int{
	0, -10, 0, 0, 0, 0, -10, 0,
	0, 0, 0, 5, 5, 0, 0, 0,
	0, 0, 10, 10, 10, 10, 0, 0,
	0, 0, 10, 20, 20, 10, 5, 0,
	5, 10, 15, 20, 20, 15, 10, 5,
	5, 10, 10, 20, 20, 10, 10, 5,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var BishopTable = [64]int{
	0, 0, -10, 0, 0, -10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 0, 10, 15, 15, 10, 0, 0,
	0, 0, 0, 10, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var RookTable = [64]int{
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	0, 0, 5, 10, 10, 5, 0, 0,
	25, 25, 25, 25, 25, 25, 25, 25,
	0, 0, 5, 10, 10, 5, 0, 0,
}

// Encourage castling.
var KingE = [64]int{
	-50, -10, 0, 0, 0, 0, -10, -50,
	-10, 0, 10, 10, 10, 10, 0, -10,
	0, 10, 15, 15, 15, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 20, 20, 15, 10, 0,
	0, 10, 15, 15, 15, 15, 10, 0,
	-10, 0, 10, 10, 10, 10, 0, -10,
	-50, -10, 0, 0, 0, 0, -10, -50,
}

var KingO = [64]int{
	0, 5, 5, -10, -10, 0, 10, 5,
	-30, -30, -30, -30, -30, -30, -30, -30,
	-50, -50, -50, -50, -50, -50, -50, -50,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
	-70, -70, -70, -70, -70, -70, -70, -70,
}

var MIRROR64 = [64]int{
	56, 57, 58, 59, 60, 61, 62, 63,
	48, 49, 50, 51, 52, 53, 54, 55,
	40, 41, 42, 43, 44, 45, 46, 47,
	32, 33, 34, 35, 36, 37, 38, 39,
	24, 25, 26, 27, 28, 29, 30, 31,
	16, 17, 18, 19, 20, 21, 22, 23,
	8, 9, 10, 11, 12, 13, 14, 15,
	0, 1, 2, 3, 4, 5, 6, 7,
}

// Evaluate position from White POV == mirror position from Black POV.
func MirrorBoard(pos *Board) {

	tempPiecesArray := [64]Piece{}
	tempSide := pos.Side ^ 1
	SwapPiece := [13]Piece{EMPTY, BlackPawn, BlackKnight, BlackBishop, BlackRook, BlackQueen, BlackKing, WhitePawn, WhiteKnight, WhiteBishop, WhiteRook, WhiteQueen, WhiteKing}
	var tempCastlePerm CastlingRights
	tempEnPas := NO_SQ

	if (pos.castlePerm & uint8(WKCA)) != 0 {
		tempCastlePerm |= BKCA
	}
	if (pos.castlePerm & uint8(WQCA)) != 0 {
		tempCastlePerm |= BQCA
	}

	if (pos.castlePerm & uint8(BKCA)) != 0 {
		tempCastlePerm |= WKCA
	}
	if (pos.castlePerm & uint8(BQCA)) != 0 {
		tempCastlePerm |= WQCA
	}

	if pos.enPas != NO_SQ {
		v := int8(pos.enPas)
		tempEnPas = Square(SQ120(int8(MIRROR64[int(SQ64(v))])))
	}

	for sq := 0; sq < 64; sq++ {
		tempPiecesArray[sq] = pos.pieces[SQ120(int8(MIRROR64[int8(sq)]))]
	}

	ResetBoard(pos)

	for sq := 0; sq < 64; sq++ {
		tp := SwapPiece[tempPiecesArray[sq]]
		pos.pieces[SQ120(int8(sq))] = tp
	}

	pos.Side = tempSide
	pos.castlePerm = uint8(tempCastlePerm)
	pos.enPas = tempEnPas

	pos.posKey = GeneratePositionKey(*pos)

	UpdateListsMaterial(pos)

	CheckBoard(pos)
}

// TODO: Need to confirm this.
func MaterialDraw(pos *Board) bool {
	if pos.pceNum[WhiteRook] == 0 && pos.pceNum[BlackRook] == 0 && pos.pceNum[WhiteQueen] == 0 && pos.pceNum[BlackQueen] == 0 {
		if pos.pceNum[BlackBishop] == 0 && pos.pceNum[WhiteBishop] == 0 {
			if pos.pceNum[WhiteKnight] < 3 && pos.pceNum[BlackKnight] < 3 {
				return true
			}
		} else if pos.pceNum[WhiteKnight] == 0 && pos.pceNum[BlackKnight] == 0 {
			if math.Abs(float64(pos.pceNum[WhiteBishop]-pos.pceNum[BlackBishop])) < 2 {
				return true
			}
		} else if (pos.pceNum[WhiteKnight] < 3 && pos.pceNum[WhiteBishop] == 0) || (pos.pceNum[WhiteBishop] == 1 && pos.pceNum[WhiteKnight] == 0) {
			if (pos.pceNum[BlackKnight] < 3 && pos.pceNum[BlackBishop] == 0) || (pos.pceNum[BlackBishop] == 1 && pos.pceNum[BlackKnight] == 0) {
				return true
			}
		}
	} else if pos.pceNum[WhiteQueen] == 0 && pos.pceNum[BlackQueen] == 0 {
		if pos.pceNum[WhiteRook] == 1 && pos.pceNum[BlackRook] == 1 {
			if (pos.pceNum[WhiteKnight]+pos.pceNum[WhiteBishop]) < 2 && (pos.pceNum[BlackKnight]+pos.pceNum[BlackBishop]) < 2 {
				return true
			}
		} else if pos.pceNum[WhiteRook] == 1 && pos.pceNum[BlackRook] == 0 {
			if (pos.pceNum[WhiteKnight]+pos.pceNum[WhiteBishop] == 0) && (((pos.pceNum[BlackKnight] + pos.pceNum[BlackBishop]) == 1) || ((pos.pceNum[BlackKnight] + pos.pceNum[BlackBishop]) == 2)) {
				return true
			}
		} else if pos.pceNum[BlackRook] == 1 && pos.pceNum[WhiteRook] == 0 {
			if (pos.pceNum[BlackKnight]+pos.pceNum[BlackBishop] == 0) && (((pos.pceNum[WhiteKnight] + pos.pceNum[WhiteBishop]) == 1) || ((pos.pceNum[WhiteKnight] + pos.pceNum[WhiteBishop]) == 2)) {
				return true
			}
		}
	}
	return false
}

// When we're in an endgame situation.
func ENDGAME_MAT() int {
	return PieceValue[WhiteRook] + 2*PieceValue[WhiteKnight] + 2*PieceValue[WhitePawn]
}

func EvalPosition(pos *Board) int {
	score := pos.material[WHITE] - pos.material[BLACK]

	// fmt.Printf("%v\n", pos.pceNum)
	if pos.pceNum[WhitePawn] == 0 && pos.pceNum[BlackPawn] == 0 && MaterialDraw(pos) == true {
		// fmt.Println("Returning material draw")
		return 0
	}

	// For each piece, add individual material value.
	pce := WhitePawn
	for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
		sq := pos.pieceList[pce][pceNum]
		assert(SqOnBoard(sq))
		score += PawnTable[SQ64(int8(sq))]

		// If isolated pawn, add the penalty (-10).
		if (IsolatedMask[SQ64(int8(sq))] & pos.Pawns[WHITE]) == 0 {
			// fmt.Printf("White Pawn Isolated: %s\n", PrSq(sq))
			score += PawnIsolated
		}

		// If passed pawn for black, add the penalty.
		// Zero means no pawns ahead of the black pawns.
		if ((WhitePassedMask[SQ64(int8(sq))]) & pos.Pawns[BLACK]) == 0 {
			// fmt.Printf("White Pawn Passed: %s\n", PrSq(sq))
			score += PawnPassed[RanksBoard[sq]]
		}
	}

	pce = BlackPawn
	for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
		sq := pos.pieceList[pce][pceNum]
		assert(SqOnBoard(sq))
		score -= PawnTable[MIRROR64[SQ64(int8(sq))]]

		if (IsolatedMask[SQ64(int8(sq))] & pos.Pawns[BLACK]) == 0 {
			// fmt.Printf("Black Pawn Isolated: %s\n", PrSq(sq))
			score -= PawnIsolated
		}

		if (BlackPassedMask[SQ64(int8(sq))] & pos.Pawns[WHITE]) == 0 {
			// Black is going in the opposite direction.
			// fmt.Printf("Black Pawn Passed: %s\n", PrSq(sq))
			score -= PawnPassed[7-RanksBoard[sq]]
		}
	}

	pce = WhiteKnight
	for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
		sq := pos.pieceList[pce][pceNum]
		assert(SqOnBoard(sq))
		score += KnightTable[SQ64(int8(sq))]
	}

	pce = BlackKnight
	for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
		sq := pos.pieceList[pce][pceNum]
		assert(SqOnBoard(sq))
		score -= KnightTable[MIRROR64[SQ64(int8(sq))]]
	}

	pce = WhiteBishop
	for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
		sq := pos.pieceList[pce][pceNum]
		assert(SqOnBoard(sq))
		score += BishopTable[SQ64(int8(sq))]
	}

	pce = BlackBishop
	for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
		sq := pos.pieceList[pce][pceNum]
		assert(SqOnBoard(sq))
		score -= BishopTable[MIRROR64[SQ64(int8(sq))]]
	}

	pce = WhiteRook
	for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
		sq := pos.pieceList[pce][pceNum]
		assert(SqOnBoard(sq))
		score += RookTable[SQ64(int8(sq))]

		// Open file == 0
		if (pos.Pawns[BOTH] & FileBBMask[FilesBrd[sq]]) == 0 {
			score += RookOpenFile
		} else if (pos.Pawns[WHITE] & FileBBMask[FilesBrd[sq]]) == 0 {
			// When we don't have our own pawns, it's semi-open
			score += RookSemiOpenFile
		}
	}

	pce = BlackRook
	for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
		sq := pos.pieceList[pce][pceNum]
		assert(SqOnBoard(sq))
		score -= RookTable[MIRROR64[SQ64(int8(sq))]]

		if (pos.Pawns[BOTH] & FileBBMask[FilesBrd[sq]]) == 0 {
			score -= RookOpenFile
		} else if (pos.Pawns[BLACK] & FileBBMask[FilesBrd[sq]]) == 0 {
			score -= RookSemiOpenFile
		}
	}

	pce = WhiteQueen
	for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
		sq := pos.pieceList[pce][pceNum]
		assert(SqOnBoard(sq))

		if (pos.Pawns[BOTH] & FileBBMask[FilesBrd[sq]]) == 0 {
			score += QueenOpenFile
		} else if (pos.Pawns[WHITE] & FileBBMask[FilesBrd[sq]]) == 0 {
			score += QueenSemiOpenFile
		}
	}

	pce = BlackQueen
	for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
		sq := pos.pieceList[pce][pceNum]
		assert(SqOnBoard(sq))

		if (pos.Pawns[BOTH] & FileBBMask[FilesBrd[sq]]) == 0 {
			score -= QueenOpenFile
		} else if (pos.Pawns[BLACK] & FileBBMask[FilesBrd[sq]]) == 0 {
			score -= QueenSemiOpenFile
		}
	}

	pce = WhiteKing
	sq := pos.pieceList[pce][0]
	if pos.material[BLACK] <= ENDGAME_MAT() {
		score += KingE[SQ64(int8(sq))]
	} else {
		score += KingO[SQ64(int8(sq))]
	}

	pce = BlackKing
	sq = pos.pieceList[pce][0]
	if pos.material[WHITE] <= ENDGAME_MAT() {
		score -= KingE[MIRROR64[SQ64(int8(sq))]]
	} else {
		score -= KingO[MIRROR64[SQ64(int8(sq))]]
	}

	if pos.pceNum[WhiteBishop] >= 2 {
		score += BishopPair
	}
	if pos.pceNum[BlackBishop] >= 2 {
		score -= BishopPair
	}

	// fmt.Println(score)
	if pos.Side == WHITE {
		return score
	} else {
		return -score
	}
}
