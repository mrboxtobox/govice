package main

import (
	"log"
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
	log.Print("Hello world.")
	rand.Seed(time.Now().Unix())
	board.Init()

	fen1 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	fen2 := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"
	fen3 := "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"

	b := board.Board{}

	board.ParseFEN(&b, fen1)
	board.PrintBoard(b)

	board.ParseFEN(&b, fen2)
	board.PrintBoard(b)

	board.ParseFEN(&b, fen3)
	board.PrintBoard(b)
}
