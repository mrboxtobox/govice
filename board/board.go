package board

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

type Rank uint8

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

type File uint8

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
type Square uint8

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
