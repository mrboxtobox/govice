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
	board.InitHashTable(b.HashTable, 64)
	board.InitPolyBook()

	board.UCILoop()
}
