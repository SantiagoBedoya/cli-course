package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")

	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	switch {
	case !countLines && !countBytes:
		scanner.Split(bufio.ScanWords)
	case countBytes:
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
