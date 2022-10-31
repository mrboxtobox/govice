package main

import (
	"fmt"
	"log"
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
	log.Print("Hello world.")
	rand.Seed(time.Now().Unix())
	board.Init()

	p1 := rand.Int()
	p2 := rand.Int()
	p3 := rand.Int()
	p4 := rand.Int()

	fmt.Printf("P1: %X\n", p1)
	fmt.Printf("P2: %X\n", p2)
	fmt.Printf("P3: %X\n", p3)
	fmt.Printf("P4: %X\n", p4)
}
