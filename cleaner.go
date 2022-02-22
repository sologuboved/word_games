package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func clean (fname string) {
	output := make([]string, 0)
	f, _ := os.Open(fname)
	input := bufio.NewScanner(f)
	for input.Scan() {
		isSuitable := true
		word := input.Text()
		for _, char := range(word) {
			if unicode.IsLetter(char) {
				if unicode.IsUpper(char) {
					isSuitable = false
					break
				}
			} else {
				isSuitable = false
				break
			}
		}
		if isSuitable {
			output = append(output, word)
		}
	}
	fmt.Printf("Found %d suitable words\n", len(output))
	f, _ = os.Create(fname)
	defer f.Close()
	f.WriteString(strings.Join(output, "\n"))
}
