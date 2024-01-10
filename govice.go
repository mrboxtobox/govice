package main

import (
	"math/rand"
	"time"

	"govice/board"
	// . "govice/board"
)

const (
	Name         = "govice"
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
	board.InitHashTable(b.HashTable, 64)
	board.InitPolyBook()

	board.UCILoop()
}
