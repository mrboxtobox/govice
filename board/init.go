package board

import (
	"math/rand"
)

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
var PieceColor = [13]Color{BOTH, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, BLACK, BLACK, BLACK, BLACK, BLACK, BLACK}

var FilesBrd [BoardSquareCount]int
var RanksBoard [BoardSquareCount]int

func InitFilesRankBoard() {
	for index := 0; index < BoardSquareCount; index++ {
		FilesBrd[index] = int(OFFBOARD)
		RanksBoard[index] = int(OFFBOARD)
	}

	for rank := Rank1; rank <= Rank8; rank++ {
		for file := FileA; file <= FileH; file++ {
			sq := FileRankTo120Square(file, rank)
			FilesBrd[sq] = int(file)
			RanksBoard[sq] = int(rank)
		}
	}

	// for index := 0; index < BoardSquareCount; index++ {
	// 	if index%10 == 0 && index != 0 {
	// 		fmt.Printf("\n")
	// 	}
	// 	fmt.Printf("%4d", FilesBoard[index])
	// }

	// fmt.Printf("\n")
	// for index := 0; index < BoardSquareCount; index++ {
	// 	if index%10 == 0 && index != 0 {
	// 		fmt.Printf("\n")
	// 	}
	// 	fmt.Printf("%4d", RankBoard[index])
	// }
}

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
	InitFilesRankBoard()
	InitEvalMasks()
	InitMvvLva()
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

func InitEvalMasks() {
	var sq, tsq int

	for sq = 0; sq < 8; sq++ {
		FileBBMask[sq] = 0
		RankBBMask[sq] = 0
	}

	// Create the mask for getting File A, B, C.
	// Just a line of 1s
	for r := Rank8; r >= Rank1; r-- {
		for f := FileA; f <= FileH; f++ {
			sq = int(r)*8 + int(f)
			FileBBMask[f] |= (1 << sq)
			RankBBMask[r] |= (1 << sq)
		}
	}

	for sq = 0; sq < 64; sq++ {
		IsolatedMask[sq] = 0
		BlackPassedMask[sq] = 0
		WhitePassedMask[sq] = 0
	}

	for sq = 0; sq < 64; sq++ {
		tsq = sq + 8

		for tsq < 64 {
			WhitePassedMask[sq] |= (1 << tsq)
			tsq += 8
		}

		tsq = sq - 8
		for tsq >= 0 {
			BlackPassedMask[sq] |= (1 << tsq)
			tsq -= 8
		}

		if FilesBrd[SQ120(int8(sq))] > int(FileA) {
			// Set to the right.
			IsolatedMask[sq] |= FileBBMask[FilesBrd[SQ120(int8(sq))]-1]

			// Go left
			tsq = sq + 7
			for tsq < 64 {
				WhitePassedMask[sq] |= (1 << tsq)
				tsq += 8
			}

			// Go left
			tsq = sq - 9
			for tsq >= 0 {
				BlackPassedMask[sq] |= (1 << tsq)
				tsq -= 8
			}
		}

		// Set to the left.
		if FilesBrd[SQ120(int8(sq))] < int(FileH) {
			IsolatedMask[sq] |= FileBBMask[FilesBrd[SQ120(int8(sq))]+1]

			// Go right
			tsq = sq + 9
			for tsq < 64 {
				WhitePassedMask[sq] |= (1 << tsq)
				tsq += 8
			}

			// Go left
			tsq = sq - 7
			for tsq >= 0 {
				BlackPassedMask[sq] |= (1 << tsq)
				tsq -= 8
			}
		}
	}

	for sq := 0; sq < 64; sq++ {
		// println(sq)
		// PrintBitBoard(IsolatedMask[sq])
		// PrintBitBoard(BlackPassedMask[sq])
		// println("\n\n")
		// PrintBitBoard(FileBBMask[sq])
		// PrintBitBoard(RankBBMask[sq])
	}
}
