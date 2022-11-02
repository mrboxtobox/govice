package board

func GeneratePositionKey(board Board) uint64 {
	var key uint64

	// Pieces.
	for sq := 0; sq < BoardSquareCount; sq++ {
		piece := board.pieces[sq]
		if piece != EMPTY && piece != OFFBOARD {
			// fmt.Println(piece, WhitePawn, BlackKing)
			assert(piece >= WhitePawn && piece <= BlackKing)
			key ^= PieceKeys[piece][sq]
		}
	}

	// Side.
	if board.side == WHITE {
		key ^= SideKey
	}

	// En Passant
	if board.enPas != NO_SQ {
		// fmt.Println(board.enPassant, 0, BoardSquareCount)
		assert(board.enPas >= 0 && board.enPas <= BoardSquareCount)
		key ^= PieceKeys[EMPTY][board.enPas]
	}

	// Castling.
	assert(board.castlePerm >= 0 && board.castlePerm <= 15)
	key ^= CastleKeys[board.castlePerm]

	return key
}
