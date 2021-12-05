package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var nFlag uint
var debugFlag bool

func init() {
	flag.UintVar(&nFlag, "n", 5, "number of lines")
	flag.BoolVar(&debugFlag, "debug", false, "write debug info to stderr")
	flag.Parse()
	if debugFlag {
		fmt.Fprintf(os.Stderr, "n=%d\n", nFlag)
	}
	if nFlag == 0 {
		os.Exit(0)
	}
}

func main() {
	var err error

	headLines := make([]string, 0, nFlag)
	tailLines := make([]*string, 0, nFlag)
	tailStart := 0

	lineNum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for ; scanner.Scan(); lineNum++ {
		err = scanner.Err()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error scanning: %v", err)
			os.Exit(1)
		}
		line := scanner.Text()

		if uint(lineNum+1) <= nFlag {
			headLines = append(headLines, line)
		}

		if uint(len(tailLines)) < nFlag {
			tailLines = append(tailLines, &line)
		} else {
			tailLines[tailStart] = &line
			tailStart = (tailStart + 1) % len(tailLines)
		}
	}

	// Realign tailLines if necessary.
	for i := 0; i < tailStart; i++ {
		rotateLeft(tailLines)
	}

	// If regions are mutually exclusive, just print each set.
	if uint(lineNum) >= 2*nFlag {
		if debugFlag {
			fmt.Fprintf(os.Stderr, "mutually exclusive regions\n")
		}
		for _, line := range headLines {
			fmt.Printf("%s\n", line)
		}
		for _, line := range tailLines {
			fmt.Printf("%s\n", *line)
		}
		// There is overlap.  See if everything gathered into the head slice.
	} else if uint(lineNum) <= nFlag {
		if debugFlag {
			fmt.Fprintf(os.Stderr, "everything fit into head slice\n")
		}
		for _, line := range headLines {
			fmt.Printf("%s\n", line)
		}
	} else {
		// Otherwise, we end up printing all of the lines;
		// all from the head slice and the remainder from tail slice.
		leftoverCnt := lineNum - len(headLines)
		tailStart := len(tailLines) - leftoverCnt
		if debugFlag {
			fmt.Fprintf(os.Stderr, "basically cat with extra steps\n")
			fmt.Fprintf(os.Stderr, "n=%d, maxlines=%d\n", nFlag, lineNum)
			fmt.Fprintf(os.Stderr, "headlines=%d\n", len(headLines))
			fmt.Fprintf(os.Stderr, "taillines=%d, leftovercnt=%d, tailstart=%d\n", len(tailLines), leftoverCnt, tailStart)
		}
		for _, line := range headLines {
			fmt.Printf("%s\n", line)
		}
		for _, line := range tailLines[tailStart:] {
			fmt.Printf("%s\n", *line)
		}
	}
}

func rotateLeft(slice []*string) {
	if len(slice) == 0 {
		return
	}
	lead := slice[0]
	for i := 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	slice[len(slice)-1] = lead
}
