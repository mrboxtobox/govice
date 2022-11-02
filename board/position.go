package board

import (
	"log"
)

const (
	StartPos = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

var (
	strToPiece = map[string]Piece{
		"p": BlackPawn,
		"b": BlackBishop,
		"n": BlackKnight,
		"r": BlackRook,
		"k": BlackKing,
		"q": BlackQueen,
		"P": WhitePawn,
		"B": WhiteBishop,
		"N": WhiteKnight,
		"R": WhiteRook,
		"K": WhiteKing,
		"Q": WhiteQueen,
	}
)

func UpdateListsMaterial(pos *Board) {
	for index := 0; index < BoardSquareCount; index++ {
		piece := pos.pieces[index]
		if piece != OffBoard && piece != NoPiece {
			color := PieceColor[piece]
			if BigPiece[index] {
				pos.bigPieceCounts[piece]++
			}
			if MajorPiece[index] {
				pos.majorPieceCounts[piece]++
			}
			if MinorPiece[index] {
				pos.minorPieceCounts[piece]++
			}

			pos.material[color] += PieceValue[piece]

			// TODO: We update the piece list and increase the count. Ch 18.
			pos.pieceList[piece][pos.pieceCounts[piece]] = index
			pos.pieceCounts[piece] += 1

			if piece == WhiteKing {
				pos.kings[White] = Square(index)
			}
			if piece == BlackBishop {
				pos.kings[Black] = Square(index)
			}
		}
	}

}

func ParseFEN(pos *Board, fen string) {
	ResetBoard(pos)

	rank := Rank8
	file := FileA
	piece := NoPiece
	var count int
	var curr int

	for rank >= Rank1 && curr < len(fen) {
		count = 1
		switch string(fen[curr]) {
		case "p", "b", "n", "r", "k", "q", "P", "B", "N", "R", "K", "Q":
			piece = strToPiece[string(fen[curr])]
		case "1", "2", "3", "4", "5", "6", "7", "8":
			piece = NoPiece
			count = int(fen[curr] - '0')
		case "/", " ":
			rank--
			file = FileA
			curr++
			continue
		default:
			log.Fatalf("Error at %d while parsing fen: %v", curr, fen)
		}

		for i := 0; i < count; i++ {
			// fmt.Println(count, rank, file)
			sq64 := int8(rank)*8 + int8(file)
			sq120 := SQ120(sq64)
			if piece != NoPiece {
				pos.pieces[sq120] = piece
			}
			file++
		}
		curr++
	}

	assert(fen[curr] == 'w' || fen[curr] == 'b')
	if fen[curr] == 'w' {
		pos.side = White
	} else {
		pos.side = Black
	}

	curr += 2

	// TODO: Make tihs cleaner.
	for j := curr; j < curr+4; j++ {
		curr++
		if fen[j] == ' ' {
			break
		}

		switch fen[j] {
		case 'K':
			pos.castlePermissions |= uint8(WhiteKingCastle)
		case 'Q':
			pos.castlePermissions |= uint8(WhiteQueenCastle)
		case 'k':
			pos.castlePermissions |= uint8(BlackKingCastle)
		case 'q':
			pos.castlePermissions |= uint8(BlackQueenCastle)
		}
	}

	// En passant.
	assert(pos.castlePermissions >= 0 && pos.castlePermissions <= 15)

	if fen[curr] != '-' {
		file := File(int(fen[curr] - 'a'))
		rank := Rank(int(fen[curr+1] - '1'))

		pos.enPassant = Square(FileRankTo120Square(file, rank))
	}

	pos.positionKey = GeneratePositionKey(*pos)
}
