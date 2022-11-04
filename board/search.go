package board

const (
	INF  = 30000
	MATE = 29000
)

//  func CheckUp(info *SearchInfo) {
// 	 if info.timeset == true && GetTimeMs() > info.stoptime)
// 		 info.stopped := true;

// 	 ReadInput(info);
//  }

//  func PickNextMove(moveNum int, list *MoveList) {
// 	 Move temp;
// 	 int index := 0;
// 	 int bestScore := 0;
// 	 int bestNum := moveNum;

// 	 for index := moveNum; index < list.count; index++ {
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

//  func ClearForSearch(pos *Board, info *SearchInfo) {
// 	 int index := 0;
// 	 int index2 := 0;

// 	 for index := 0; index < 13; index++)
// 		 for index2 := 0; index2 < BRD_SQ_NUM; index++2)
// 			 pos.searchHistory[index][index2] := 0;

// 	 for index := 0; index < 2; index++)
// 		 for index2 := 0; index2 < MAXDEPTH; index++2)
// 			 pos.searchKillers[index][index2] := 0;

// 	 ClearHashTable(pos.HashTable);
// 	 pos.ply := 0;

// 	 info.stopped := 0;
// 	 info.nodes := 0;
// 	 info.fh := 0;
// 	 info.fhf := 0;
//  }

//  func Quiescence(int alpha, int beta, pos *Board, info *SearchInfo {
// 	 int MoveNum := 0;
// 	 int Legal := 0;
// 	 int Score;
// 	 MoveLIST list[1];

// 	 assert(CheckBoard(pos));
// 	 assert(beta > alpha);

// 	 if (info.nodes & 2047) == 0)
// 		 CheckUp(info);

// 	 info.nodes++;

// 	 if IsRepetition(pos) || pos.fiftyMove >= 100)
// 		 return 0;

// 	 if pos.ply > MAXDEPTH - 1)
// 		 return EvalPosition(pos);

// 	 Score := EvalPosition(pos);

// 	 assert(Score > -INF && Score < INF);

// 	 if Score >= beta)
// 		 return beta;

// 	 if Score > alpha)
// 		 alpha := Score;

// 	 GenerateAllCaps(pos, list);

// 	 Score := -INF;

// 	 for MoveNum := 0; MoveNum < list.count; ++MoveNum {
// 		 PickNextMove(MoveNum, list);

// 		 if (!MakeMove(pos, list.Moves[MoveNum].move))
// 			 continue;

// 		 Legal++;
// 		 Score := -Quiescence(-beta, -alpha, pos, info);
// 		 TakeMove(pos);

// 		 if info.stopped == true)
// 			 return 0;

// 		 if Score > alpha {
// 			 if Score >= beta {
// 				 if Legal == 1)
// 					 info.fhf++;
// 				 info.fh++;
// 				 return beta;
// 			 }
// 			 alpha := Score;
// 		 }
// 	 }

// 	 return alpha;
//  }

//  func AlphaBeta(int alpha, int beta, int depth, pos *Board, info *SearchInfo, int DoNull {
// 	 MoveLIST list[1];
// 	 int InCheck;
// 	 int Score := -INF;
// 	 int PvMove := NOMOVE;
// 	 int MoveNum := 0;
// 	 int Legal := 0;
// 	 int OldAlpha := alpha;
// 	 int BestMove := NOMOVE;
// 	 int BestScore := -INF;

// 	 assert(CheckBoard(pos));
// 	 assert(beta > alpha);
// 	 assert(depth >= 0);

// 	 if depth <= 0)
// 		 return Quiescence(alpha, beta, pos, info);

// 	 if (info.nodes & 2047) == 0)
// 		 CheckUp(info);

// 	 info.nodes++;

// 	 if (IsRepetition(pos) || pos.fiftyMove >= 100) && pos.ply)
// 		 return 0;

// 	 if pos.ply > MAXDEPTH - 1)
// 		 return EvalPosition(pos);

// 	 InCheck := SqAttacked(pos.KingSq[pos.side], pos.side ^ 1, pos);

// 	 if InCheck == true)
// 		 depth++;

// 	 if ProbeHashEntry(pos, &PvMove, &Score, alpha, beta, depth) == true {
// 		 pos.HashTable.cut++;
// 		 return Score;
// 	 }

// 	 if DoNull && !InCheck && pos.ply && (pos.bigPce[pos.side] > 0) && depth >= 4 {
// 		 MakeNullMove(pos);
// 		 Score := -AlphaBeta(-beta, -beta + 1, depth - 4, pos, info, false);
// 		 TakeNullMove(pos);
// 		 if info.stopped == true)
// 			 return 0;

// 		 if (Score >= beta && abs(Score) < ISMATE {
// 			 info.nullCut++;
// 			 return beta;
// 		 }
// 	 }

// 	 GenerateAllMoves(pos, list);

// 	 Score := -INF;

// 	 if PvMove != NOMOVE {
// 		 for MoveNum := 0; MoveNum < list.count; ++MoveNum {
// 			 if  list.Moves[MoveNum].move == PvMove {
// 				 list.Moves[MoveNum].score := 2000000;
// 				 break;
// 			 }
// 		 }
// 	 }

// 	 for MoveNum := 0; MoveNum < list.count; ++MoveNum {
// 		 PickNextMove(MoveNum, list);

// 		 if !MakeMove(pos,list.Moves[MoveNum].move))
// 			 continue;

// 		 Legal++;
// 		 Score := -AlphaBeta(-beta, -alpha, depth - 1, pos, info, true);
// 		 TakeMove(pos);

// 		 if info.stopped == true)
// 			 return 0;

// 		 if Score > BestScore {
// 			 BestScore := Score;
// 			 BestMove := list.Moves[MoveNum].move;
// 			 if Score > alpha {
// 				 if Score >= beta {
// 					 if Legal == 1)
// 						 info.fhf++;
// 					 info.fh++;

// 					 if !(list.Moves[MoveNum].move & MFLAGCAP) {
// 						 pos.searchKillers[1][pos.ply] := pos.searchKillers[0][pos.ply];
// 						 pos.searchKillers[0][pos.ply] := list.Moves[MoveNum].move;
// 					 }

// 					 StoreHashEntry(pos, BestMove, beta, HFBETA, depth);
// 					 return beta;
// 				 }
// 				 alpha := Score;

// 				 if !(list.Moves[MoveNum].move & MFLAGCAP))
// 					 pos.searchHistory[pos.pieces[FROMSQ(BestMove)]][TOSQ(BestMove)] += depth;
// 			 }
// 		 }
// 	 }

// 	 if Legal == 0 {
// 		 if InCheck {
// 			 return -INF + pos.ply;
// 		 } else {
// 			 return 0;
// 		 }
// 	 }

// 	 assert(alpha >= OldAlpha);

// 	 if alpha != OldAlpha {
// 		 StoreHashEntry(pos, BestMove, BestScore, HFEXACT, depth);
// 	 } else {
// 		 StoreHashEntry(pos, BestMove, alpha, HFALPHA, depth);
// 	 }

// 	 return alpha;
//  }

//  void SearchPosition(pos *Board, info *SearchInfo {
// 	 int bestMove := NOMOVE;
// 	 int bestScore := -INF;
// 	 int currentDepth := 0;
// 	 int pvMoves := 0;
// 	 int pvNum := 0;

// 	 ClearForSearch(pos, info);

// 	 if EngineOptions.UseBook == true)
// 		 bestMove := GetBookMove(pos);

// 	 if bestMove == NOMOVE {
// 		 for currentDepth := 1; currentDepth <= info.depth; ++currentDepth {
// 			 bestScore := AlphaBeta(-INF, INF, currentDepth, pos, info, true);

// 			 if info.stopped == true)
// 				 break;

// 			 pvMoves := GetPvLine(currentDepth, pos);
// 			 bestMove := pos.PvArray[0];

// 			 if info.GAME_MODE == UCIMODE {
// 				 printf("info score cp %d depth %d nodes %ld time %d ",
// 					 bestScore, currentDepth, info.nodes, GetTimeMs() - info.starttime);
// 			 } else if info.GAME_MODE == XBOARDMODE && info.POST_THINKING == true {
// 				 printf("%d %d %d %ld ",
// 					 currentDepth, bestScore, (GetTimeMs() - info.starttime) / 10, info.nodes);
// 			 } else if info.POST_THINKING == true {
// 				 printf("score:%d depth:%d nodes:%ld time:%d(ms) ",
// 					 bestScore, currentDepth, info.nodes, GetTimeMs() - info.starttime);
// 			 }

// 			 if info.GAME_MODE == UCIMODE || info.POST_THINKING == true {
// 				 pvMoves := GetPvLine(currentDepth, pos);
// 				 printf("pv");
// 				 for pvNum := 0; pvNum < pvMoves; ++pvNum)
// 					 printf(" %s", PrMove(pos.PvArray[pvNum]));
// 				 printf("\n");
// 			 }
// 		 }
// 	 }

// 	 if info.GAME_MODE == UCIMODE {
// 		 printf("bestmove %s\n", PrMove(bestMove));
// 	 } else if info.GAME_MODE == XBOARDMODE {
// 		 printf("move %s\n", PrMove(bestMove));
// 		 MakeMove(pos, bestMove);
// 	 } else {
// 		 printf("\n\n***!! Vice makes move %s !!***\n\n", PrMove(bestMove));
// 		 MakeMove(pos, bestMove);
// 		 PrintBoard(pos);
// 	 }
//  }
