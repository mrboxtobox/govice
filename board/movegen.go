package board

var LoopSlidePce = [8]Piece{WhiteBishop, WhiteRook, WhiteQueen, 0, BlackBishop, BlackRook, BlackQueen, 0}
var LoopSlideIndex = [2]int{0, 4}
var LoopNonSlidePce = [6]Piece{WhiteKnight, WhiteKing, 0, BlackKnight, BlackKing, 0}
var LoopNonSlideIndex = [2]int{0, 3}

// Directions for each piece.
var PceDir = [13][8]int{
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{-8, -19, -21, -12, 8, 19, 21, 12},
	{-9, -11, 11, 9, 0, 0, 0, 0},
	{-1, -10, 1, 10, 0, 0, 0, 0},
	{-1, -10, 1, 10, -9, -11, 11, 9},
	{-1, -10, 1, 10, -9, -11, 11, 9},
	{0, 0, 0, 0, 0, 0, 0},
	{-8, -19, -21, -12, 8, 19, 21, 12},
	{-9, -11, 11, 9, 0, 0, 0, 0},
	{-1, -10, 1, 10, 0, 0, 0, 0},
	{-1, -10, 1, 10, -9, -11, 11, 9},
	{-1, -10, 1, 10, -9, -11, 11, 9},
}

// NumDir the number of moves per piece.
var NumDir = [13]int{0, 0, 8, 4, 4, 8, 8, 0, 8, 4, 4, 8, 8}

var VictimScore = [13]int{0, 100, 200, 300, 400, 500, 600, 100, 200, 300, 400, 500, 600}

var MvvLvaScores = [13][13]int{}

func InitMvvLva() {
	for Attacker := WhitePawn; Attacker <= BlackKing; Attacker++ {
		for Victim := WhitePawn; Victim <= BlackKing; Victim++ {
			MvvLvaScores[Victim][Attacker] = VictimScore[Victim] + 6 - (VictimScore[Attacker] / 100)
		}
	}
}

func PackMove(from, to int, cap, pro Piece, fl int) int {
	return from | (to << 7) | (int(cap) << 14) | (int(pro) << 20) | fl
}

func AddQuietMove(pos *Board, move int, list *MoveList) {

	list.Moves[list.Count].Move = move

	if pos.searchKillers[0][pos.ply] == move {
		list.Moves[list.Count].score = 900_000
	} else if pos.searchKillers[1][pos.ply] == move {
		list.Moves[list.Count].score = 800_000
	} else {
		list.Moves[list.Count].score = pos.searchHistory[pos.pieces[FromSQ(move)]][ToSQ(move)]
	}

	list.Moves[list.Count].score = 0
	list.Count++
}

func AddCaptureMove(pos *Board, move int, list *MoveList) {
	list.Moves[list.Count].Move = move
	list.Moves[list.Count].score = MvvLvaScores[Captured(move)][pos.pieces[FromSQ(move)]] + 1_000_000
	list.Count++
}

func AddEnPassantMove(pos *Board, move int, list *MoveList) {
	list.Moves[list.Count].Move = move
	// Pawn Takes Pawn according to MvvLvaScores.
	// Using 1000000 to make sure these are prioritized.
	list.Moves[list.Count].score = 105 + 1_000_000
	list.Count++
}

// Make functions performance intensive.
func AddWhitePawnCapMove(pos *Board, from, to int, cap Piece, list *MoveList) {
	assert(PieceValidEmpty(cap))
	assert(SqOnBoard(from))
	assert(SqOnBoard(to))

	if RanksBoard[from] == int(Rank7) {
		AddCaptureMove(pos, PackMove(from, to, cap, WhiteQueen, 0), list)
		AddCaptureMove(pos, PackMove(from, to, cap, WhiteRook, 0), list)
		AddCaptureMove(pos, PackMove(from, to, cap, WhiteBishop, 0), list)
		AddCaptureMove(pos, PackMove(from, to, cap, WhiteKnight, 0), list)
	} else {
		AddCaptureMove(pos, PackMove(from, to, cap, EMPTY, 0), list)
	}
}

func AddWhitePawnMove(pos *Board, from, to int, list *MoveList) {
	assert(SqOnBoard(from))
	assert(SqOnBoard(to))

	if RanksBoard[from] == int(Rank7) {
		AddQuietMove(pos, PackMove(from, to, EMPTY, WhiteQueen, 0), list)
		AddQuietMove(pos, PackMove(from, to, EMPTY, WhiteRook, 0), list)
		AddQuietMove(pos, PackMove(from, to, EMPTY, WhiteBishop, 0), list)
		AddQuietMove(pos, PackMove(from, to, EMPTY, WhiteKnight, 0), list)
	} else {
		AddQuietMove(pos, PackMove(from, to, EMPTY, EMPTY, 0), list)
	}
}

func AddBlackPawnCapMove(pos *Board, from, to int, cap Piece, list *MoveList) {
	assert(PieceValidEmpty(cap))
	assert(SqOnBoard(from))
	assert(SqOnBoard(to))

	if RanksBoard[from] == int(Rank2) {
		AddCaptureMove(pos, PackMove(from, to, cap, BlackQueen, 0), list)
		AddCaptureMove(pos, PackMove(from, to, cap, BlackRook, 0), list)
		AddCaptureMove(pos, PackMove(from, to, cap, BlackBishop, 0), list)
		AddCaptureMove(pos, PackMove(from, to, cap, BlackKnight, 0), list)
	} else {
		AddCaptureMove(pos, PackMove(from, to, cap, EMPTY, 0), list)
	}
}

func AddBlackPawnMove(pos *Board, from, to int, list *MoveList) {
	assert(SqOnBoard(from))
	assert(SqOnBoard(to))

	if RanksBoard[from] == int(Rank2) {
		AddQuietMove(pos, PackMove(from, to, EMPTY, BlackQueen, 0), list)
		AddQuietMove(pos, PackMove(from, to, EMPTY, BlackRook, 0), list)
		AddQuietMove(pos, PackMove(from, to, EMPTY, BlackBishop, 0), list)
		AddQuietMove(pos, PackMove(from, to, EMPTY, BlackKnight, 0), list)
	} else {
		AddQuietMove(pos, PackMove(from, to, EMPTY, EMPTY, 0), list)
	}
}

func GenerateAllMoves(pos *Board, list *MoveList) {
	// pce := EMPTY
	side := pos.Side
	// assert(CheckBoard(pos))
	list.Count = 0

	if side == WHITE {
		for pceNum := 0; int8(pceNum) < pos.pceNum[WhitePawn]; pceNum++ {
			sq := int(pos.pieceList[WhitePawn][pceNum])
			assert(SqOnBoard(sq))

			if pos.pieces[sq+10] == EMPTY {
				AddWhitePawnMove(pos, sq, sq+10, list)
				if RanksBoard[sq] == int(Rank2) && pos.pieces[sq+20] == EMPTY {
					AddQuietMove(pos, PackMove(sq, (sq+20), EMPTY, EMPTY, MFLAGPS), list)
				}
			}

			if !SQOFFBOARD(sq+9) && PieceColor[pos.pieces[sq+9]] == BLACK {
				AddWhitePawnCapMove(pos, sq, sq+9, pos.pieces[sq+9], list)
			}
			if !SQOFFBOARD(sq+11) && PieceColor[pos.pieces[sq+11]] == BLACK {
				AddWhitePawnCapMove(pos, sq, sq+11, pos.pieces[sq+11], list)
			}

			if pos.enPas != NO_SQ {
				if int(sq)+9 == int(pos.enPas) {
					AddEnPassantMove(pos, PackMove(sq, sq+9, EMPTY, EMPTY, MFLAGEP), list)
				}
				if int(sq)+11 == int(pos.enPas) {
					AddEnPassantMove(pos, PackMove(sq, sq+11, EMPTY, EMPTY, MFLAGEP), list)
				}
			}
		}

		if (pos.castlePerm & uint8(WKCA)) != 0 {
			if pos.pieces[F1] == EMPTY && pos.pieces[G1] == EMPTY {
				// We will check if King ends up in check during MakeMove.
				if !SqAttacked(*pos, E1, BLACK) && !SqAttacked(*pos, F1, BLACK) {
					// fmt.Printf("WKCA")
					AddQuietMove(pos, PackMove(int(E1), int(G1), EMPTY, EMPTY, MFLAGCA), list)
				}
			}
		}

		if (pos.castlePerm & uint8(WQCA)) != 0 {
			if pos.pieces[D1] == EMPTY && pos.pieces[C1] == EMPTY && pos.pieces[B1] == EMPTY {
				if !SqAttacked(*pos, E1, BLACK) && !SqAttacked(*pos, D1, BLACK) {
					// fmt.Printf("WQCA")
					AddQuietMove(pos, PackMove(int(E1), int(C1), EMPTY, EMPTY, MFLAGCA), list)
				}
			}
		}
	} else {
		for pceNum := 0; int8(pceNum) < pos.pceNum[BlackPawn]; pceNum++ {
			sq := pos.pieceList[BlackPawn][pceNum]
			assert(SqOnBoard(sq))

			if pos.pieces[sq-10] == EMPTY {
				AddBlackPawnMove(pos, sq, sq-10, list)
				if RanksBoard[sq] == int(Rank7) && pos.pieces[sq-20] == EMPTY {
					AddQuietMove(pos, PackMove(sq, (sq-20), EMPTY, EMPTY, MFLAGPS), list)
				}
			}

			if !SQOFFBOARD(sq-9) && PieceColor[pos.pieces[sq-9]] == WHITE {
				AddBlackPawnCapMove(pos, sq, sq-9, pos.pieces[sq-9], list)
			}
			if !SQOFFBOARD(sq-11) && PieceColor[pos.pieces[sq-11]] == WHITE {
				AddBlackPawnCapMove(pos, sq, sq-11, pos.pieces[sq-11], list)
			}

			if pos.enPas != NO_SQ {
				if sq-9 == int(pos.enPas) {
					AddEnPassantMove(pos, PackMove(sq, sq-9, EMPTY, EMPTY, MFLAGEP), list)
				}
				if sq-11 == int(pos.enPas) {
					AddEnPassantMove(pos, PackMove(sq, sq-11, EMPTY, EMPTY, MFLAGEP), list)
				}
			}
		}

		if (pos.castlePerm & uint8(BKCA)) != 0 {
			if pos.pieces[F8] == EMPTY && pos.pieces[G8] == EMPTY {
				if !SqAttacked(*pos, E8, WHITE) && !SqAttacked(*pos, F8, WHITE) {
					// fmt.Printf("BKCA")
					AddQuietMove(pos, PackMove(int(E8), int(G8), EMPTY, EMPTY, MFLAGCA), list)
				}
			}
		}

		if (pos.castlePerm & uint8(BQCA)) != 0 {
			if pos.pieces[D8] == EMPTY && pos.pieces[C8] == EMPTY && pos.pieces[B8] == EMPTY {
				if (!SqAttacked(*pos, E8, WHITE)) && (!SqAttacked(*pos, D8, WHITE)) {
					// fmt.Printf("BQCA")
					AddQuietMove(pos, PackMove(int(E8), int(C8), EMPTY, EMPTY, MFLAGCA), list)
				}
			}
		}
	}

	// wB, wR, wQ -- bB, bR, bQ
	// Sliding pieces

	pceIndex := LoopSlideIndex[side]
	pce := int(LoopSlidePce[pceIndex])
	pceIndex++
	for pce != 0 {
		assert(PieceValid(pce))
		// fmt.Printf("Piece: %c %v\n", PieceChar[pce], pce)

		for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
			sq := pos.pieceList[pce][pceNum]
			assert(SqOnBoard(sq))

			for index := 0; index < NumDir[pce]; index++ {
				dir := PceDir[pce][index]
				t_sq := sq + dir

				for !SQOFFBOARD(t_sq) {
					if pos.pieces[t_sq] != EMPTY {
						if PieceColor[pos.pieces[t_sq]] == (side ^ 1) {
							// fmt.Printf("%v %v %v %v\n", PieceColor[pos.pieces[t_sq]], side, side^1, pos.pieces[t_sq])
							// fmt.Printf("  Capture Move: %s\n", PrMove(PackMove(sq, t_sq, pos.pieces[t_sq], EMPTY, 0)))
							AddCaptureMove(pos, PackMove(sq, t_sq, pos.pieces[t_sq], EMPTY, 0), list)
						}
						break
					}
					// fmt.Printf("  Quiet Move: %s\n", PrMove(PackMove(sq, t_sq, pos.pieces[t_sq], EMPTY, 0)))
					AddQuietMove(pos, PackMove(sq, t_sq, EMPTY, EMPTY, 0), list)
					t_sq += dir
				}
			}
		}
		pce = int(LoopSlidePce[pceIndex])
		pceIndex++
	}

	pceIndex = LoopNonSlideIndex[side]
	pce = int(LoopNonSlidePce[pceIndex])
	pceIndex++
	for pce != 0 {
		assert(PieceValid(pce))
		// fmt.Printf("Piece: %c\n", PieceChar[pce])

		for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
			sq := pos.pieceList[pce][pceNum]
			assert(SqOnBoard(sq))

			for index := 0; index < NumDir[pce]; index++ {
				dir := PceDir[pce][index]
				t_sq := sq + dir

				if SQOFFBOARD(t_sq) {
					continue
				}

				if pos.pieces[t_sq] != EMPTY {
					if PieceColor[pos.pieces[t_sq]] == (side ^ 1) {
						// fmt.Printf("%v %v %v %v\n", PieceColor[pos.pieces[t_sq]], side, side^1, pos.pieces[t_sq])
						// fmt.Printf("  Capture Move: %s\n", PrMove(PackMove(sq, t_sq, pos.pieces[t_sq], EMPTY, 0)))
						AddCaptureMove(pos, PackMove(sq, t_sq, pos.pieces[t_sq], EMPTY, 0), list)
					}
					continue
				}
				// fmt.Printf("  Quiet Move: %s\n", PrMove(PackMove(sq, t_sq, pos.pieces[t_sq], EMPTY, 0)))
				AddQuietMove(pos, PackMove(sq, t_sq, EMPTY, EMPTY, 0), list)
			}
		}
		pce = int(LoopNonSlidePce[pceIndex])
		pceIndex++
	}
}

func MoveExists(pos *Board, move int) bool {
	list := &MoveList{}
	GenerateAllMoves(pos, list)

	for MoveNum := 0; MoveNum < list.Count; MoveNum++ {
		if !MakeMove(pos, list.Moves[MoveNum].Move) {
			continue
		}
		TakeMove(pos)
		if list.Moves[MoveNum].Move == move {
			return true
		}
	}
	return false
}

func GenerateAllCaps(pos *Board, list *MoveList) {
	// pce := EMPTY
	side := pos.Side
	// assert(CheckBoard(pos))
	list.Count = 0

	if side == WHITE {
		for pceNum := 0; int8(pceNum) < pos.pceNum[WhitePawn]; pceNum++ {
			sq := int(pos.pieceList[WhitePawn][pceNum])
			assert(SqOnBoard(sq))

			if !SQOFFBOARD(sq+9) && PieceColor[pos.pieces[sq+9]] == BLACK {
				AddWhitePawnCapMove(pos, sq, sq+9, pos.pieces[sq+9], list)
			}
			if !SQOFFBOARD(sq+11) && PieceColor[pos.pieces[sq+11]] == BLACK {
				AddWhitePawnCapMove(pos, sq, sq+11, pos.pieces[sq+11], list)
			}

			if pos.enPas != NO_SQ {
				if int(sq)+9 == int(pos.enPas) {
					AddEnPassantMove(pos, PackMove(sq, sq+9, EMPTY, EMPTY, MFLAGEP), list)
				}
				if int(sq)+11 == int(pos.enPas) {
					AddEnPassantMove(pos, PackMove(sq, sq+11, EMPTY, EMPTY, MFLAGEP), list)
				}
			}
		}
	} else {
		for pceNum := 0; int8(pceNum) < pos.pceNum[BlackPawn]; pceNum++ {
			sq := pos.pieceList[BlackPawn][pceNum]
			assert(SqOnBoard(sq))

			if !SQOFFBOARD(sq-9) && PieceColor[pos.pieces[sq-9]] == WHITE {
				AddBlackPawnCapMove(pos, sq, sq-9, pos.pieces[sq-9], list)
			}
			if !SQOFFBOARD(sq-11) && PieceColor[pos.pieces[sq-11]] == WHITE {
				AddBlackPawnCapMove(pos, sq, sq-11, pos.pieces[sq-11], list)
			}

			if pos.enPas != NO_SQ {
				if sq-9 == int(pos.enPas) {
					AddEnPassantMove(pos, PackMove(sq, sq-9, EMPTY, EMPTY, MFLAGEP), list)
				}
				if sq-11 == int(pos.enPas) {
					AddEnPassantMove(pos, PackMove(sq, sq-11, EMPTY, EMPTY, MFLAGEP), list)
				}
			}
		}
	}

	// wB, wR, wQ -- bB, bR, bQ
	// Sliding pieces

	pceIndex := LoopSlideIndex[side]
	pce := int(LoopSlidePce[pceIndex])
	pceIndex++
	for pce != 0 {
		assert(PieceValid(pce))
		// fmt.Printf("Piece: %c %v\n", PieceChar[pce], pce)

		for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
			sq := pos.pieceList[pce][pceNum]
			assert(SqOnBoard(sq))

			for index := 0; index < NumDir[pce]; index++ {
				dir := PceDir[pce][index]
				t_sq := sq + dir

				for !SQOFFBOARD(t_sq) {
					if pos.pieces[t_sq] != EMPTY {
						if PieceColor[pos.pieces[t_sq]] == (side ^ 1) {
							// fmt.Printf("%v %v %v %v\n", PieceColor[pos.pieces[t_sq]], side, side^1, pos.pieces[t_sq])
							// fmt.Printf("  Capture Move: %s\n", PrMove(PackMove(sq, t_sq, pos.pieces[t_sq], EMPTY, 0)))
							AddCaptureMove(pos, PackMove(sq, t_sq, pos.pieces[t_sq], EMPTY, 0), list)
						}
						break
					}
					t_sq += dir
				}
			}
		}
		pce = int(LoopSlidePce[pceIndex])
		pceIndex++
	}

	pceIndex = LoopNonSlideIndex[side]
	pce = int(LoopNonSlidePce[pceIndex])
	pceIndex++
	for pce != 0 {
		assert(PieceValid(pce))
		// fmt.Printf("Piece: %c\n", PieceChar[pce])

		for pceNum := 0; pceNum < int(pos.pceNum[pce]); pceNum++ {
			sq := pos.pieceList[pce][pceNum]
			assert(SqOnBoard(sq))

			for index := 0; index < NumDir[pce]; index++ {
				dir := PceDir[pce][index]
				t_sq := sq + dir

				if SQOFFBOARD(t_sq) {
					continue
				}

				if pos.pieces[t_sq] != EMPTY {
					if PieceColor[pos.pieces[t_sq]] == (side ^ 1) {
						// fmt.Printf("%v %v %v %v\n", PieceColor[pos.pieces[t_sq]], side, side^1, pos.pieces[t_sq])
						// fmt.Printf("  Capture Move: %s\n", PrMove(PackMove(sq, t_sq, pos.pieces[t_sq], EMPTY, 0)))
						AddCaptureMove(pos, PackMove(sq, t_sq, pos.pieces[t_sq], EMPTY, 0), list)
					}
					continue
				}
			}
		}
		pce = int(LoopNonSlidePce[pceIndex])
		pceIndex++
	}
}

// Move Ordering
// 1. PV Move
// 2. Cap . MvvLva
// 3. Killers
// 4. History score
