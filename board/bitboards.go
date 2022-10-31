package board

import "fmt"

var (
	// See Matt Taylor's Folding trick.
	// TODO: This could be optimized depending on architecture.
	// TODO: Using De Bruijn Multiplication 64-bit mode.
	BitTable = [64]int8{
		0, 1, 48, 2, 57, 49, 28, 3,
		61, 58, 50, 42, 38, 29, 17, 4,
		62, 55, 59, 36, 53, 51, 43, 22,
		45, 39, 33, 30, 24, 18, 12, 5,
		63, 47, 56, 27, 60, 41, 37, 16,
		54, 35, 52, 21, 44, 32, 23, 11,
		46, 26, 40, 15, 34, 20, 31, 10,
		25, 14, 19, 9, 13, 8, 7, 6,
	}
)

// Takes the first least significant bit and returns the index and sets it to zero.
// Different from Vice's implementation.
// TODO: need to profile this.
// TODO: I had to rely on the 64-bit implementation to get this to work.
func PopBit(bb *uint64) int8 {
	res := BitTable[((*bb&-*bb)*0x03f79d71b4cb0a89)>>58]
	*bb &= (*bb - 1)
	return res
}

// Counts and returns the number of bits that are 1.
func CountBits(b uint64) int8 {
	var r int8
	for b != 0 {
		b &= b - 1
		r++
	}
	return r
}

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
