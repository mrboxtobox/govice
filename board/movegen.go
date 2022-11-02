package board

func PackMove(from, to int, cap, pro Piece, fl int) int {
	return from | (to << 7) | (int(cap) << 14) | (int(pro) << 20) | fl
}

func AddQuietMove(pos *Board, move int, list *MoveList) {
	list.Moves[list.Count].move = move
	list.Moves[list.Count].score = 0
	list.Count++
}

func AddCaptureMove(pos *Board, move int, list *MoveList) {
	list.Moves[list.Count].move = move
	list.Moves[list.Count].score = 0
	list.Count++
}

func AddEnPassantMove(pos *Board, move int, list *MoveList) {
	list.Moves[list.Count].move = move
	list.Moves[list.Count].score = 0
	list.Count++
}

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
	side := pos.side
	// int sq = 0, t_sq = 0;
	// int pceNum = 0;

	// int dir = 0;
	// int index = 0;
	// int pceIndex = 0;

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

		// if (pos.castlePerm & uint8(WKCA)) != 0 {
		// 	if pos.pieces[F1] == EMPTY && pos.pieces[G1] == EMPTY {
		// 		if !SqAttacked(*pos, E1, BLACK) && !SqAttacked(*pos, F1, BLACK) {
		// 			AddQuietMove(pos, PackMove(int(E1), int(G1), EMPTY, EMPTY, MFLAGCA), list)
		// 		}
		// 	}
		// }

		// if (pos.castlePerm & uint8(WQCA)) != 0 {
		// 	if pos.pieces[D1] == EMPTY && pos.pieces[C1] == EMPTY && pos.pieces[B1] == EMPTY {
		// 		if !SqAttacked(*pos, E1, BLACK) && !SqAttacked(*pos, D1, BLACK) {
		// 			AddQuietMove(pos, PackMove(int(E1), int(C1), EMPTY, EMPTY, MFLAGCA), list)
		// 		}
		// 	}
		// }
	} else {
		// for pceNum := 0; int8(pceNum) < pos.pceNum[BlackPawn]; pceNum++ {
		// 	sq := pos.pieceList[BlackPawn][pceNum]
		// 	assert(SqOnBoard(sq))

		// 	if pos.pieces[sq-10] == EMPTY {
		// 		AddBlackPawnMove(pos, sq, sq-10, list)
		// 		if RanksBoard[sq] == int(Rank7) && pos.pieces[sq-20] == EMPTY {
		// 			AddQuietMove(pos, PackMove(sq, (sq-20), EMPTY, EMPTY, MFLAGPS), list)
		// 		}
		// 	}

		// 	if !SQOFFBOARD(sq-9) && PieceColor[pos.pieces[sq-9]] == WHITE {
		// 		AddBlackPawnCapMove(pos, sq, sq-9, pos.pieces[sq-9], list)
		// 	}
		// 	if !SQOFFBOARD(sq-11) && PieceColor[pos.pieces[sq-11]] == WHITE {
		// 		AddBlackPawnCapMove(pos, sq, sq-11, pos.pieces[sq-11], list)
		// 	}

		// 	if pos.enPas != NO_SQ {
		// 		if sq-9 == int(pos.enPas) {
		// 			AddEnPassantMove(pos, PackMove(sq, sq-9, EMPTY, EMPTY, MFLAGEP), list)
		// 		}
		// 		if sq-11 == int(pos.enPas) {
		// 			AddEnPassantMove(pos, PackMove(sq, sq-11, EMPTY, EMPTY, MFLAGEP), list)
		// 		}
		// 	}
		// }

		// if pos.castlePerm & uint8(BKCA) {
		// 	if pos.pieces[F8] == EMPTY && pos.pieces[G8] == EMPTY {
		// 		if !SqAttacked(E8, WHITE, pos) && !SqAttacked(F8, WHITE, pos) {
		// 			AddQuietMove(pos, PackMove(E8, G8, EMPTY, EMPTY, MFLAGCA), list);
		// 		}
		// 	}
		// }

		// if(pos.castlePerm & uint8(BQCA)) {
		// 	if(pos.pieces[D8] == EMPTY && pos.pieces[C8] == EMPTY && pos.pieces[B8] == EMPTY) {
		// 		if((!SqAttacked(E8, WHITE, pos)) && (!SqAttacked(D8, WHITE, pos))) {
		// 			AddQuietMove(pos, PackMove(E8, C8, EMPTY, EMPTY, MFLAGCA), list);
		// 		}
		// 	}
		// }
	}

	// pceIndex = LoopSlideIndex[side];
	// pce = LoopSlidePce[pceIndex++];
	// for pce != 0 {
	// 	assert(PieceValid(pce));

	// 	for pceNum = 0; pceNum < pos.pceNum[pce]; ++pceNum {
	// 		sq = pos.pieceList[pce][pceNum];
	// 		assert(SqOnBoard(sq));

	// 		for index = 0; index < NumDir[pce]; ++index {
	// 			dir = PceDir[pce][index];
	// 			t_sq = sq + dir;

	// 			 for !SQOFFBOARD(t_sq) {

	// 				if(pos.pieces[t_sq] != EMPTY {
	// 					if(PieceColor[pos.pieces[t_sq]] == (side ^ 1) {
	// 						AddCaptureMove(pos, PackMove(sq, t_sq, pos.pieces[t_sq], EMPTY, 0), list);
	// 					}
	// 					break;
	// 				}
	// 				AddQuietMove(pos, PackMove(sq, t_sq, EMPTY, EMPTY, 0), list);
	// 				t_sq += dir;
	// 			}
	// 		}
	// 	}
	// 	pce = LoopSlidePce[pceIndex++];
	// }

	// pceIndex = LoopNonSlideIndex[side];
	// pce = LoopNonSlidePce[pceIndex++];
	//  for pce != 0 {
	// 	assert(PieceValid(pce));

	// 	for pceNum = 0; pceNum < pos.pceNum[pce]; ++pceNum {
	// 		sq = pos.pieceList[pce][pceNum];
	// 		assert(SqOnBoard(sq));

	// 		for index = 0; index < NumDir[pce]; ++index {
	// 			dir = PceDir[pce][index];
	// 			t_sq = sq + dir;

	// 			if(SQOFFBOARD(t_sq))
	// 				continue;

	// 			if(pos.pieces[t_sq] != NoPiece {
	// 				if(PieceColor[pos.pieces[t_sq]] == (side ^ 1) {
	// 					AddCaptureMove(pos, PackMove(sq, t_sq, pos.pieces[t_sq], NoPiece, 0), list);
	// 				}
	// 				continue;
	// 			}
	// 			AddQuietMove(pos, PackMove(sq, t_sq, NoPiece, NoPiece, 0), list);
	// 		}
	// 	}
	// 	pce = LoopNonSlidePce[pceIndex++];
	// }
}

// func GenerateAllCaps(pos *Board, list *MoveList) {
// 	int pce = NoPiece;
// 	int side = pos.side;
// 	int sq = 0, t_sq = 0;
// 	int pceNum = 0;

// 	int dir = 0;
// 	int index = 0;
// 	int pceIndex = 0;

// 	assert(CheckBoard(pos));
// 	list.count = 0;

// 	if(side == White) {
// 		for pceNum = 0; pceNum < pos.pceNum[WhitePawn]; ++pceNum) {
// 			sq = pos.pieceList[WhitePawn][pceNum];
// 			assert(SqOnBoard(sq));

// 			if(!SQOFFBOARD(sq + 9) && PieceColor[pos.pieces[sq + 9]] == Black) {
// 				AddWhitePawnCapMove(pos, sq, sq + 9, pos.pieces[sq + 9], list);
// 			}
// 			if(!SQOFFBOARD(sq + 11) && PieceColor[pos.pieces[sq + 11]] == Black) {
// 				AddWhitePawnCapMove(pos, sq, sq + 11, pos.pieces[sq + 11], list);
// 			}

// 			if(pos.enPas != NoSquare) {
// 				if(sq + 9 == pos.enPas) {
// 					AddEnPassantMove(pos, PackMove(sq, sq + 9, NoPiece, NoPiece, MFLAGEP), list);
// 				}
// 				if(sq + 11 == pos.enPas) {
// 					AddEnPassantMove(pos, PackMove(sq, sq + 11, NoPiece, NoPiece, MFLAGEP), list);
// 				}
// 			}
// 		}

// 	} else {
// 		for pceNum = 0; pceNum < pos.pceNum[bP]; ++pceNum) {
// 			sq = pos.pieceList[bP][pceNum];
// 			assert(SqOnBoard(sq));

// 			if(!SQOFFBOARD(sq - 9) && PieceColor[pos.pieces[sq - 9]] == White) {
// 				AddBlackPawnCapMove(pos, sq, sq - 9, pos.pieces[sq - 9], list);
// 			}
// 			if(!SQOFFBOARD(sq - 11) && PieceColor[pos.pieces[sq - 11]] == White) {
// 				AddBlackPawnCapMove(pos, sq, sq - 11, pos.pieces[sq - 11], list);
// 			}

// 			if(pos.enPas != NoSquare) {
// 				if(sq - 9 == pos.enPas) {
// 					AddEnPassantMove(pos, PackMove(sq, sq - 9, NoPiece, NoPiece, MFLAGEP), list);
// 				}
// 				if(sq - 11 == pos.enPas) {
// 					AddEnPassantMove(pos, PackMove(sq, sq - 11, NoPiece, NoPiece, MFLAGEP), list);
// 				}
// 			}
// 		}

// 	}

// 	pceIndex = LoopSlideIndex[side];
// 	pce = LoopSlidePce[pceIndex++];
// 	 for pce != 0 {
// 		assert(PieceValid(pce));

// 		for pceNum = 0; pceNum < pos.pceNum[pce]; ++pceNum {
// 			sq = pos.pieceList[pce][pceNum];
// 			assert(SqOnBoard(sq));

// 			for index = 0; index < NumDir[pce]; ++index {
// 				dir = PceDir[pce][index];
// 				t_sq = sq + dir;

// 				 for !SQOFFBOARD(t_sq) {

// 					if(pos.pieces[t_sq] != NoPiece {
// 						if(PieceColor[pos.pieces[t_sq]] == (side ^ 1) {
// 							AddCaptureMove(pos, PackMove(sq, t_sq, pos.pieces[t_sq], NoPiece, 0), list);
// 						}
// 						break;
// 					}
// 					t_sq += dir;
// 				}
// 			}
// 		}
// 		pce = LoopSlidePce[pceIndex++];
// 	}

// 	pceIndex = LoopNonSlideIndex[side];
// 	pce = LoopNonSlidePce[pceIndex++];
// 	 for pce != 0 {
// 		assert(PieceValid(pce));

// 		for pceNum = 0; pceNum < pos.pceNum[pce]; ++pceNum {
// 			sq = pos.pieceList[pce][pceNum];
// 			assert(SqOnBoard(sq));

// 			for index = 0; index < NumDir[pce]; ++index {
// 				dir = PceDir[pce][index];
// 				t_sq = sq + dir;

// 				if(SQOFFBOARD(t_sq))
// 					continue;

// 				if(pos.pieces[t_sq] != NoPiece {
// 					if(PieceColor[pos.pieces[t_sq]] == (side ^ 1) {
// 						AddCaptureMove(pos, PackMove(sq, t_sq, pos.pieces[t_sq], NoPiece, 0), list);
// 					}
// 					continue;
// 				}
// 			}
// 		}
// 		pce = LoopNonSlidePce[pceIndex++];
// 	}
// }
// Footer
