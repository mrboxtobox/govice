package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"achebe/board"
	// . "achebe/board"
)

const (
	Name = "Achebe"
)

func initialize() {

}

func main() {
	rand.Seed(time.Now().Unix())
	board.Init()

	b := &board.Board{
		HashTable: &board.PVTable{},
	}
	// TODO: Free memory.
	board.InitHashTable(b.HashTable, 10)

	// board.PerftMain()
	// fen := "n1n5/PPPk4/8/8/8/8/4Kppp/5N1N w - - 0 1"
	// fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	fen2 := "2rr3k/pp3pp1/1nnqbN1p/3pN3/2pP4/2P3Q1/PPB4P/R4RK1 w - -" // Mate in 3.
	board.ParseFEN(b, fen2)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		board.PrintBoard(*b)
		fmt.Printf("Please enter a move > ")

		if !scanner.Scan() {
			break
		}
		txt := scanner.Text()

		info := board.SearchInfo{
			Depth: 5,
		}
		board.SearchPosition(b, &info)
		if len(txt) == 0 {
			continue
		} else if txt[0] == 'q' {
			break
		} else if txt[0] == 't' {
			board.TakeMove(b)
		} else if txt[0] == 's' {
			info := board.SearchInfo{
				Depth: 4,
			}
			board.SearchPosition(b, &info)
		} else if len(txt) >= 4 {
			move := board.ParseMove(txt, b)
			if move != board.NOMOVE {
				board.StorePvMove(b, move)
				board.MakeMove(b, move)
				if board.IsRepetition(b) {
					fmt.Println("Repetition")
				}
				// board.PrintBoard(*b)
			} else {
				fmt.Printf("Move not parsed: %s", txt)
			}
		} else {
			fmt.Printf("Bad txt: %s\n", txt)
		}
	}
}
