package board

func SqOnBoard(sq int) bool {
	return FilesBrd[sq] != int(OFFBOARD)
}
func SideValid(side Color) bool {
	return (side == WHITE) || (side == BLACK)
}

func FileRankValid(fr int) bool {
	return fr >= 0 && fr <= 7
}
func PieceValidEmpty(pce Piece) bool {
	return pce >= EMPTY && pce <= BlackKing
}
func PieceValid(pce int) bool {
	return pce >= int(WhitePawn) && pce <= int(BlackKing)
}
