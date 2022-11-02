package board

import "math/rand"

var Square120ToSquare64 [BoardSquareCount]int8
var Square64ToSquare120 [64]int8
var SetMask [64]uint64
var ClearMask [64]uint64

// Hash keys
var PieceKeys [13][BoardSquareCount]uint64
var SideKey uint64
var CastleKeys [16]uint64

var BigPiece = [13]bool{false, false, true, true, true, true, true, false, true, true, true, true, true}
var MajorPiece = [13]bool{false, false, false, false, true, true, true, false, false, false, true, true, true}
var MinorPiece = [13]bool{false, false, true, true, false, false, false, false, true, true, false, false, false}
var PieceValue = [13]int{0, 100, 325, 325, 550, 1_000, 50_000, 100, 325, 325, 550, 1_000, 50_000}
var PieceColor = [13]Color{Both, White, White, White, White, White, Black, Black, Black, Black, Black, Black, Black}

func InitHashKeys() {
	for index := 0; index < 13; index++ {
		for index2 := 0; index2 < BoardSquareCount; index2++ {
			PieceKeys[index][index2] = rand.Uint64()
		}
	}
	SideKey = rand.Uint64()
	for index := 0; index < 13; index++ {
		CastleKeys[index] = rand.Uint64()
	}
}

func InitBitMasks() {
	for index := 0; index < 64; index++ {
		SetMask[index] = 0
		ClearMask[index] = 0
	}
	for index := 0; index < 64; index++ {
		SetMask[index] |= (1 << index)
		ClearMask[index] = ^SetMask[index] // Bitwise Complement
	}
}

func Init() {
	InitSquare120ToSquare64()
	InitBitMasks()
	InitHashKeys()
}

func InitSquare120ToSquare64() {
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
