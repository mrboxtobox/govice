package main

import (
	"fmt"
	"log"

	"achebe/board"
)

const (
	Name = "Achebe"
)

func initialize() {

}

func main() {
	log.Print("Hello world.")
	board.Init()

	var playBitBoard uint64 = 0
	fmt.Printf("Start:\n")
	board.PrintBitBoard(playBitBoard)

	playBitBoard |= (1 << board.SQ64(int8(board.A1)))
	playBitBoard |= (1 << board.SQ64(int8(board.H8)))
	playBitBoard |= (1 << board.SQ64(int8(board.D2)))
	playBitBoard |= (1 << board.SQ64(int8(board.D3)))
	playBitBoard |= (1 << board.SQ64(int8(board.D4)))
	board.PrintBitBoard(playBitBoard)

	fmt.Printf("Count: %v\n", board.CountBits(playBitBoard))
	index := board.PopBit(&playBitBoard)
	fmt.Printf("Index: %v\n", index)
	board.PrintBitBoard(playBitBoard)

	var sq64 int
	for playBitBoard != 0 {
		sq64 = int(board.PopBit(&playBitBoard))
		fmt.Printf("Popped: %d\n", sq64)
		fmt.Printf("Count: %v\n", board.CountBits(playBitBoard))
		board.PrintBitBoard(playBitBoard)
	}
}
