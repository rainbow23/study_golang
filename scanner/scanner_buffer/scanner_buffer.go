package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := "abcdefghijkl" + strings.Repeat("X", 65536)

	cnt := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		fmt.Printf("%t\t%d\t%s\n", atEOF, len(data), data)
		cnt += 1
		return 0, nil, nil
	}
	scanner.Split(split)
	buf := make([]byte, 10000)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	fmt.Println("cnt ", cnt)
}
