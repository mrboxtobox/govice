package main

import (
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
	rand.Seed(time.Now().Unix())
	board.Init()

	board.PerftMain()

	// b := board.Board{}
	// fen := "rnbqkbnr/p1p1p3/3p3p/1p1p4/2P1Pp2/8/PP1P1PpP/RNBQKB1R b KQkq - 0 1"
	// fen := "5k2/1n6/4n3/6N1/8/3N4/8/5K2 b - - 0 1"

	// bishops := "6k1/1b6/4n3/8/1n4B1/1B3N2/1N6/2b3K1 w - - 0 1"
	// rooks := "6k1/8/5r2/8/1nR5/5N2/8/6K1 w - - 0 1"
	// queens := "6k1/1b6/4n3/8/1n4B1/1B3N2/1N6/2b3K1 w - - 0 1"
	// castle := "3rk2r/8/8/8/8/8/6p1/R3K2R b KQk - 0 1"
	// complex := "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"

	// board.ParseFEN(&b, complex)
	// fmt.Println(b.Side)

	// list := board.MoveList{
	// 	Moves: [256]board.Move{},
	// }
	// board.GenerateAllMoves(&b, &list)
	// // board.PrintMoveList(&list)

	// input := bufio.NewScanner(os.Stdin)
	// // input.Scan()

	// var num int
	// for input.Scan() {
	// 	fmt.Println("input")
	// 	move := list.Moves[num].Move

	// 	if !board.MakeMove(&b, move) {
	// 		continue
	// 	}

	// 	fmt.Println("made")
	// 	board.PrintBoard(b)

	// 	board.TakeMove(&b)
	// 	fmt.Println("taken")
	// 	board.PrintBoard(b)
	// 	num++
	// }
}
