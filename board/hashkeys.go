package board

func GeneratePositionKey(board Board) uint64 {
	var key uint64

	// Pieces.
	for sq := 0; sq < BoardSquareCount; sq++ {
		piece := board.pieces[sq]
		if piece != NoPiece && piece != OffBoard {
			// fmt.Println(piece, WhitePawn, BlackKing)
			assert(piece >= WhitePawn && piece <= BlackKing)
			key ^= PieceKeys[piece][sq]
		}
	}

	// Side.
	if board.side == White {
		key ^= SideKey
	}

	// En Passant
	if board.enPassant != NoSquare {
		// fmt.Println(board.enPassant, 0, BoardSquareCount)
		assert(board.enPassant >= 0 && board.enPassant <= BoardSquareCount)
		key ^= PieceKeys[NoPiece][board.enPassant]
	}

	// Castling.
	assert(board.castlePermissions >= 0 && board.castlePermissions <= 15)
	key ^= CastleKeys[board.castlePermissions]

	return key
}
