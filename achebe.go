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
	// fen := "5k2/1n6/4n3/6N1/8/3N4/8/5K2 b - - 0 1"

	// bishops := "6k1/1b6/4n3/8/1n4B1/1B3N2/1N6/2b3K1 w - - 0 1"
	rooks := "6k1/8/5r2/8/1nR5/5N2/8/6K1 w - - 0 1"
	// queens := "6k1/1b6/4n3/8/1n4B1/1B3N2/1N6/2b3K1 w - - 0 1"
	board.ParseFEN(&b, rooks)
	fmt.Println(b.Side)

	list := board.MoveList{
		Moves: [256]board.Move{},
	}
	board.GenerateAllMoves(&b, &list)
	board.PrintMoveList(&list)
}
