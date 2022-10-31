package board

const (
	BoardSquareCount = 120
	// Maximum number of half-moves.
	MaxGameMoves = 2048
)

type Board struct {
	pieces [BoardSquareCount]uint8
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

var Square120ToSquare64 [BoardSquareCount]int8
var Square64ToSquare120 [64]int8

func FileRankTo120Square(file File, rank Rank) int8 {
	return (21 + int8(file)) + int8(rank)*10
}

func Init() {
	for index := 0; index < BoardSquareCount; index++ {
		Square120ToSquare64[index] = 65
	}

	for index := 0; index < 64; index++ {
		Square64ToSquare120[index] = 120
	}

	var sq64 int8
	for rank := Rank1; rank <= Rank8; rank++ {
		for file := FileA; file <= FileH; file++ {
			sq := FileRankTo120Square(file, rank)
			Square64ToSquare120[sq64] = sq
			Square120ToSquare64[sq] = sq64
			sq64++
		}
	}
}

func SQ64(sq120 int8) int8 {
	return Square120ToSquare64[sq120]
}
