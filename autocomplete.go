package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ryancaille/trie"
)

func main() {
	filename := os.Args[1]
	count, _ := strconv.Atoi(os.Args[2])
	words := parseWordFile(filename)

	t := trie.NewTrie()

	start := time.Now()
	for _, w := range words {
		t.Insert(w)
	}
	elapsed := time.Since(start)
	log.Printf("Load took %s to load %d words", elapsed, t.Count())

	for {
		var input string
		fmt.Scan(&input)
		suggestions := t.Like(input, count)

		fmt.Println(strings.Join(suggestions, " "))
	}

}

func parseWordFile(file string) []string {
	words := make([]string, 0)
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}
