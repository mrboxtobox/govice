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

// var ParseMove(char *ptrChar, S_BOARD *pos) {
// 	var from, to;
// 	var MoveNum = 0, Move = 0, PromPce = EMPTY;
// 	S_MOVELIST list[1];

// 	if(ptrChar[1] > '8' || ptrChar[1] < '1') return NOMOVE;
// 	if(ptrChar[3] > '8' || ptrChar[3] < '1') return NOMOVE;
// 	if(ptrChar[0] > 'h' || ptrChar[0] < 'a') return NOMOVE;
// 	if(ptrChar[2] > 'h' || ptrChar[2] < 'a') return NOMOVE;

// 	from = FR2SQ(ptrChar[0] - 'a', ptrChar[1] - '1');
// 	to = FR2SQ(ptrChar[2] - 'a', ptrChar[3] - '1');

// 	ASSERT(SqOnBoard(from) && SqOnBoard(to));

// 	GenerateAllMoves(pos, list);

// 	for(MoveNum = 0; MoveNum < list->count; ++MoveNum) {
// 		Move = list->moves[MoveNum].move;
// 		if(FromSQ(Move) == from && ToSQ(Move) == to) {
// 			PromPce = Promoted(Move);
// 			if(PromPce != EMPTY) {
// 				if(IsRQ(PromPce) && !IsBQ(PromPce) && ptrChar[4] == 'r') {
// 					return Move;
// 				} else if(!IsRQ(PromPce) && IsBQ(PromPce) && ptrChar[4] == 'b') {
// 					return Move;
// 				} else if(IsRQ(PromPce) && IsBQ(PromPce) && ptrChar[4] == 'q') {
// 					return Move;
// 				} else if(IsKn(PromPce) && ptrChar[4] == 'n') {
// 					return Move;
// 				}
// 				continue;
// 			}
// 			return Move;
// 		}
// 	}

// 	return NOMOVE;
// }

// void PrintMoveList(const S_MOVELIST *list) {
// 	var index = 0;
// 	var score = 0;
// 	var move = 0;
// 	printf("MoveList:\n");

// 	for(index = 0; index < list->count; ++index) {
// 		move = list->moves[index].move;
// 		score = list->moves[index].score;

// 		printf("Move:%d > %s (score:%d)\n", index + 1, PrMove(move), score);
// 	}

// 	printf("MoveList Total %d Moves.\n\n", list->count);
// }
