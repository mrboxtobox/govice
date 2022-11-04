package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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

	b := &board.Board{}
	// board.PerftMain()
	// fen := "n1n5/PPPk4/8/8/8/8/4Kppp/5N1N w - - 0 1"
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	board.ParseFEN(b, fen)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		board.PrintBoard(*b)
		fmt.Printf("Please enter a move > ")

		if !scanner.Scan() {
			break
		}
		txt := scanner.Text()
		if len(txt) == 0 {
			continue
		} else if txt[0] == 'q' {
			break
		} else if txt[0] == 't' {
			board.TakeMove(b)
		} else if txt[0] == 'p' {
			// board.PerftTest(4, b)
			max := board.GetPvLine(4, b)
			fmt.Printf("PVLine: ")
			for num := 0; num < max; num++ {
				move := b.PvArray[num]
				fmt.Printf(" %s", board.PrMove(move))
			}
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
