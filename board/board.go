package board

import (
	"fmt"
)

func ShowSqAtSide(side Color, pos *Board) {
	for rank := Rank8; rank >= Rank1; rank-- {
		fmt.Printf("%d  ", rank+1)
		for file := FileA; file <= FileH; file++ {
			sq := FileRankTo120Square(file, rank)
			if SqAttacked(*pos, Square(sq), side) {
				fmt.Printf("X")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func PrintBoard(pos Board) {
	fmt.Printf("Game Board:\n\n")
	for rank := Rank8; rank >= Rank1; rank-- {
		fmt.Printf("%d  ", rank+1)
		for file := FileA; file <= FileH; file++ {
			sq := FileRankTo120Square(file, rank)
			piece := pos.pieces[sq]
			fmt.Printf("%3c", PieceChar[piece])
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n   ")
	for file := FileA; file <= FileH; file++ {
		fmt.Printf("%3c", 'a'+file)
	}
	fmt.Print("\n\n")

	fmt.Printf("Side: %v\n", pos.Side)
	fmt.Printf("En Passant:%d\n", pos.enPas)
	var castle string
	if pos.castlePerm&uint8(WKCA) > 0 {
		castle = castle + "K"
	} else {
		castle = castle + "-"
	}
	if pos.castlePerm&uint8(WQCA) > 0 {
		castle = castle + "Q"
	} else {
		castle = castle + "-"
	}
	if pos.castlePerm&uint8(BKCA) > 0 {
		castle = castle + "k"
	} else {
		castle = castle + "-"
	}
	if pos.castlePerm&uint8(BQCA) > 0 {
		castle = castle + "q"
	} else {
		castle = castle + "-"
	}
	fmt.Printf("Castle: %s\n", castle)

	fmt.Printf("PositionKey: %x\n\n", pos.positionKey)

}

func CheckBoard(pos *Board) {
	if true {
		return
	}

	var tmpPieceCount [13]int
	var tmpBigPiece [2]uint8
	var tmpMajorPiece [2]uint8
	var tmpMinorPiece [2]uint8
	var tmpMaterial [2]int

	var tmpPawns [3]uint64
	tmpPawns[WHITE] = pos.Pawns[WHITE]
	tmpPawns[BLACK] = pos.Pawns[BLACK]
	tmpPawns[Both] = pos.Pawns[Both]

	// Check piece list.
	for piece := WhitePawn; piece <= BlackKing; piece++ {
		for pieceNum := 0; pieceNum < int(pos.pceNum[piece]); pieceNum++ {
			sq120 := pos.pieceList[piece][pieceNum]
			assert(pos.pieces[sq120] == piece)
		}
	}

	// Piece count and counters
	for sq64 := int8(0); sq64 < 64; sq64++ {
		sq120 := SQ120(sq64)
		piece := pos.pieces[sq120]
		tmpPieceCount[piece]++
		color := PieceColor[piece]
		if BigPiece[piece] {
			tmpBigPiece[color]++
		}
		if MinorPiece[piece] {
			tmpMinorPiece[color]++
		}
		if MajorPiece[piece] {
			tmpMajorPiece[color]++
		}

		// fmt.Println(color, piece, sq64)
		if color != Both {
			tmpMaterial[color] += PieceValue[piece]
		}
	}

	for piece := uint8(WhitePawn); piece <= uint8(BlackKing); piece++ {
		// fmt.Println(tmpPieceCount[piece], int(pos.pieceCounts[piece]))
		assert(tmpPieceCount[piece] == int(pos.pceNum[piece]))
	}

	// Check bit boards.
	pcount := CountBits(tmpPawns[WHITE])
	assert(pcount == pos.pceNum[WhitePawn])
	pcount = CountBits(tmpPawns[BLACK])
	assert(pcount == pos.pceNum[BlackPawn])
	pcount = CountBits(tmpPawns[Both])
	assert(pcount == pos.pceNum[WhitePawn]+pos.pceNum[BlackPawn])

	// Check bitboard squares.
	for tmpPawns[WHITE] != 0 {
		sq64 := PopBit(&tmpPawns[WHITE])
		assert(pos.pieces[SQ120(sq64)] == WhitePawn)

	}

	for tmpPawns[BLACK] != 0 {
		sq64 := PopBit(&tmpPawns[BLACK])
		assert(pos.pieces[SQ120(sq64)] == BlackPawn)
	}

	for tmpPawns[Both] != 0 {
		sq64 := PopBit(&tmpPawns[Both])
		assert(pos.pieces[SQ120(sq64)] == WhitePawn || pos.pieces[SQ120(sq64)] == BlackPawn)
	}

	assert(tmpMaterial[WHITE] == pos.material[WHITE] && tmpMaterial[BLACK] == pos.material[BLACK])
	assert(tmpMinorPiece[WHITE] == pos.minorPieceCounts[WHITE] && tmpMinorPiece[BLACK] == pos.minorPieceCounts[BLACK])
	assert(tmpMajorPiece[WHITE] == pos.majorPieceCounts[WHITE] && tmpMajorPiece[BLACK] == pos.majorPieceCounts[BLACK])
	assert(tmpBigPiece[WHITE] == pos.bigPieceCounts[WHITE] && tmpBigPiece[BLACK] == pos.bigPieceCounts[BLACK])

	assert(pos.Side == WHITE || pos.Side == BLACK)
	assert(GeneratePositionKey(*pos) == pos.positionKey)

	assert(pos.enPas == NO_SQ ||
		(Rank(RanksBoard[pos.enPas]) == Rank6 && pos.Side == WHITE) ||
		(Rank(RanksBoard[pos.enPas]) == Rank3 && pos.Side == BLACK))

	// fmt.Println(pos.pieces, pos.pieces[pos.kings[White]])
	assert(pos.pieces[pos.KingSq[WHITE]] == WhiteKing)
	assert(pos.pieces[pos.KingSq[BLACK]] == BlackKing)

}

func SQOFFBOARD(sq int) bool {
	return FilesBrd[sq] == int(OFFBOARD)
}

func ResetBoard(board *Board) {
	var index int8
	for index = 0; index < BoardSquareCount; index++ {
		board.pieces[index] = OFFBOARD
	}

	for index = 0; index < 64; index++ {
		board.pieces[SQ120(index)] = EMPTY
	}

	for index = 0; index < 2; index++ {
		board.bigPieceCounts[index] = 0
		board.majorPieceCounts[index] = 0
		board.minorPieceCounts[index] = 0
		board.material[index] = 0
	}

	// Pawns indexed to 3 for bitboards.
	for index := 0; index < 3; index++ {
		board.Pawns[index] = 0
	}

	for index = 0; index < 13; index++ {
		board.pceNum[index] = 0
	}

	board.KingSq[WHITE] = NO_SQ
	board.KingSq[BLACK] = NO_SQ

	board.Side = Both
	board.enPas = NO_SQ
	board.fiftyMove = 0

	board.ply = 0
	board.hisPly = 0
	board.castlePerm = 0
	board.positionKey = 0
}

const (
	BoardSquareCount = 120
	// Maximum number of half-moves.
	MaxGameMoves = 2048

	MaxPositionMoves = 256
)

type Board struct {
	pieces [BoardSquareCount]Piece
	// Bit representation of pawn positions.
	// One for White, Black and Both.
	Pawns [3]uint64

	KingSq [2]Square
	enPas  Square
	Side   Color

	// When this hits 100, the game is drawn (fifty-move rule).
	fiftyMove uint8

	// Number of half-moves in current search.
	ply uint16

	// Number of half-moves in total game so far.
	hisPly int

	// Castling permission encoded in 4 bits.
	castlePerm uint8

	// Unique key generated for each position.
	positionKey uint64

	pceNum [13]int8
	// Big piece is everything that's not a pawn.
	bigPieceCounts   [2]uint8
	majorPieceCounts [2]uint8
	minorPieceCounts [2]uint8
	// Material scores for Black and White
	material [2]int

	// Indexed by hisPly and used to undo to the last move.
	history [MaxGameMoves]UndoInfo

	// Piece List (as Map).
	pieceList [13][10]int
}

type UndoInfo struct {
	move        int
	castlePerm  uint8
	enPas       Square
	fiftyMoves  uint8
	positionKey uint64
}

type CastlingRights uint8

const (
	// Use 4 bits to represent castling permission.
	WKCA CastlingRights = 1 << iota
	WQCA
	BKCA
	BQCA
)

type Piece uint8

const (
	EMPTY Piece = iota
	WhitePawn
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing

	OFFBOARD // TODO: This was under the square.
)

// Use signed bits since Go doesn't error out on int overflow.
// When checking in for loops, we could wrap around to a larger positive number.
type Rank int8

const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	NoRank
)

type File int8

const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
	NoFile
)

type Color uint8

const (
	WHITE Color = iota
	BLACK
	Both
)

// Using VICE's Board Representation. See github.com/peterwankman/vice.
type Square int8

const (
	A1 Square = iota + 21
	B1
	C1
	D1
	E1
	F1
	G1
	H1
)

const (
	A2 Square = iota + 31
	B2
	C2
	D2
	E2
	F2
	G2
	H2
)

const (
	A3 Square = iota + 41
	B3
	C3
	D3
	E3
	F3
	G3
	H3
)

const (
	A4 Square = iota + 51
	B4
	C4
	D4
	E4
	F4
	G4
	H4
)

const (
	A5 Square = iota + 61
	B5
	C5
	D5
	E5
	F5
	G5
	H5
)

const (
	A6 Square = iota + 71
	B6
	C6
	D6
	E6
	F6
	G6
	H6
)

const (
	A7 Square = iota + 81
	B7
	C7
	D7
	E7
	F7
	G7
	H7
)

const (
	A8 Square = iota + 91
	B8
	C8
	D8
	E8
	F8
	G8
	H8

	NO_SQ // Square is off the board.
)

func FileRankTo120Square(file File, rank Rank) int8 {
	return (21 + int8(file)) + int8(rank)*10
}

func SQ64(sq120 int8) int8 {
	return Square120ToSquare64[sq120]
}

func SQ120(sq64 int8) int8 {
	return Square64ToSquare120[sq64]
}

func assert(condition bool) {
	if !condition {
		panic(condition)
	}
}
