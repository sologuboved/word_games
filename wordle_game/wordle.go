package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"strings"
	"time"
	// "/Users/sologuboved/Golang/src/word_games/misc"
)

func getWord (maxlen int) string {
	f, _ := os.Open(getSrcFname("english.txt"))
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
	incl := make([]string, 0)
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
	maxlen := 5
	mysteryWord := getWord(maxlen)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Any ideas?")
	for input.Scan() {
		candidate := input.Text()
		if candidate == "fin" {
			fmt.Printf("\nMystery word was %s\n", mysteryWord)
			break
		}
		masked, incl, excl := getVerdict(candidate, mysteryWord)
		fmt.Println()
		fmt.Println(masked)
		fmt.Printf("Wrongly positioned: %v\n", incl)
		fmt.Printf("Not there: %v\n", excl)
		if !strings.Contains(masked, "_") {
			fmt.Printf("Right, it was %s\n", mysteryWord)
			break
		}
		fmt.Println("\nNext idea?")
	}
}
