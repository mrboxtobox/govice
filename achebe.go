package main

import (
	"fmt"
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

	// fen1 := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	// fen2 := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"
	// fen3 := "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"
	// fen4 := "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"
	fen := "8/3q1p2/8/5P2/4Q3/8/8/8 w - - 0 1"
	b := board.Board{}

	board.ParseFEN(&b, fen)
	board.PrintBoard(b)

	// board.CheckBoard(&b)

	from := 6
	to := 12
	cap := board.WhiteRook
	promo := board.BlackRook
	move := (from) | (to << 7) | (int(cap) << 14) | (int(promo) << 20)

	fmt.Printf("\ndex: %d\nhex %x\n", move, move)
	fmt.Printf("\bin:%b\n", move)

	fmt.Printf("from:%d, to:%d, cap: %d, prom: %d\n", board.FromSQ(move), board.ToSQ(move), board.Captured(move), board.Promoted(move))

	move |= board.MFlagPawnStart
	fmt.Printf("Is pawn start: %v ", move&board.MFlagPawnStart)
}
