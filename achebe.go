package main

import (
	"fmt"
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
	// fen := "rnbqkbnr/p1p1p3/3p3p/1p1p4/2P1Pp2/8/PP1P1PpP/RNBQKB1R b KQkq - 0 1"
	fen := "5k2/1n6/4n3/6N1/8/3N4/8/5K2 b - - 0 1"
	board.ParseFEN(&b, fen)
	fmt.Println(b.Side)

	list := board.MoveList{
		Moves: [256]board.Move{},
	}
	board.GenerateAllMoves(&b, &list)
	board.PrintMoveList(&list)
}
