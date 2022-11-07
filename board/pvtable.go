package board

import (
	"fmt"
	"unsafe"
)

const PvSize = 0x10 * 2
const MaxDepth = 64
const ISMATE = INF - MaxDepth

// const MaxDepth =

type TTFlag int

const (
	HFNONE TTFlag = iota
	HFALPHA
	HFBETA
	HFEXACT
)

// Optimization option is size ordering.
type PVEntry struct {
	posKey uint64
	move   int
	// TODO: Where was this set?
	score int
	// These can be packed.
	depth int
	flags TTFlag
}

// TODO: This could be made more memory efficient.
type PVTable struct {
	// TODO: Should I use a pointer?
	pTable     []PVEntry
	numEntries int
	cut        int
	hit        int
	overWrite  int
	nullCut    int
	newWrite   int
}

// TODO: Make the TT global.

func GetPvLine(depth int, pos *Board) int {
	move := ProbePvMove(pos)
	count := 0

	assert(depth < MaxDepth)

	for move != NOMOVE && count < depth {
		assert(count < MaxDepth)

		// TODO: Move is done twice.
		if MoveExists(pos, move) {
			MakeMove(pos, move)
			pos.PvArray[count] = move
			count++
		} else {
			break
		}
		move = ProbePvMove(pos)
	}

	// Reset to initial state.
	for pos.ply > 0 {
		TakeMove(pos)
	}
	return count
}

func ClearHashTable(table *PVTable) {
	for _, tableEntry := range table.pTable {
		tableEntry.posKey = 0
		tableEntry.move = NOMOVE
		// TODO: These can be packed.
		tableEntry.depth = 0
		tableEntry.score = 0
		tableEntry.flags = 0
	}
	table.newWrite = 0
}

func InitHashTable(table *PVTable, MB int) {
	HashSize := 0x100000 * MB
	// TODO: Run garbage collection after?.
	table.numEntries = HashSize / int(unsafe.Sizeof(PVEntry{}))
	fmt.Printf("Hash table size: %d\n", table.numEntries)
	table.pTable = make([]PVEntry, table.numEntries)
	// Subtract 2 to be safe. Could be removed.
	table.numEntries -= 2
	assert(table.numEntries > 1)
	assert(len(table.pTable) > 1)

	// table.pTable = nil
	// table.pTable = make([]PVEntry, table.numEntries)

	ClearHashTable(table)
	// fmt.Printf("HashTable init complete with %d entries\n", table.numEntries)
}

func StorePvMove(pos *Board, move, score int, flags TTFlag, depth int) {

	index := pos.posKey % uint64(pos.HashTable.numEntries)

	assert(index >= 0 && index <= uint64(pos.HashTable.numEntries-1))
	assert(depth >= 1 && depth < MaxDepth)
	assert(flags >= HFALPHA && flags <= HFEXACT)
	assert(score >= -INF && score <= INF)
	assert(pos.ply >= 0 && pos.ply < MaxDepth)

	// TODO: Can change overwrite condition.
	if pos.HashTable.pTable[index].posKey == 0 {
		pos.HashTable.newWrite++
	} else {
		pos.HashTable.overWrite++
	}

	// Reset the mate score to be infinite since we might be using it shallower
	// in teh tree. It will be converted to the appropriate score late.
	if score > ISMATE {
		score += int(pos.ply)
	} else if score < -ISMATE {
		score -= int(pos.ply)
	}

	pos.HashTable.pTable[index].move = move
	pos.HashTable.pTable[index].posKey = pos.posKey
	pos.HashTable.pTable[index].flags = flags
	pos.HashTable.pTable[index].score = score
	pos.HashTable.pTable[index].depth = depth
}

func ProbePvTable(pos *Board, move *int, score *int, alpha, beta, depth int) bool {
	index := pos.posKey % uint64(pos.HashTable.numEntries)

	assert(index >= 0 && index <= uint64(pos.HashTable.numEntries-1))
	assert(depth >= 1 && depth < MaxDepth)
	assert(alpha < beta)
	assert(alpha >= -INF && alpha <= INF)
	assert(beta >= -INF && beta <= INF)
	assert(pos.ply >= 0 && pos.ply < MaxDepth)

	if pos.HashTable.pTable[index].posKey == pos.posKey {
		// TODO: We could return a pari, (move, score) instead and update score.

		*move = pos.HashTable.pTable[index].move
		// From Bruce Moreland's site.
		if pos.HashTable.pTable[index].depth >= depth {
			pos.HashTable.hit++
			assert(pos.HashTable.pTable[index].depth >= 1 && pos.HashTable.pTable[index].depth < MaxDepth)
			assert(pos.HashTable.pTable[index].flags >= HFALPHA && pos.HashTable.pTable[index].flags <= HFEXACT)

			*score = pos.HashTable.pTable[index].score
			if *score > ISMATE { // score will never be equal to ISMATE since we will make at least 1 move.
				*score -= int(pos.ply)
			} else if *score < -ISMATE {
				*score += int(pos.ply)
			}
			assert(*score >= -INF && *score <= INF)

			switch pos.HashTable.pTable[index].flags {
			// We have an exact score for the position
			case HFALPHA:
				// We aren't going to beat this alpha, just return the stored alpha.
				if *score <= alpha {
					*score = alpha
					return true
				}
			case HFBETA:
				// Beta (bettered) so we're guaranteed to beat the beta as well.
				// If the previous search's beta is greater than the current beta, we are
				// guaranteed to beat this beta (think back to the pruning). This node
				// does will basically be ignored.
				if *score >= beta {
					*score = beta
					return true
				}
			case HFEXACT:
				// We had an exact score for the position.
				return true
			default:
				assert(false)
			}
		}
	}
	return false
}

func ProbePvMove(pos *Board) int {
	index := pos.posKey % uint64(pos.HashTable.numEntries)
	assert(index >= 0 && index <= uint64(pos.HashTable.numEntries-1))
	if pos.HashTable.pTable[index].posKey == pos.posKey {
		return pos.HashTable.pTable[index].move
	}
	return NOMOVE
}
