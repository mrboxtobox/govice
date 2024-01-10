package board

// TODO: Still searching significantly more nodes.
// r1b1kb1r/2pp1ppp/1np1q3/p3P3/2P5/1P6/PB1NQPPP/R3KB1R b KQkq - 0 1
// D8: ~21M // 26 seconds
import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
	StartTime time.Time
	stoptime  time.Time
	Depth     int
	depthset  int
	timeset   bool
	movestogo int
	infinite  bool

	nodes int64

	quit    bool
	stopped bool

	GAME_MODE     Mode
	POST_THINKING bool

	nullCut int

	// Fail high: Look at move ordering.
	fh float32

	// Fail high first
	fhf float32
}

func InputWaiting(info *SearchInfo) bool {
	reader := bufio.NewReader(os.Stdin)
	b, err := reader.Peek(16)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("after peak")

	// Read if data exists.
	if len(string(b)) > 0 {
		text, _ := reader.ReadString('\n')
		if text == "quit" {
			info.quit = true
		}
		return true
	}
	// Read the input.
	return false
}

// TODO: This could be solved by channels.
// If the GUI receives any input, assume it's being interrupted to stop.
func ReadInput(info *SearchInfo) {
	// if InputWaiting(info) {
	// 	fmt.Println("input waiting")

	// 	info.stopped = true
	// }
}

func CheckUp(info *SearchInfo) {
	if info.timeset && time.Now().After(info.stoptime) {
		// fmt.Println("Stopping time")
		info.stopped = true
	}

	ReadInput(info)
}

// Makes the best scoring move the next move.
// Previously, this was computed for all moves.
func PickNextMove(moveNum int, list *MoveList) {
	bestScore := 0
	bestNum := moveNum

	for index := moveNum; index < list.Count; index++ {
		if list.Moves[index].score > bestScore {
			bestScore = list.Moves[index].score
			bestNum = index
		}
	}

	temp := list.Moves[moveNum]
	list.Moves[moveNum] = list.Moves[bestNum]
	list.Moves[bestNum] = temp
}

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

	// Don't Clear the has table; keep it for later.

	pos.HashTable.overWrite = 0
	pos.HashTable.hit = 0
	pos.HashTable.cut = 0
	pos.HashTable.nullCut = 0
	pos.ply = 0

	info.stopped = false
	info.nodes = 0
	info.fh = 0
	info.fhf = 0
	info.nullCut = 0
}

func Quiescence(alpha, beta int, pos *Board, info *SearchInfo) int {
	CheckBoard(pos)
	assert(beta > alpha)

	if (info.nodes & 2047) == 0 {
		CheckUp(info)
	}

	info.nodes++

	if IsRepetition(pos) || pos.fiftyMove >= 100 {
		// TODO: Make draw more sensible.
		return 0
	}

	if pos.ply > MaxDepth-1 {
		return EvalPosition(pos)
	}

	Score := EvalPosition(pos)

	assert(Score > -INF && Score < INF)

	// We're not going to make any move that reduces beta so break early.
	// Standing pat (how you're doing without making a move)
	if Score >= beta {
		return beta
	}
	// https://github.com/ChrisWhittington/Chess-EPDs
	if Score > alpha {
		alpha = Score
	}

	list := &MoveList{}
	GenerateAllCaps(pos, list)

	Score = -INF
	// var BestMove int

	var Legal int
	for MoveNum := 0; MoveNum < list.Count; MoveNum++ {
		PickNextMove(MoveNum, list)

		if !MakeMove(pos, list.Moves[MoveNum].Move) {
			continue
		}

		// fmt.Printf("q %s %d\n", PrMove(list.Moves[MoveNum].Move), EvalPosition(pos))

		Legal++
		Score := -Quiescence(-beta, -alpha, pos, info)
		TakeMove(pos)

		if info.stopped {
			// TODO: Return 0? Not alpha?
			// Because we're breaking out here?
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
			// BestMove = list.Moves[MoveNum].Move
		}
	}

	return alpha
}

func AlphaBeta(alpha, beta, depth int, pos *Board, info *SearchInfo, DoNull bool) int {
	// fmt.Println("in ab")
	PvMove := NOMOVE
	Legal := 0
	OldAlpha := alpha
	BestMove := NOMOVE
	BestScore := -INF
	Score := -INF

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

	// Make at least 1 move.
	if (IsRepetition(pos) || (pos.fiftyMove >= 100)) && (pos.ply > 0) {
		// Return draw estimate count.
		// fmt.Println("Is repetition")
		return 0
	}

	if pos.ply > MaxDepth-1 {
		// fmt.Println("Reached max depth")
		return EvalPosition(pos)
	}

	InCheck := SqAttacked(*pos, pos.KingSq[pos.Side], pos.Side^1)

	// Add more depth when we're in check to allow us search deeper.
	// Depth extension to get outselves out of check. Otherwise, quiesence might
	// end at an opponents sacrifice when, maybe there's checkmate later on.
	if InCheck == true {
		depth++
	}

	// If we found a PVMove, just return that score.
	if ProbePvTable(pos, &PvMove, &Score, alpha, beta, depth) {
		pos.HashTable.cut++
		return Score
	}

	// Give opponent a free move (only once)
	// Check for zuzgwang with the bigPce check.
	if DoNull && !InCheck && (pos.ply > 0) && (pos.bigPieceCounts[pos.Side] > 0) && depth >= 4 {
		MakeNullMove(pos)
		Score := -AlphaBeta(-beta, -beta+1, depth-4, pos, info, false)
		TakeNullMove(pos)
		if info.stopped {
			return 0
		}

		if Score >= beta && math.Abs(float64(Score)) < ISMATE {
			info.nullCut++
			return beta
		}
	}

	list := &MoveList{}
	GenerateAllMoves(pos, list)

	Score = -INF
	if PvMove != NOMOVE {
		for MoveNum := 0; MoveNum < list.Count; MoveNum++ {
			if list.Moves[MoveNum].Move == PvMove {
				list.Moves[MoveNum].score = 2_000_000
				break
			}
		}
	}

	for MoveNum := 0; MoveNum < list.Count; MoveNum++ {
		PickNextMove(MoveNum, list)

		if !MakeMove(pos, list.Moves[MoveNum].Move) {
			continue
		}

		// fmt.Printf("a %s %d\n", PrMove(list.Moves[MoveNum].Move), EvalPosition(pos))

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
				// fmt.Printf("alpha cutoff a=%d, b=%d, s=%d, m=%s\n", alpha, beta, Score, PrMove(BestMove))
				if Score >= beta {
					// fmt.Printf("beta cutoff a=%d, b=%d, s=%d, m=%s\n", alpha, beta, Score, PrMove(BestMove))
					if Legal == 1 {
						info.fhf++
					}
					info.fh++

					// Caused a capture but increased beta?
					// Beta cutoffs still important since they reduce the tree.
					if list.Moves[MoveNum].Move&MFLAGCAP == 0 {
						// Shuffle down.
						pos.searchKillers[1][pos.ply] = pos.searchKillers[0][pos.ply]
						pos.searchKillers[0][pos.ply] = list.Moves[MoveNum].Move
					}

					StorePvMove(pos, BestMove, beta, HFBETA, depth)
					return beta
				}
				alpha = Score
				BestMove = list.Moves[MoveNum].Move
				if (list.Moves[MoveNum].Move & MFLAGCAP) == 0 {
					// Prioritize moves nearer the root of the tree.
					// Depth is decreasing so closer to root is higher.
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
			// fmt.Println("Returning bad non-legal move")
			return 0
		}
	}

	assert(alpha >= OldAlpha)

	// We didn't beat beta but improved alpha.
	// If we improved the best move, we can store the best move.
	if alpha != OldAlpha {
		// We have an exact flag.
		StorePvMove(pos, BestMove, BestScore, HFEXACT, depth)
	} else {
		StorePvMove(pos, BestMove, alpha, HFALPHA, depth)
	}

	return alpha
}

func SearchPosition(pos *Board, info *SearchInfo) {
	bestMove := NOMOVE
	bestScore := -INF
	pvMoves := 0

	ClearForSearch(pos, info)

	if EngineOptions.UseBook {
		bestMove = int(GetBookMove(pos))
	}

	if bestMove == NOMOVE {
		for currentDepth := 1; currentDepth <= info.Depth; currentDepth++ {
			// fmt.Println("doing ab")
			bestScore = AlphaBeta(-INF, INF, currentDepth, pos, info, true)
			if info.stopped {
				break
			}
			// fmt.Println("out ab")

			pvMoves = GetPvLine(currentDepth, pos)
			bestMove = pos.PvArray[0]

			if info.GAME_MODE == UCIMODE {
				fmt.Printf("info score cp %d depth %d nodes %d time %d ",
					bestScore, currentDepth, info.nodes, time.Since(info.StartTime).Milliseconds())
			} else if info.POST_THINKING {
				fmt.Printf("score:%d depth:%d nodes:%d time:%d(ms) ",
					bestScore, currentDepth, info.nodes, time.Since(info.StartTime).Milliseconds())
			} else {
				fmt.Printf("score:%d depth:%d nodes:%d time:%d(ms) ",
					bestScore, currentDepth, info.nodes, time.Since(info.StartTime).Milliseconds())
			}

			if info.GAME_MODE == UCIMODE || info.POST_THINKING || true {
				pvMoves = GetPvLine(currentDepth, pos)
				fmt.Printf("pv")
				for pvNum := 0; pvNum < pvMoves; pvNum++ {
					fmt.Printf(" %s", PrMove(pos.PvArray[pvNum]))
				}
				fmt.Printf("\n")
			}
			fmt.Printf("\n")
			// fmt.Printf("Ordering:%.2f\n", (info.fhf / info.fh))
		}
	}

	if info.GAME_MODE == UCIMODE {
		fmt.Printf("bestmove %s\n", PrMove(bestMove))
	} else {
		fmt.Printf("\n\n***!! govice makes move %s (%d) !!***\n\n", PrMove(bestMove), bestScore)
		// MakeMove(pos, bestMove)
		// PrintBoard(*pos)
	}
}
