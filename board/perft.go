package board

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://raw.githubusercontent.com/mishoo/queen.lisp/master/perftsuite.epd
type TestPosition struct {
	fen    string
	depths map[int]int
}

func PerftMain() {
	log.Print("Running tests...")
	// Read all entries.
	positions := []TestPosition{}
	f, err := os.Open("board/perft.epd")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		// fen am/bm; bm
		parts := strings.Split(scanner.Text(), " ;")
		// fmt.Println(len(parts))
		positions = append(positions, TestPosition{
			fen:    parts[0],
			depths: make(map[int]int),
		})
		// fmt.Println(count)

		for i := 1; i <= 5; i++ {
			p := strings.Fields(parts[i])[1]
			if v, err := strconv.Atoi(p); err != nil {
				panic(err)
			} else {
				positions[count].depths[i] = v
			}
		}

		if scanner.Err() != nil {
			log.Fatal(err)
		}
		count++
	}
	for d := 1; d <= 6; d++ {
		fmt.Printf("depth=%d\n", d)
		for _, p := range positions {
			pos := Board{}
			ParseFEN(&pos, p.fen)
			nodes := Perft(d, &pos)
			if nodes == p.depths[d] {
				// fmt.Printf("y: match\n")
			} else {
				fmt.Print(p.fen)
				fmt.Printf("n: got %d; want %d\n", nodes, p.depths[d])
			}
		}
	}
}

func Perft(depth int, pos *Board) int {
	var nodes int
	if depth == 0 {
		return 1
	}

	list := MoveList{
		Moves: [256]Move{},
	}
	GenerateAllMoves(pos, &list)
	for i := 0; i < list.Count; i++ {
		if !MakeMove(pos, list.Moves[i].Move) {
			continue
		}
		nodes += Perft(depth-1, pos)
		TakeMove(pos)
	}

	return nodes
}
