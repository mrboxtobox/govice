package board

import "fmt"

// TODO: Review this.
func PrSq(sq int) string {
	file := 'a' + FilesBrd[sq]
	rank := '1' + RanksBoard[sq]

	return fmt.Sprintf("%c%c", file, rank)
}

func PrMove(move int) string {
	ff := FilesBrd[FromSQ(move)]
	rf := RanksBoard[FromSQ(move)]
	ft := FilesBrd[ToSQ(move)]
	rt := RanksBoard[ToSQ(move)]

	promoted := Promoted(move)

	if promoted != 0 {
		pchar := 'q'
		if IsKn(Piece(promoted)) {
			pchar = 'n'
		} else if IsRQ(Piece(promoted)) && !IsBQ(Piece(promoted)) {
			pchar = 'r'
		} else if !IsRQ(Piece(promoted)) && IsBQ(Piece(promoted)) {
			pchar = 'b'
		}

		return fmt.Sprintf("%c%c%c%c%c", ('a' + ff), ('1' + rf), ('a' + ft), ('1' + rt), pchar)
	} else {
		return fmt.Sprintf("%c%c%c%c", ('a' + ff), ('1' + rf), ('a' + ft), ('1' + rt))
	}
}

func ParseMove(ptrChar string, pos *Board) int {
	list := MoveList{}

	if ptrChar[1] > '8' || ptrChar[1] < '1' {
		return NOMOVE
	}
	if ptrChar[3] > '8' || ptrChar[3] < '1' {
		return NOMOVE
	}
	if ptrChar[0] > 'h' || ptrChar[0] < 'a' {
		return NOMOVE
	}
	if ptrChar[2] > 'h' || ptrChar[2] < 'a' {
		return NOMOVE
	}

	from := FileRankTo120Square(File(ptrChar[0]-'a'), Rank(ptrChar[1]-'1'))
	to := FileRankTo120Square(File(ptrChar[2]-'a'), Rank(ptrChar[3]-'1'))

	assert(SqOnBoard(int(from)) && SqOnBoard(int(to)))

	GenerateAllMoves(pos, &list)

	for MoveNum := 0; MoveNum < list.Count; MoveNum++ {
		// fmt.Printf("Found move: %d", MoveNum)
		Move := list.Moves[MoveNum].Move
		if FromSQ(Move) == int(from) && ToSQ(Move) == int(to) {
			PromPce := Piece(Promoted(Move))
			if PromPce != EMPTY {
				if IsRQ(PromPce) && !IsBQ(PromPce) && ptrChar[4] == 'r' {
					// fmt.Printf("Found rq: %d", MoveNum)
					return Move
				} else if !IsRQ(PromPce) && IsBQ(PromPce) && ptrChar[4] == 'b' {
					return Move
				} else if IsRQ(PromPce) && IsBQ(PromPce) && ptrChar[4] == 'q' {
					return Move
				} else if IsKn(PromPce) && ptrChar[4] == 'n' {
					return Move
				}
				continue
			}
			return Move
		}
	}

	return NOMOVE
}

func PrintMoveList(list *MoveList) {
	fmt.Print("MoveList:\n")

	for index := 0; index < list.Count; index++ {
		move := list.Moves[index].Move
		score := list.Moves[index].Move

		fmt.Printf("Move:%d > %s (score:%d)\n", index+1, PrMove(move), score)
	}

	fmt.Printf("MoveList Total %d Moves.\n\n", list.Count)
}
