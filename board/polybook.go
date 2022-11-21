package board

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"unsafe"
)

type PolyBookEntry struct {
	key    uint64
	move   uint16
	weight uint16
	learn  uint32
}

var (
	NumEntries = 0
	entries    = []PolyBookEntry{}
)

var PolyKindOfPiece = [13]int{
	-1, 1, 3, 5, 7, 9, 11, 0, 2, 4, 6, 8, 10,
}

type Options struct {
	UseBook bool
}

var EngineOptions = Options{}

const (
	path      = "board/gm2001.bin"
	entrySize = 16
)

func InitPolyBook() {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("os.Open(%v) failed: %v", path, err)
	}
	defer file.Close()

	EngineOptions.UseBook = false

	// Create a buffer for reading all PolyGlot entry.
	reader := bufio.NewReader(file)
	buf := make([]byte, 16)

	info, err := os.Stat(path)
	if err != nil {
		log.Fatalf("os.Stat(%v) failed: %v", path, err)
	}
	x := info.Size()

	numEntries := x / int64(unsafe.Sizeof(PolyBookEntry{}))
	// fmt.Printf("Found %d entries in book\n", numEntries)
	entries = make([]PolyBookEntry, numEntries)

	// Read all entries from the PolyGlot file.
	for {
		n, err := reader.Read(buf)
		// reader.Read() returns EOF at the end of a file.
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("reader.Read(%v) failed with error: %v", buf, err)
		}
		if n != entrySize {
			log.Fatalf("reader.Read(%v) got %d bytes; want %d bytes", buf, n, entrySize)
		}

		// 64-bit hash of a board position. First 8 bytes.
		key := binary.BigEndian.Uint64(buf[0:8])
		// From Zahak. Do the swap while reading.
		// 64-bit hash of a board position. First 8 bytes.
		// key := binary.BigEndian.Uint64(buf[0:8])
		// Bit field encoding the move. Next 2 bytes.
		move := binary.BigEndian.Uint16(buf[8:10])
		// Measure of the move's quality. Next 2 bytes.
		weight := binary.BigEndian.Uint16(buf[10:12])

		// We do not need the 'learn' field.
		entry := PolyBookEntry{
			key:    key,
			move:   uint16(move),
			weight: uint16(weight),
		}

		entries = append(entries, entry)
	}

	if numEntries > 0 {
		EngineOptions.UseBook = true
	}
}

func CleanPolyBook() {
	//
}

func HasPawnForCapture(board *Board) bool {
	var sqWithPawn = 0
	var targetPce Piece
	if board.Side == WHITE {
		targetPce = WhitePawn
	} else {
		targetPce = BlackPawn
	}

	if board.enPas != NO_SQ {
		if board.Side == WHITE {
			sqWithPawn = int(board.enPas - 10)
		} else {
			sqWithPawn = int(board.enPas + 10)
		}

		if board.pieces[sqWithPawn+1] == targetPce {
			return true
		} else if board.pieces[sqWithPawn-1] == targetPce {
			return true
		}
	}
	return false
}

func PolyKeyFromBoard(board *Board) uint64 {
	var sq, rank, file int
	var finalKey = uint64(0)
	var piece = EMPTY
	var polyPiece = 0
	var offset = 0

	for sq = 0; sq < BoardSquareCount; sq++ {
		piece = board.pieces[sq]
		if piece != EMPTY && piece != OFFBOARD {
			assert(piece >= WhitePawn && piece <= BlackKing)
			polyPiece = PolyKindOfPiece[piece]
			rank = RanksBoard[sq]
			file = FilesBrd[sq]

			finalKey ^= Random64Poly[(64*polyPiece)+(8*rank)+file]
		}
	}

	offset = 768
	if (board.castlePerm & uint8(WKCA)) != 0 {
		finalKey ^= Random64Poly[offset+0]
	}
	if (board.castlePerm & uint8(WQCA)) != 0 {
		finalKey ^= Random64Poly[offset+1]
	}
	if (board.castlePerm & uint8(BKCA)) != 0 {
		finalKey ^= Random64Poly[offset+2]
	}
	if (board.castlePerm & uint8(BQCA)) != 0 {
		finalKey ^= Random64Poly[offset+3]
	}

	offset = 772
	if HasPawnForCapture(board) == true {
		file = FilesBrd[board.enPas]
		finalKey ^= Random64Poly[offset+file]
	}

	if board.Side == WHITE {
		finalKey ^= Random64Poly[780]
	}

	return finalKey
}

// toMove maps a PolyGlot move to its corresponding chess representation.
// 'move' encodes the following in 16 bits (bit 0 is the least significant bit).
// bits                meaning
// ===================================
// 0,1,2               to file
// 3,4,5               to rank
// 6,7,8               from file
// 9,10,11             from rank
// 12,13,14            promotion piece

func ConvertPolyMoveToInternalMove(polyMove uint16, board *Board) int {
	var ff = (polyMove >> 6) & 7
	var fr = (polyMove >> 9) & 7
	var tf = (polyMove >> 0) & 7
	var tr = (polyMove >> 3) & 7
	var pp = (polyMove >> 12) & 7
	var promChar = 'q'

	var moveString string
	if pp == 0 {
		moveString = fmt.Sprintf("%c%c%c%c",
			FileChar[ff],
			RankChar[fr],
			FileChar[tf],
			RankChar[tr])
	} else {
		switch pp {
		case 1:
			promChar = 'n'
		case 2:
			promChar = 'b'
		case 3:
			promChar = 'r'
		}

		moveString = fmt.Sprintf("%c%c%c%c%c",
			FileChar[ff],
			RankChar[fr],
			FileChar[tf],
			RankChar[tr],
			promChar)
	}

	return ParseMove(moveString, board)
}

func GetBookMove(board *Board) uint16 {
	// var PolyBookEntry *en try
	var move uint16
	const MAXBOOKMOVES = 32
	var bookMoves = []uint16{} /* !! */
	// var weight = []uint16{}
	// var cumWeight = uint16(0)
	var tempMove = NOMOVE
	var count = 0

	var polyKey = PolyKeyFromBoard(board)

	for _, entry := range entries {
		if polyKey == entry.key {
			move = entry.move
			tempMove = ConvertPolyMoveToInternalMove(move, board)
			if tempMove != NOMOVE {
				bookMoves = append(bookMoves, uint16(tempMove))
				count++
				if count > MAXBOOKMOVES {
					break
				}
			}
		}
	}

	if count != 0 {
		// for _, move := range bookMoves {
		// 	// fmt.Println(PrMove(int(move)))
		// }
		return bookMoves[rand.Intn(len(bookMoves))]
	} else {
		return NOMOVE
	}
}
