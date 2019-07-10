package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// CountWords represents the file values
type CountWords struct {
	wordCount      int
	lineCount      int
	characterCount int
}

func countWordsInFile(s string) int {
	words := strings.Fields(s)
	return len(words)
}

func main() {

	/*
		var lines bool
		var words bool
		var chars bool
		flag.BoolVar(&lines, "l", false, "")
		flag.BoolVar(&words, "w", false, "")
		flag.BoolVar(&chars, "c", false, "")
		flag.Parse()

		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error retrieving directory", err)
			return
		}
	*/

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	wc := countWordsInFile(string(data))

	counts := CountWords{
		wordCount:      wc,
		lineCount:      1,
		characterCount: 10,
	}

	fmt.Println(counts.wordCount, counts.lineCount, counts.characterCount)
}
