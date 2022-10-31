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

	playBitBoard |= (1 << board.SQ64(int8(board.D2)))
	playBitBoard |= (1 << board.SQ64(int8(board.D3)))
	playBitBoard |= (1 << board.SQ64(int8(board.D4)))
	fmt.Printf("D2 Added:\n")
	board.PrintBitBoard(playBitBoard)
}
