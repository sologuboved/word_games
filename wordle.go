package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"strings"
	"time"
)

func getWord (maxlen int) string {
	f, _ := os.Open("english.txt")
	input := bufio.NewScanner(f)
	words := make([]string, 0)
	for input.Scan() {
		word := input.Text()
		if len(word) == maxlen {
			words = append(words, word)
		}
	}
	return words[rand.Intn(len(words))]
}

func getVerdict(candidate, word string) (string, []string, []string) {
	len_word := len(word)
	masked := make([]string, len_word)
	for i := range masked {
		masked[i] = "_"
	}
	incl := make([]string, len_word)
	excl := make([]string, 0)
	if len(candidate) != len_word {
		fmt.Println("Wrong length")
	} else {
		for i := 0; i < len_word; i++ {
			letter := candidate[i]
			if word[i] == letter {
				masked[i] = string(letter)
				continue
			}
			if strings.Contains(word, string(letter)) {
				incl = append(incl, string(letter))
			} else {
				excl = append(excl, string(letter))
			}
		}
	}
	return strings.Join(masked, ""), incl, excl
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// fmt.Println(getVerdict("skate", "place"))
	fmt.Println(getWord(5))
}
