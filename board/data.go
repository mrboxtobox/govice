package board

const (
	PieceChar = ".PNBRQKpnbrqk"
	SideChar  = "wb-"
	RankChar  = "12345678"
	FileChar  = "abcdefgh"
)

var PiecePawn = [13]bool{false, true, false, false, false, false, false, true, false, false, false, false, false}
var PieceKnight = [13]bool{false, false, true, false, false, false, false, false, true, false, false, false, false}
var PieceKing = [13]bool{false, false, false, false, false, false, true, false, false, false, false, false, true}
var PieceRookQueen = [13]bool{false, false, false, false, true, true, false, false, false, false, true, true, false}
var PieceBishopQueen = [13]bool{false, false, false, true, false, true, false, false, false, true, false, true, false}
var PieceSlides = [13]bool{false, false, false, true, true, true, false, false, false, true, true, true, false}

var IsolatedMask [64]uint64

// Passed Pawn--no pawn is on the direct or adjacent file
var WhitePassedMask [64]uint64
var BlackPassedMask [64]uint64
var FileBBMask [8]uint64
var RankBBMask [8]uint64

// var IsolatedMask = [64]int{}

func IsBQ(piece Piece) bool {
	return PieceBishopQueen[piece]
}

func IsRQ(piece Piece) bool {
	return PieceRookQueen[piece]
}

func IsKn(piece Piece) bool {
	return PieceKnight[piece]
}

func IsKi(piece Piece) bool {
	return PieceKing[piece]
}
