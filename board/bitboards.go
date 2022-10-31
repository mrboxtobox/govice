package board

import "fmt"

func PrintBitBoard(bb uint64) {
	var shift uint64 = 1

	fmt.Println()
	for rank := Rank8; rank >= Rank1; rank-- {
		for file := FileA; file <= FileH; file++ {
			// fmt.Printf("rank file %v %v\n", rank, file)
			sq := FileRankTo120Square(file, rank)
			sq64 := SQ64(sq)
			if (shift<<sq64)&bb != 0 {
				fmt.Printf("%s", "X")
			} else {
				fmt.Printf("%s", "-")
			}
		}
		fmt.Println()
	}
	fmt.Print("\n\n")
}
