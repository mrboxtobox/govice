package board

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// TODO: Vice takes around 3.9 seconds (with some printing).

func MirrorEvalTest() {
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
			fen: parts[0],
		})
		// fmt.Println(count)

		if scanner.Err() != nil {
			log.Fatal(err)
		}
		count++
	}

	for _, p := range positions {
		pos := &Board{}
		ParseFEN(pos, p.fen)
		ev1 := EvalPosition(pos)
		MirrorBoard(pos)
		ev2 := EvalPosition(pos)
		if ev1 != ev2 {
			println("***")
			ParseFEN(pos, p.fen)
			PrintBoard(*pos)
			MirrorBoard(pos)
			PrintBoard(*pos)
		} else {
			// println("passed")
		}
	}
}
