# govice: Go Implementation of VICE (Video Instructional Chess Engine)
(Soon-to-be) UCI-compliant chess engine in Go.

## Local Development

Build from source. Binary will be written to `bin/govice`.
```
make build
```

Build and run
```
make run
```

## Acknowledgements

`govice` exists thanks to the following:

Instructional material
* [Chess Programming Wiki](https://www.chessprogramming.org/)
* [Computer Chess Programming Topics](https://web.archive.org/web/20071026090003/http://www.brucemo.com/compchess/programming/index.htm)
* [VICE](https://www.youtube.com/playlist?list=PLZ1QII7yudbc-Ky058TEaOstZHVbT-2hg)
* [How do modern chess engines work?](https://www.youtube.com/watch?v=pUyURF1Tqvg&t=567s&ab_channel=TNGTechnologyConsultingGmbH)
* [A Review of Game-Tree Pruning](https://webdocs.cs.ualberta.ca/~tony/OldPapers/icca.Mar1986.pp3-18.pdf)
* [Alpha-Beta with Sibling Prediction Pruning in Chess](https://homepages.cwi.nl/~paulk/theses/Carolus.pdf)
* [Chess Programming](https://www.youtube.com/channel/UCB9-prLkPwgvlKKqDgXhsMQ)

Libraries and tools
* [@notnil/chess](https://github.com/notnil/chess)
* [@cutechess/cutechess](https://github.com/cutechess/cutechess)

Open source engines
* [VICE](https://github.com/peterwankman/vice) (board representation, move generation, UCI, static evaluation)
* [Lc0](https://github.com/LeelaChessZero/lc0)
* [Zahak](https://github.com/amanjpro/zahak) (UCI, static evaluation, parsing EPDS)
* [CounterGo](https://github.com/ChizhovVadim/CounterGo) (UCI, architecture)
* [Combusken](https://github.com/mhib/combusken) (UCI, architecture)
* [FrankyGo](https://github.com/frankkopp/FrankyGo) (performance profiling)
* [TSCP](http://www.tckerrigan.com/Chess/TSCP/)
* [Mediocre](http://mediocrechess.blogspot.com/)
* [CPW-Engine](https://github.com/nescitus/cpw-engine)

Resources
* [Computer Chess Ratings List](http://www.computerchess.org.uk/ccrl/4040/about.html)

Test EPDs
* [Win at Chess](https://www.chessprogramming.org/Win_at_Chess)
* [Bratko-Kopec Test](https://www.chessprogramming.org/Bratko-Kopec_Test)
* [CCR One Hour Test](https://www.chessprogramming.org/CCR_One_Hour_Test)
* [Eigenmann Rapid Engine Test](https://www.chessprogramming.org/Eigenmann_Rapid_Engine_Test)
* [Kaufman Test](https://www.chessprogramming.org/Kaufman_Test)
* [LCT II](https://www.chessprogramming.org/LCT_II)
* [The Nolot Suite](https://www.chessprogramming.org/The_Nolot_Suite)
* [Null Move Test Positions](https://www.chessprogramming.org/Null_Move_Test-Positions)
* [Silent but Deadly](https://www.chessprogramming.org/Silent_but_Deadly)
* [Strategic Test Suite](https://www.chessprogramming.org/Strategic_Test_Suite)

Individuals
* Vincent Z.
* Thomas A.
* James H.
* Alex M.
* Nicholas B.
* Prof. Michael Genesereth ([CS227B])
* Antonio T. III ([CS227B] teammate)
* Aaron B. ([CS227B] teammate)

Chess Channels
* [GothamChess](https://www.youtube.com/c/GothamChess)
* [Daniel Naroditsky](https://www.youtube.com/c/DanielNaroditskyGM)

[CS227B]: http://logic.stanford.edu/classes/cs227/2015/index.html

SDG
