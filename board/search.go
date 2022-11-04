package board

import (
	"fmt"
	"time"
)

const (
	INF  = 30000
	MATE = 29000
)

type Mode int8

const (
	None Mode = iota
	UCIMODE
)

type SearchInfo struct {
	starttime time.Time
	stoptime  time.Time
	depth     int
	depthset  int
	timeset   bool
	movestogo int
	infinite  bool

	nodes int64

	quit    bool
	stopped bool

	GAME_MODE     Mode
	POST_THINKING bool

	// Fail high: Look at move ordering.
	fh float32

	// Fail high first
	fhf float32
}

func CheckUp(info *SearchInfo) {
	if info.timeset && time.Now().After(info.stoptime) {
		info.stopped = true
	}

	//  ReadInput(info);
}

//  func PickNextMove(moveNum int, list *MoveList) {
// 	 Move temp;
// 	 int index := 0;
// 	 int bestScore := 0;
// 	 int bestNum := moveNum;

// 	 for index := moveNum; index < list.Count; index++ {
// 		 if list.Moves[index].score > bestScore {
// 			 bestScore := list.Moves[index].score;
// 			 bestNum := index;
// 		 }
// 	 }

// 	 temp := list.Moves[moveNum];
// 	 list.Moves[moveNum] := list.Moves[bestNum];
// 	 list.Moves[bestNum] := temp;
//  }

func IsRepetition(pos *Board) bool {
	// Time since the 50-move rule was last reset. When it's reset, we can
	// never be in the same state we previously were.
	for index := pos.hisPly - int(pos.fiftyMove); index < pos.hisPly; index++ {
		assert(index >= 0 && index < MaxGameMoves)
		if pos.posKey == pos.history[index].positionKey {
			return true
		}
	}

	return false
}

func ClearForSearch(pos *Board, info *SearchInfo) {
	for index := 0; index < 13; index++ {
		for index2 := 0; index2 < BoardSquareCount; index2++ {
			// TODO: Use Golang clear slice = nil.
			pos.searchHistory[index][index2] = 0
		}
	}

	for index := 0; index < 2; index++ {
		for index2 := 0; index2 < MaxDepth; index2++ {
			pos.searchKillers[index][index2] = 0
		}
	}

	// TODO: Why do we clear at the beginning of search. Don't we use the PV Line
	// later?
	ClearHashTable(pos.HashTable)
	pos.ply = 0

	info.starttime = time.Now()
	info.stopped = false
	info.nodes = 0
	info.fh = 0
	info.fhf = 0
}

func Quiescence(alpha, beta int, pos *Board, info *SearchInfo) int {
	//  int MoveNum := 0;
	Legal := 0
	//  int Score;
	list := &MoveList{}
	//  MoveLIST list[1];

	CheckBoard(pos)
	assert(beta > alpha)

	if (info.nodes & 2047) == 0 {
		CheckUp(info)
	}

	info.nodes++

	if IsRepetition(pos) || pos.fiftyMove >= 100 {
		return 0
	}

	if pos.ply > MaxDepth-1 {
		return EvalPosition(pos)
	}

	Score := EvalPosition(pos)

	assert(Score > -INF && Score < INF)

	if Score >= beta {
		return beta
	}

	if Score > alpha {
		alpha = Score
	}

	GenerateAllCaps(pos, list)

	Score = -INF

	for MoveNum := 0; MoveNum < list.Count; MoveNum++ {
		//  PickNextMove(MoveNum, list);

		//  if (!MakeMove(pos, list.Moves[MoveNum].Move)){
		// 	 continue
		// 	}

		Legal++
		Score := -Quiescence(-beta, -alpha, pos, info)
		TakeMove(pos)

		if info.stopped {
			return 0
		}

		if Score > alpha {
			if Score >= beta {
				// We searched the best move first.
				if Legal == 1 {
					info.fhf++
				}
				// Otherwise, we just failed high.
				info.fh++
				return beta
			}
			alpha = Score
		}
	}

	return alpha
}

func AlphaBeta(alpha, beta, depth int, pos *Board, info *SearchInfo, DoNull bool) int {
	PvMove := NOMOVE
	Legal := 0
	OldAlpha := alpha
	BestMove := NOMOVE
	BestScore := -INF

	CheckBoard(pos)
	assert(beta > alpha)
	assert(depth >= 0)

	if depth <= 0 {
		return Quiescence(alpha, beta, pos, info)
	}

	if (info.nodes & 2047) == 0 {
		CheckUp(info)
	}

	info.nodes++

	if (IsRepetition(pos) || (pos.fiftyMove >= 100)) && (pos.ply > 0) {
		// Return draw estimate count.
		return 0
	}

	if pos.ply > MaxDepth-1 {
		return EvalPosition(pos)
	}

	InCheck := SqAttacked(*pos, pos.KingSq[pos.Side], pos.Side^1)
	if InCheck == true {
		depth++
	}

	// if ProbePvTable(pos, &PvMove, &Score, alpha, beta, depth) {
	// 	pos.HashTable.cut++
	// 	return Score
	// }

	// if DoNull && !InCheck && pos.ply && (pos.bigPce[pos.Side] > 0) && depth >= 4 {
	// 	MakeNullMove(pos)
	// 	Score := -AlphaBeta(-beta, -beta+1, depth-4, pos, info, false)
	// 	TakeNullMove(pos)
	// 	if info.stopped {
	// 		return 0
	// 	}

	// 	if Score >= beta && math.Abs(Score) < ISMATE {
	// 		info.nullCut++
	// 		return beta
	// 	}
	// }

	list := &MoveList{}
	GenerateAllMoves(pos, list)

	Score := -INF
	if PvMove != NOMOVE {
		for MoveNum := 0; MoveNum < list.Count; MoveNum++ {
			if list.Moves[MoveNum].Move == PvMove {
				list.Moves[MoveNum].score = 2000000
				break
			}
		}
	}

	for MoveNum := 0; MoveNum < list.Count; MoveNum++ {
		// PickNextMove(MoveNum, list)

		if !MakeMove(pos, list.Moves[MoveNum].Move) {
			continue
		}

		Legal++
		Score = -AlphaBeta(-beta, -alpha, depth-1, pos, info, true)
		TakeMove(pos)

		if info.stopped {
			return 0
		}

		if Score > BestScore {
			BestScore = Score
			BestMove = list.Moves[MoveNum].Move
			if Score > alpha {
				if Score >= beta {
					if Legal == 1 {
						// info.fhf++
					}
					// info.fh++

					if list.Moves[MoveNum].Move&MFLAGCAP == 0 {
						pos.searchKillers[1][pos.ply] = pos.searchKillers[0][pos.ply]
						pos.searchKillers[0][pos.ply] = list.Moves[MoveNum].Move
					}

					StorePvMove(pos, BestMove /*, beta, HFBETA, depth*/)
					return beta
				}
				alpha = Score

				if (list.Moves[MoveNum].Move & MFLAGCAP) == 0 {
					pos.searchHistory[pos.pieces[FromSQ(BestMove)]][ToSQ(BestMove)] += depth
				}
			}
		}
	}

	// Check mate or stalemate
	if Legal == 0 {
		if InCheck {
			return -INF + int(pos.ply) // TODO: Why not use MATE
		} else {
			return 0
		}
	}

	assert(alpha >= OldAlpha)

	// If we improved the best move, we can store the best move.
	if alpha != OldAlpha {
		StorePvMove(pos, BestMove /*, BestScore, HFEXACT, depth*/)
	} else {
		StorePvMove(pos, BestMove /*, alpha, HFALPHA, depth*/)
	}

	return alpha
}

func SearchPosition(pos *Board, info *SearchInfo) {
	bestMove := NOMOVE
	bestScore := -INF
	pvMoves := 0

	ClearForSearch(pos, info)

	//  if EngineOptions.UseBook {
	// 	 bestMove := GetBookMove(pos);
	//  }

	if bestMove == NOMOVE {
		for currentDepth := 1; currentDepth <= info.depth; currentDepth++ {
			bestScore = AlphaBeta(-INF, INF, currentDepth, pos, info, true)

			if info.stopped {
				break
			}

			pvMoves = GetPvLine(currentDepth, pos)
			bestMove = pos.PvArray[0]

			if info.GAME_MODE == UCIMODE {
				fmt.Printf("info score cp %d depth %d nodes %d time %d ",
					bestScore, currentDepth, info.nodes, time.Since(info.starttime))
			} else if info.POST_THINKING {
				fmt.Printf("score:%d depth:%d nodes:%d time:%d(ms) ",
					bestScore, currentDepth, info.nodes, time.Since(info.starttime))
			}

			if info.GAME_MODE == UCIMODE || info.POST_THINKING {
				pvMoves = GetPvLine(currentDepth, pos)
				fmt.Printf("pv")
				for pvNum := 0; pvNum < pvMoves; pvNum++ {
					fmt.Printf(" %s", PrMove(pos.PvArray[pvNum]))
					fmt.Printf("\n")
				}
			}
			fmt.Printf("\n")
			fmt.Printf("Ordering:%.2f\n", (info.fhf / info.fh))
		}
	}

	if info.GAME_MODE == UCIMODE {
		fmt.Printf("bestmove %s\n", PrMove(bestMove))
	} else {
		fmt.Printf("\n\n***!! Achebe makes move %s !!***\n\n", PrMove(bestMove))
		MakeMove(pos, bestMove)
		PrintBoard(*pos)
	}
}
