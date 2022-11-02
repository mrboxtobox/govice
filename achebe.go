package main

import (
	"math/rand"
	"time"

	"achebe/board"
)

const (
	Name = "Achebe"
)

func initialize() {

}

func main() {
	rand.Seed(time.Now().Unix())
	board.Init()

	b := board.Board{}
	fen := "rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P4/P1P1P2P/RNBQKBNR w KQkq - 0 1"
	board.ParseFEN(&b, fen)

	list := board.MoveList{
		Moves: [256]board.Move{},
	}
	board.GenerateAllMoves(&b, &list)
	board.PrintMoveList(&list)
}
