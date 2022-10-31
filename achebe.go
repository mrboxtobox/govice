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

	for index := 0; index < board.BoardSquareCount; index++ {
		if index%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", board.Square120ToSquare64[index])
	}
	fmt.Println()

	for index := 0; index < 64; index++ {
		if index%8 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", board.Square64ToSquare120[index])
	}
}
