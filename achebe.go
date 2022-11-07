package main

import (
	"math/rand"
	"time"

	"achebe/board"
	// . "achebe/board"
)

const (
	Name         = "Achebe"
	DEFAULT_HASH = 64
)

func initialize() {

}

func main() {
	rand.Seed(time.Now().Unix())
	board.Init()
	// board.MirrorEvalTest()

	b := &board.Board{
		HashTable: &board.PVTable{},
	}
	// TODO: Free memory.
	board.InitHashTable(b.HashTable, 16)

	// board.UCILoop()

	board.ParseFEN(b, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	info := &board.SearchInfo{
		GAME_MODE: board.UCIMODE,
		StartTime: time.Now(),
		Depth:     2,
	}
	board.SearchPosition(b, info)

	// board.ParseFEN(b, "6rk/7p/Q2R1p2/1p2n3/4b3/1P4NP/P1P2PPK/5q2 b - - 0 1")
	// info := &board.SearchInfo{
	// 	GAME_MODE: board.UCIMODE,
	// 	StartTime: time.Now(),
	// 	Depth:     6,
	// }
	// board.SearchPosition(b, info)

	// board.ParseFEN(b, "6rk/7p/Q2R1p2/1p2n3/4b3/1P4NP/P1P2PPK/2q5 b - - 0 1")
	// board.SearchPosition(b, info)
}
