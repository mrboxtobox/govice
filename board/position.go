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
		// fmt.Println(pos.pieces, index)
		if piece != OFFBOARD && piece != EMPTY {
			color := PieceColor[piece]
			if BigPiece[piece] {
				pos.bigPieceCounts[color]++
			}
			if MajorPiece[piece] {
				pos.majorPieceCounts[color]++
			}
			if MinorPiece[piece] {
				pos.minorPieceCounts[color]++
			}

			// TODO: What happens when color is BOTH.
			pos.material[color] += PieceValue[piece]

			// TODO: We update the piece list and increase the count. Ch 18.
			pos.pieceList[piece][pos.pceNum[piece]] = index
			pos.pceNum[piece] += 1

			if piece == WhiteKing {
				pos.KingSq[WHITE] = Square(index)
			}
			if piece == BlackKing {
				pos.KingSq[BLACK] = Square(index)
			}
			sq := int8(index)

			// TODO: Need to verify this pointer is correctly passed.
			if piece == WhitePawn {
				SetBit(&pos.Pawns[WHITE], int(SQ64(sq)))
				SetBit(&pos.Pawns[Both], int(SQ64(sq)))
			} else if piece == BlackPawn {
				SetBit(&pos.Pawns[BLACK], int(SQ64(sq)))
				SetBit(&pos.Pawns[Both], int(SQ64(sq)))
			}

		}
	}
}

func ParseFEN(pos *Board, fen string) {
	ResetBoard(pos)

	rank := Rank8
	file := FileA
	piece := EMPTY
	var count int
	var curr int

	for rank >= Rank1 && curr < len(fen) {
		count = 1
		switch string(fen[curr]) {
		case "p", "b", "n", "r", "k", "q", "P", "B", "N", "R", "K", "Q":
			piece = strToPiece[string(fen[curr])]
		case "1", "2", "3", "4", "5", "6", "7", "8":
			piece = EMPTY
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
			if piece != EMPTY {
				pos.pieces[sq120] = piece
			}
			file++
		}
		curr++
	}

	assert(fen[curr] == 'w' || fen[curr] == 'b')
	if fen[curr] == 'w' {
		pos.Side = WHITE
	} else {
		pos.Side = BLACK
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
			pos.castlePerm |= uint8(WKCA)
		case 'Q':
			pos.castlePerm |= uint8(WQCA)
		case 'k':
			pos.castlePerm |= uint8(BKCA)
		case 'q':
			pos.castlePerm |= uint8(BQCA)
		}
	}

	// En passant.
	assert(pos.castlePerm >= 0 && pos.castlePerm <= 15)

	if fen[curr] != '-' {
		file := File(int(fen[curr] - 'a'))
		rank := Rank(int(fen[curr+1] - '1'))

		pos.enPas = Square(FileRankTo120Square(file, rank))
	}

	pos.posKey = GeneratePositionKey(*pos)

	UpdateListsMaterial(pos)
}
