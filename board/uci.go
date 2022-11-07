package board

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var debug = false

const MAX_HASH = 100
const DEFAULT_HASH = 64

// Iterates until we get a space.
func GetWord(line string) string {
	var res string
	var idx int
	if len(line) == 0 {
		return res
	}

	for idx < len(line) && line[idx] != ' ' {
		res = res + string(line[idx])
		idx++
	}
	return res
}

// go depth 6 wtime 1800 btime 1000 binc 1000 winc 1000 movetim 1000 movestogo
func ParseGo(line string, info *SearchInfo, pos *Board) {
	depth := -1
	movestogo := 30
	movetime := -1
	sideTimed := -1
	inc := 0
	info.timeset = false
	sideTimed = -1
	var err error
	// Get values after each item.
	// go wtime 180000 btime 180000
	// fmt.Printf("go2: %q", line)
	if idx := strings.Index(line, "btime"); idx != -1 && pos.Side == WHITE {
		val := GetWord(line[idx+6:])
		if sideTimed, err = strconv.Atoi(val); err != nil {
			log.Fatal(err)
		}
	}

	if idx := strings.Index(line, "btime"); idx != -1 && pos.Side == BLACK {
		val := GetWord(line[idx+6:])
		if sideTimed, err = strconv.Atoi(val); err != nil {
			log.Fatal(err)
		}
	}

	if idx := strings.Index(line, "winc"); idx != -1 && pos.Side == WHITE {
		val := GetWord(line[idx+5:])
		if inc, err = strconv.Atoi(GetWord(val)); err != nil {
			log.Fatal(err)
		}
	}

	if idx := strings.Index(line, "binc"); idx != -1 && pos.Side == BLACK {
		val := GetWord(line[idx+5:])
		if inc, err = strconv.Atoi(val); err != nil {
			log.Fatal(err)
		}
	}

	if idx := strings.Index(line, "movestogo"); idx != -1 {
		val := GetWord(line[idx+10:])
		if movestogo, err = strconv.Atoi(val); err != nil {
			log.Fatal(err)
		}
	}

	if idx := strings.Index(line, "movetime"); idx != -1 {
		val := GetWord(line[idx+9:])
		if movetime, err = strconv.Atoi(val); err != nil {
			log.Fatal(err)
		}
	}

	if idx := strings.Index(line, "depth"); idx != -1 {
		val := GetWord(line[idx+6:])
		if depth, err = strconv.Atoi(val); err != nil {
			log.Fatal(err)
		}
	}

	if movetime != -1 {
		sideTimed = movetime
		movestogo = 1
	}

	info.StartTime = time.Now()
	info.Depth = depth

	if sideTimed != -1 {
		info.timeset = true
		sideTimed /= movestogo
		sideTimed -= 50 // Buffer
		info.stoptime = info.StartTime.Add(time.Duration(sideTimed+inc) * time.Millisecond)
	}

	if depth == -1 {
		info.Depth = MaxDepth
	}

	fmt.Printf("time:%d depth:%d timeset:%v\n", sideTimed, info.Depth, info.timeset)
	SearchPosition(pos, info)
}

// Reads UCI commands from stdin and outputs to stdout.
func UCILoop() {
	pos := &Board{
		HashTable: &PVTable{},
	}
	info := &SearchInfo{
		GAME_MODE: UCIMODE,
	}

	// 	= New()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("id name achebe")
	fmt.Println("id author mrboxtobox")
	fmt.Printf("option name Hash type spin default 64 min 4 max %d\n", MAX_HASH)
	fmt.Println("option name Book type check default true")
	fmt.Println("uciok")

	InitHashTable(pos.HashTable, 64)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "isready") {
			fmt.Println("readyok")
			continue
		} else if strings.HasPrefix(line, "position") {
			ParsePosition(line, pos)
		} else if strings.HasPrefix(line, "ucinewgame") {
			ParsePosition("position startpos", pos)
		} else if strings.HasPrefix(line, "go") {
			// fmt.Println("ggo: %v", line)
			ParseGo(line, info, pos)
		} else if strings.HasPrefix(line, "uci") {
			reply("id name achebe")
			reply("id author mrboxtobox")
			reply("uciok")
		} else if strings.HasPrefix(line, "quit") {
			info.quit = true
			break
		} else if strings.HasPrefix(line, "setoption name Hash value ") {
			MB, err := strconv.Atoi((line[26:]))
			if err != nil {
				log.Fatal(err)
			}
			MB = int(math.Max(math.Min(float64(MB), 4), 2048))

		}
		if info.quit {
			break
		}

		if scanner.Err() != nil {
			panic("bad scanner")
		}
	}
}

const START_FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

// position fen rnbqkb1r/pppp1ppp/8/4P3/6n1/7P/PPPNPPP1/R1BQKBNR b KQkq - 0 1
func ParsePosition(line string, pos *Board) {
	startIndex := 9
	line = line[startIndex:]
	if strings.HasPrefix(line, "startpos") {
		ParseFEN(pos, START_FEN)
	} else {
		idx := strings.Index(line, "fen")
		if idx == -1 {
			ParseFEN(pos, START_FEN)
		} else {
			line = line[idx+4:] // For the "fen "
			ParseFEN(pos, line)
		}
	}
	idx := strings.Index(line, "moves")
	// fmt.Printf("%q", line)
	if idx != -1 {
		line = line[idx+6:]
		fmt.Printf("%q\n", line)
		moves := strings.Fields(line)
		for _, moveStr := range moves {
			move := ParseMove(moveStr, pos)
			if move == NOMOVE {
				fmt.Println("Cannot parse move")
				break
			}
			MakeMove(pos, move)
			pos.ply = 0
		}
	}
	PrintBoard(*pos)
}

func reply(msg string) {
	fmt.Println(msg)
}
