package board

import "fmt"

/*
                            3 F
0000 0000 0000 0000 0000 0111 1111 -> From: & 0x7F
0000 0000 0000 0011 1111 1000 0000 -> To: >> 7, & 0x7F
0000 0000 0011 1100 0000 0000 0000 -> Captured: >> 14, & F
0000 0000 0100 0000 0000 0000 0000 -> En Passant: 0x40000 (non zero)
0000 0000 1000 0000 0000 0000 0000 -> En Passant: 0x80000 (non zero)
0000 1111 0000 0000 0000 0000 0000 -> Promoted: >> 20, 0xF
0001 0000 0000 0000 0000 0000 0000 -> Castle:0x100000
*/

type Move struct {
	move  int
	score int
}

type MoveList struct {
	Moves [MaxPositionMoves]Move
	Count int
}

func PrintMoveList(list *MoveList) {
	for index := 0; index < list.Count; index++ {
		fmt.Printf("%d Move: %s\n", index, PrMove(list.Moves[index].move))
		if list.Moves[index].move&MFLAGCA != 0 {
			fmt.Println("---> Capture")
		}
	}
}

func FromSQ(move int) int {
	return move & 0x7F
}

func ToSQ(move int) int {
	return (move >> 7) & 0x7F
}

func Captured(m int) int {
	return (m >> 14) & 0xF
}

func Promoted(m int) int {
	return (m >> 20) & 0xF
}

const (
	MFLAGEP     = 0x40000
	MFLAGPS     = 0x80000
	MFlagCastle = 0x1000000

	MFLAGCA    = 0x7c000
	MFlagPromo = 0xF00000
)
