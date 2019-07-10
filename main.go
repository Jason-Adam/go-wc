package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func countWordsInFile(s string) int {
	words := strings.Fields(s)
	return len(words)
}

// Place holder until I can figure it out
func countLinesInFile(s int) int {
	return s
}

// Place holder until I can figure it out
func countCharsInFile(s int) int {
	return s
}

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func retrieveCounts(d string) []int {
	s := make([]int, 3)
	s[0] = countWordsInFile(d)
	s[1] = countLinesInFile(10) // I hardcoded these two until I can figure out how to do them
	s[2] = countCharsInFile(100)
	return s
}

func main() {

	words := flag.Bool("w", false, "") // Only added "w" for now until I figure out the other functions
	flag.Parse()

	switch flag.NArg() {
	case 0:
		data, err := ioutil.ReadAll(os.Stdin)
		errorCheck(err)

		fileCounts := retrieveCounts(string(data))

		if *words {
			fmt.Println(fileCounts[0])
		} else {
			fmt.Println(fileCounts)
		}

	case 1:
		data, err := ioutil.ReadFile(flag.Arg(0))
		errorCheck(err)

		fileCounts := retrieveCounts(string(data))

		if *words {
			fmt.Println(fileCounts[0], flag.Arg(0))
		} else {
			fmt.Println(fileCounts, flag.Arg(0))
		}

	default:
		fmt.Println("No input given")
		os.Exit(1)

	}
}
