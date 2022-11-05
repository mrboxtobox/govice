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

	for i, p := range positions {
		// if i > 1 {
		// 	continue
		// }
		pos := &Board{}
		ParseFEN(pos, p.fen)
		// ParseFEN(pos, "1k2r2r/Bpq3pp/3b4/3Bp3/8/7b/PPP1QP2/R3R1K1 w - - 0 1")
		ev1 := EvalPosition(pos)
		// fmt.Println(ev1)
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
		if i%10 == 0 {
			println(i)
		}
	}
}
