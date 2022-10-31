package board

import "fmt"

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

	fmt.Printf("Side: %v\n", pos.side)
	fmt.Printf("En Passant:%d\n", pos.enPassant)
	var castle string
	if pos.castlePermissions&uint8(WhiteKingCastle) > 0 {
		castle = castle + "K"
	} else {
		castle = castle + "-"
	}
	if pos.castlePermissions&uint8(WhiteQueenCastle) > 0 {
		castle = castle + "Q"
	} else {
		castle = castle + "-"
	}
	if pos.castlePermissions&uint8(BlackKingCastle) > 0 {
		castle = castle + "k"
	} else {
		castle = castle + "-"
	}
	if pos.castlePermissions&uint8(BlackQueenCastle) > 0 {
		castle = castle + "q"
	} else {
		castle = castle + "-"
	}
	fmt.Printf("Castle: %s\n", castle)

	fmt.Printf("KePositionKey: %x\n\n", pos.positionKey)

}

func ResetBoard(board *Board) {
	var index int8
	for index = 0; index < BoardSquareCount; index++ {
		board.pieces[index] = OffBoard
	}

	for index = 0; index < 64; index++ {
		board.pieces[SQ120(index)] = NoPiece
	}

	for index = 0; index < 3; index++ {
		board.bigPieceCounts[index] = 0
		board.majorPieceCounts[index] = 0
		board.minorPieceCounts[index] = 0
		board.pawns[index] = 0
	}

	for index = 0; index < 13; index++ {
		board.pieceCounts[index] = 0
	}

	board.kings[White] = NoSquare
	board.kings[Black] = NoSquare

	board.side = Both
	board.enPassant = NoSquare
	board.fiftyMoveCount = 0

	board.plyCount = 0
	board.historyPlyCount = 0
	board.castlePermissions = 0
	board.positionKey = 0
}

const (
	BoardSquareCount = 120
	// Maximum number of half-moves.
	MaxGameMoves = 2048
)

type Board struct {
	pieces [BoardSquareCount]Piece
	// Bit representation of pawn positions.
	// One for White, Black and Both.
	pawns [3]uint64

	kings     [2]Square
	enPassant Square
	side      Color

	// When this hits 100, the game is drawn (fifty-move rule).
	fiftyMoveCount uint8

	// Number of half-moves in current search.
	plyCount uint16

	// Number of half-moves in total game so far.
	historyPlyCount uint16

	// Castling permission encoded in 4 bits.
	castlePermissions uint8

	// Unique key generated for each position.
	positionKey uint64

	pieceCounts [13]uint8
	// Big piece is everything that's not a pawn.
	bigPieceCounts   [3]uint8
	majorPieceCounts [3]uint8
	minorPieceCounts [3]uint8

	// Indexed by hisPly and used to undo to the last move.
	undoInfo [MaxGameMoves]UndoInfo

	// Piece List (as Map).
	pieceMap map[Piece][]Square
}

type UndoInfo struct {
	move              int
	castlePermissions uint8
	enPassant         Square
	fiftyMoveCount    uint8
	positionKey       uint64
}

type CastlingRights uint8

const (
	// Use 4 bits to represent castling permission.
	WhiteKingCastle CastlingRights = 1 << iota
	WhiteQueenCastle
	BlackKingCastle
	BlackQueenCastle
)

type Piece uint8

const (
	NoPiece Piece = iota
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

	OffBoard // TODO: This was under the square.
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
	White Color = iota
	Black
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

	NoSquare // Square is off the board.
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
