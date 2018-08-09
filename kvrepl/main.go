package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Exercise: k/v REPL
// Author: Vladimir Vivien (@vladimirvivien)
// 02/26/18

// This is a simple implementation of a key/value pair REPL with
// in memory storage.  The implementation uses a variable store
// for storage and txLog for transaction log.  When a transaction
// starts, using txOn, all k/v operations are done using txLog until
// the transaction is committed, then the data is copy into store.

var (
	promt = "kv> "
	store = make(map[string]string)
	txLog = make(map[string]string)
	txOn  bool
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(promt)
		cmdLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input invalid, type help")
			continue
		}

		cmd, p0, p1 := parseCmd(cmdLine)

		if cmd == "" {
			continue
		}

		switch cmd {
		case "START":
			txOn = true
		case "COMMIT":
			if !txOn {
				fmt.Println("no transaction started")
			}
			for k, v := range txLog {
				store[k] = v
			}
			txOn = false
			txLog = make(map[string]string)
		case "ABORT":
			if !txOn {
				fmt.Println("no transaction started")
			}
			txOn = false
			txLog = make(map[string]string)
		case "READ":
			if p0 == "" {
				fmt.Println("missing argument, usage: READ <key>")
				continue
			}

			var val string
			var ok bool

			if txOn {
				val, ok = txLog[p0]
			} else {
				val, ok = store[p0]
			}

			if ok {
				fmt.Println(val)
			} else {
				fmt.Println("key not found:", p0)
			}

		case "WRITE":
			if p0 == "" || p1 == "" {
				fmt.Println("missing arguments, usage: WRITE <key> <value>")
				continue
			}
			if txOn {
				txLog[p0] = p1
			} else {
				store[p0] = p1
			}
		case "DELETE":
			if p0 == "" {
				fmt.Println("missing argument, usage: DELETE <key>")
				continue
			}
			if txOn {
				delete(txLog, p0)
			} else {
				delete(store, p0)
			}
		case "QUIT":
			goto quit
		default:
			fmt.Println("invalid command")
			continue
		}
	}
quit:
	fmt.Println("good bye")
}

func parseCmd(cmdLine string) (cmd, p0, p1 string) {
	parts := strings.Split(cmdLine, " ")
	if len(parts) < 1 {
		return "", "", ""
	}
	cmd = strings.ToUpper(strings.TrimSpace(parts[0]))

	if len(parts) == 1 {
		return cmd, "", ""
	}
	if len(parts) == 2 {
		p0 = strings.TrimSpace(parts[1])
	}
	if len(parts) > 2 {
		p0 = strings.TrimSpace(parts[1])
		p1 = strings.TrimSpace(parts[2])
	}
	return
}
