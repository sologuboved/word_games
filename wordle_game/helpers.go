package main

import (
	"bufio"
	"fmt"
	"path/filepath"
	"os"
	"strings"
	"unicode"
)

func getSrcFname(fname string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	return parent + "/words/" + fname 
}

func clean (fname string) {
	fname = getSrcFname(fname)
	output := make([]string, 0)
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
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
	f, err = os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(strings.Join(output, "\n"))
}