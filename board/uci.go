package board

// func ParseGo(line string, info *SearchInfo, pos *Board) {
// 	 depth := -1
// 	 movestogo := 30
// 	 movetime := -1
// 	 time := -1
// 	 inc := 0
// 	 info.timeset = false

// 	//  if((ptr := strstr(line,"INF")));

// 	//  if((ptr := strstr(line,"binc")) && pos.side == BLACK)
// 	// 	 inc := atoi(ptr + 5);
// 	//  if((ptr := strstr(line,"winc")) && pos.side == WHITE)
// 	// 	 inc := atoi(ptr + 5);
// 	//  if((ptr := strstr(line,"wtime")) && pos.side == WHITE)
// 	// 	 time := atoi(ptr + 6);
// 	//  if((ptr := strstr(line,"btime")) && pos.side == BLACK)
// 	// 	 time := atoi(ptr + 6);
// 	//  if((ptr := strstr(line,"movestogo")))
// 	// 	 movestogo := atoi(ptr + 10);
// 	//  if((ptr := strstr(line,"movetime")))
// 	// 	 movetime := atoi(ptr + 9);

// 	//  if((ptr := strstr(line,"depth")))
// 	// 	 depth := atoi(ptr + 6);

// 	//  if(movetime != -1) {
// 	// 	 time := movetime;
// 	// 	 movestogo := 1;
// 	//  }

// 	//  info.starttime := GetTimeMs();
// 	//  info.depth := depth;

// 	//  if(time != -1) {
// 	// 	 info.timeset := TRUE;
// 	// 	 time /= movestogo;
// 	// 	 time -= 50;
// 	// 	 info.stoptime := info.starttime + time + inc;
// 	//  }

// 	//  if(depth == -1)
// 	// 	 info.depth := MAXDEPTH;

// 	//  printf("time:%d start:%d stop:%d depth:%d timeset:%d\n",
// 	// 	 time, info.starttime, info.stoptime, info.depth, info.timeset);
// 	//  SearchPosition(pos, info);
//  }

// func Uci_Loop(pos *Board, info *SearchInfo) {
// 	 char line[INPUTBUFFER]
// 	  *ptrTrue := NULL
// 	 MB := 64;

// 	 info.GAME_MODE = UCIMODE;
// 	 info.quit = false;

// 	 setbuf(stdin, NULL);
// 	 setbuf(stdout, NULL);
// 	 printf("id name %s\n",NAME);
// 	 printf("id author Bluefever\n");
// 	 printf("option name Hash type spin default 64 min 4 max %d\n", MAX_HASH);
// 	 printf("option name Book type check default true\n");
// 	 printf("uciok\n");

// 	 while (TRUE) {
// 		 memset(&line[0], 0, sizeof(line));
// 		 fflush(stdout);
// 		 if (!fgets(line, INPUTBUFFER, stdin))
// 			 continue;

// 		 if (line[0] == '\n')
// 			 continue;

// 		 if (!strncmp(line, "isready", 7)) {
// 			 printf("readyok\n");
// 			 continue;
// 		 } else if (!strncmp(line, "position", 8)) {
// 			 ParsePosition(line, pos);
// 		 } else if (!strncmp(line, "ucinewgame", 10)) {
// 			 ParsePosition("position startpos\n", pos);
// 		 } else if (!strncmp(line, "go", 2)) {
// 			 ParseGo(line, info, pos);
// 		 } else if (!strncmp(line, "quit", 4)) {
// 			 info.quit := TRUE;
// 			 break;
// 		 } else if (!strncmp(line, "uci", 3)) {
// 			 printf("id name %s\n", NAME);
// 			 printf("id author Bluefever\n");
// 			 printf("uciok\n");
// 		 } else if(!strncmp(line, "setoption name Hash value ", 26)) {
// 			 sscanf(line, "%*s %*s %*s %*s %d", &MB);
// 			 if(MB < 4) MB := 4;
// 			 if(MB > MAX_HASH) MB := MAX_HASH;
// 			 printf("Set Hash to %d MB\n", MB);
// 			 InitHashTable(pos.HashTable, MB);
// 		 } else if(!strncmp(line, "setoption name Book value ", 26)) {
// 			 ptrTrue := strstr(line, "true");
// 			 if(ptrTrue != NULL) {
// 				 EngineOptions.UseBook := TRUE;
// 			 } else {
// 				 EngineOptions.UseBook := false;
// 			 }
// 		 }
// 		 if(info.quit) break;
// 	 }
//  }

// type UCI struct {
// 	engine *engine.Engine
// 	game   *chess.Game
// 	debug  bool

// 	// Callback to send signal to child functions and cancel execution.
// 	cancel context.CancelFunc
// }

// // Reads UCI commands from stdin and outputs to stdout.
// func (uci *UCI) Loop() {
// 	uci.engine = engine.New()
// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Printf("id name %s\n", NAME)
// 	fmt.Printf("id author Bluefever\n")
// 	fmt.Printf("option name Hash type spin default 64 min 4 max %d\n", MAX_HASH)
// 	fmt.Printf("option name Book type check default true\n")
// 	fmt.Printf("uciok\n")
// 	fmt.Printf("option name Hash type spin default 64 min 4 max %d\n", MAX_HASH)
// 	fmt.Printf("option name Book type check default true\n")

// 	for scanner.Scan() {
// 		uci.printDebug(fmt.Sprintf("received %q", scanner.Text()))

// 		line := strings.TrimSpace(scanner.Text())
// 		if len(line) == 0 {
// 			continue
// 		}

// 		if line == "quit" {
// 			return
// 		}

// 		if err := uci.handleCommand(line); err != nil {
// 			logger.Printf("uci.handleCommand(%q) error: %v", line, err)
// 		}
// 	}

// 	if scanner.Err() != nil {
// 		logger.Printf("scanner.Scan() error: %v", scanner.Err())
// 	}
// }

// func (uci *UCI) handleCommand(line string) error {
// 	fields := strings.Fields(line)
// 	cmd, params := fields[0], fields[1:]

// 	if cmd == "uci" {
// 		reply("id name ggchess")
// 		reply("id author mrboxtobox")
// 		reply("uciok")
// 	} else if cmd == "isready" {
// 		reply("readyok")
// 	} else if cmd == "setoption" {
// 		uci.handleOption(params)
// 	} else if cmd == "ucinewgame" {
// 		uci.handleNewGame()
// 	} else if cmd == "position" {
// 		uci.handlePosition(params)
// 	} else if cmd == "go" {
// 		uci.handleGo(params)
// 	} else if cmd == "stop" {
// 		uci.handleStop()
// 	} else if cmd == "ponderhit" {
// 		uci.handlePonderHit()
// 	} else if cmd == "debug" {
// 		uci.handleDebug(params)
// 	} else {
// 		logger.Printf("cannot parse command; cmd: %q, params: %v", cmd, params)
// 	}

// 	return nil
// }

// func (uci *UCI) handleOption(params []string) {
// 	// TODO: Implement this.
// 	uci.engine.SetOption(params)
// }

// func (uci *UCI) handleNewGame() {
// 	uci.game = chess.NewGame(un)
// }

// // The subcommands will either be 'startpos' or 'fen'. We only need to configure
// // the engine for the 'fen' case.
// func (uci *UCI) handlePosition(params []string) {
// 	subcommand := params[0]

// 	idx := find(params, "moves")
// 	if subcommand == "startpos" {
// 		uci.game = chess.NewGame(un)
// 	} else if subcommand == "fen" {
// 		var err error
// 		var fen func(*chess.Game)
// 		if idx == -1 {
// 			str := strings.Join(params[1:], " ")
// 			if fen, err = chess.FEN(str); err != nil {
// 				logger.Fatal(err)
// 			}
// 		} else {
// 			str := strings.Join(params[1:idx], " ")
// 			if fen, err = chess.FEN(str); err != nil {
// 				logger.Fatal(err)
// 			}
// 		}

// 		uci.game = chess.NewGame(un, fen)
// 	} else {
// 		logger.Fatalf("unable to parse position params %v", params)
// 	}

// 	// Apply the subsequent moves if provided.
// 	if idx == -1 {
// 		return
// 	}
// 	for _, m := range params[idx+1:] {
// 		if err := uci.game.MoveStr(m); err != nil {
// 			logger.Fatalf("game.MoveStr(%v) error: %v", m, err)
// 		}
// 	}
// }

// func (uci *UCI) handleGo(params []string) {
// 	sp := &engine.SearchOptions{
// 		Game:       uci.game.Clone(),
// 		Encoder:    chess.UCINotation{},
// 		Limits:     parseLimits(params),
// 		OnProgress: replyInfo,
// 	}

// 	ctx, cancel := context.WithCancel(context.Background())
// 	uci.cancel = cancel

// 	// Run search in a separate goroutine.
// 	// TODO: Figure out if this otherwise blocks IO.
// 	go func() {
// 		defer uci.cancel()
// 		si := uci.engine.Search(ctx, sp)

// 		if len(si.MainLine) > 1 {
// 			reply(fmt.Sprintf("bestmove %v ponder %v", si.MainLine[0], si.MainLine[1]))
// 		} else if len(si.MainLine) > 0 {
// 			reply(fmt.Sprintf("bestmove %v", si.MainLine[0]))
// 		} else {
// 			logger.Printf("Could not find mainline")
// 			reply("bestmove 0000")
// 		}
// 	}()
// }

// func (uci *UCI) handleStop() {
// 	if uci.cancel != nil {
// 		uci.cancel()
// 	}
// }

// func (uci *UCI) handlePonderHit() {
// 	// TODO: Set this at the UCI level.
// 	uci.engine.StopPondering()
// }

// func (uci *UCI) handleDebug(params []string) {
// 	if params[0] == "on" {
// 		uci.debug = true
// 	} else if params[0] == "off" {
// 		uci.debug = false
// 	}
// }

// func (uci *UCI) printDebug(str string) {
// 	if uci.debug {
// 		reply(fmt.Sprintf("info string %v\n", str))
// 	}
// }

// func replyInfo(si *engine.SearchInfo) {
// 	sb := &strings.Builder{}
// 	fmt.Fprintf(sb, "info depth %d", si.Depth)
// 	if si.MateIn != 0 {
// 		fmt.Fprintf(sb, " score mate %d", si.MateIn)
// 	} else {
// 		fmt.Fprintf(sb, " score cp %d", si.CentipawnScore)
// 	}

// 	t := si.Duration.Milliseconds()
// 	nps := si.Nodes / int(si.Duration.Seconds()+1)
// 	fmt.Fprintf(sb, " nodes %d time %d nps %d", si.Nodes, t, nps)

// 	if len(si.MainLine) > 0 {
// 		pv := strings.Join(si.MainLine, " ")
// 		fmt.Fprintf(sb, " pv %v", pv)
// 	}

// 	reply(sb.String())
// }

// func parseLimits(params []string) *engine.SearchLimits {
// 	limits := &engine.SearchLimits{}
// 	for i := 0; i < len(params); i++ {
// 		switch params[i] {
// 		case "searchmoves":
// 			var moves []string
// 			i++
// 			for ; i < len(params); i++ {
// 				// Keep searching until we find another token.
// 				if find(goParams, params[i]) != -1 {
// 					break
// 				}
// 				moves = append(moves, params[i])
// 			}
// 			limits.SearchMoves = moves
// 		case "ponder":
// 			limits.Ponder = true
// 		case "wtime":
// 			i++
// 			limits.WhiteTimeMs = parseInt(params[i])
// 		case "btime":
// 			i++
// 			limits.BlackTimeMs = parseInt(params[i])
// 		case "winc":
// 			i++
// 			limits.WhiteIncrementMs = parseInt(params[i])
// 		case "binc":
// 			i++
// 			limits.BlackIncrementMs = parseInt(params[i])
// 		case "movestogo":
// 			i++
// 			limits.MovesToGo = parseInt(params[i])
// 		case "depth":
// 			i++
// 			limits.Depth = parseInt(params[i])
// 		case "nodes":
// 			i++
// 			limits.Nodes = parseInt(params[i])
// 		case "mate":
// 			i++
// 			limits.MateIn = parseInt(params[i])
// 		case "movetime":
// 			i++
// 			limits.MoveTimeMs = parseInt(params[i])
// 		case "infinite":
// 			limits.Infinite = true
// 		}
// 	}
// 	return limits
// }

// func parseInt(val string) int {
// 	res, err := strconv.Atoi(val)
// 	if err != nil {
// 		logger.Fatalf("strconv.Atoi(%v) error: %v", val, err)
// 		return -1
// 	}
// 	return res
// }

// func find(list []string, s string) int {
// 	for i, p := range list {
// 		if strings.EqualFold(p, s) {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func reply(msg string) {
// 	fmt.Println(msg)
// }
