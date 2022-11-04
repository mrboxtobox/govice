package board

import (
	"fmt"
	"unsafe"
)

const PvSize = 0x10 * 2

// const MAXDEPTH =

// Optimization option is size ordering.
type PVEntry struct {
	posKey uint64
	move   int
}

// TODO: This could be made more memory efficient.
type PVTable struct {
	// TODO: Should I use a pointer?
	pTable     []PVEntry
	numEntries int
}

// TODO: Make the TT global.

// func GetPvLine(depth int, pos *Board) {
// 	move := ProbePvMove(pos);
// 	count  = 0;

// 	assert(depth < MAXDEPTH);

// 	 for move != NOMOVE && count < depth {
// 		 assert(count < MAXDEPTH);

// 		 if(MoveExists(pos, move) {
// 			 MakeMove(pos, move);
// 			 pos.PvArray[count++] := move;
// 		 } else {
// 			 break;
// 		 }
// 		 move := ProbePvMove(pos);
// 	 }

// 	 for pos.ply > 0 {
// 		 TakeMove(pos);
// 	 }

// 	 return count;
//  }

func ClearHashTable(table *PVTable) {
	for _, tableEntry := range table.pTable {
		tableEntry.posKey = 0
		tableEntry.move = NOMOVE
		// TODO: These can be packed.
		//  tableEntry.depth := 0;
		//  tableEntry.score := 0;
		//  tableEntry.flags := 0;
	}
	//  table.newWrite=0;
}

func InitHashTable(table *PVTable, MB int) {
	HashSize := 0x100000 * MB
	table.numEntries = HashSize / int(unsafe.Sizeof(PVTable{}))
	// Subtract 2 to be safe. Could be removed.
	table.numEntries -= 2

	table.pTable = nil
	// TODO: Run garbage collection after?.
	table.pTable = make([]PVEntry, table.numEntries)

	ClearHashTable(table)
	fmt.Printf("HashTable init complete with %d entries\n", table.numEntries)
}

func StorePvMove(pos *Board, move /*,score, flags, depth*/ int) {

	index := pos.posKey % uint64(pos.HashTable.numEntries)

	assert(index >= 0 && index <= uint64(pos.HashTable.numEntries-1))
	//  assert(depth>=1&&depth<MAXDEPTH);
	//  assert(flags>=HFALPHA&&flags<=HFEXACT);
	//  assert(score>=-INF&&score<=INF);
	//  assert(pos.ply>=0&&pos.ply<MAXDEPTH);

	// if pos.HashTable.pTable[index].posKey == 0 {
	// 	pos.HashTable.newWrite++
	// } else {
	// 	pos.HashTable.overWrite++
	// }

	//  if(score > ISMATE) score += pos.ply;
	//  else if(score < -ISMATE) score -= pos.ply;

	pos.HashTable.pTable[index].move = move
	pos.HashTable.pTable[index].posKey = pos.posKey
	//  pos.HashTable.pTable[index].flags := flags;
	//  pos.HashTable.pTable[index].score := score;
	//  pos.HashTable.pTable[index].depth := depth;
}

func ProbePvTable(pos *Board /*var  *move,var  *score,var  alpha,var  beta,var  depth*/) int {
	index := pos.posKey % uint64(pos.HashTable.numEntries)

	assert(index >= 0 && index <= uint64(pos.HashTable.numEntries-1))
	//  assert(depth>=1&&depth<MAXDEPTH);
	//  assert(alpha<beta);
	//  assert(alpha>=-INF&&alpha<=INF);
	//  assert(beta>=-INF&&beta<=INF);
	//  assert(pos.ply>=0&&pos.ply<MAXDEPTH);

	if pos.HashTable.pTable[index].posKey == pos.posKey {
		return pos.HashTable.pTable[index].move
		//  *move := pos.HashTable.pTable[index].move;
		//  if pos.HashTable.pTable[index].depth >= depth {
		// 	 pos.HashTable.hit++;

		// 	 assert(pos.HashTable.pTable[index].depth>=1&&pos.HashTable.pTable[index].depth<MAXDEPTH);
		// 	 assert(pos.HashTable.pTable[index].flags>=HFALPHA&&pos.HashTable.pTable[index].flags<=HFEXACT);

		// 	 *score := pos.HashTable.pTable[index].score;
		// 	 if(*score > ISMATE) *score -= pos.ply;
		// 	 else if(*score < -ISMATE) *score += pos.ply;

		// 	 switch(pos.HashTable.pTable[index].flags {

		// 		 assert(*score>=-INF&&*score<=INF);

		// 		 case HFALPHA: if(*score<=alpha {
		// 			 *score=alpha;
		// 			 return TRUE;
		// 			 }
		// 			 break;
		// 		 case HFBETA: if(*score>=beta {
		// 			 *score=beta;
		// 			 return TRUE;
		// 			 }
		// 			 break;
		// 		 case HFEXACT:
		// 			 return TRUE;
		// 			 break;
		// 		 default: assert(FALSE); break;
		// 	 }
		//  }
	}

	return NOMOVE
}

// var  ProbePvMove(const pos *Board {

// 	var  index := pos.posKey % pos.HashTable.numEntries;
// 	 assert(index >= 0 && index <= pos.HashTable.numEntries - 1);

// 	 if( pos.HashTable.pTable[index].posKey == pos.posKey  {
// 		 return pos.HashTable.pTable[index].move;
// 	 }

// 	 return NOMOVE;
//  }
