package main

import (
	"bufio"
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

func countLinesInFile(s string) int {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func countCharsInFile(s string) int {
	return len(s)
}

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	words := flag.Bool("w", false, "")
	lines := flag.Bool("l", false, "")
	chars := flag.Bool("c", false, "")
	flag.Parse()

	switch flag.NArg() {
	case 0:
		data, err := ioutil.ReadAll(os.Stdin)
		errorCheck(err)

		if *words {
			fmt.Println(countWordsInFile(string(data)))
		}

		if *lines {
			fmt.Println(countLinesInFile(string(data)))
		}

		if *chars {
			fmt.Println(countCharsInFile(string(data)))
		}

		if !*words && !*lines && !*chars {
			fmt.Println(countWordsInFile(string(data)), countLinesInFile(string(data)), countCharsInFile(string(data)))
		}

	case 1:
		data, err := ioutil.ReadFile(flag.Arg(0))
		errorCheck(err)

		if *words {
			fmt.Println(countWordsInFile(string(data)))
		}

		if *lines {
			fmt.Println(countLinesInFile(string(data)))
		}

		if *chars {
			fmt.Println(countCharsInFile(string(data)))
		}

		if !*words && !*lines && !*chars {
			fmt.Println(countWordsInFile(string(data)), countLinesInFile(string(data)), countCharsInFile(string(data)))
		}

	default:
		fmt.Println("No input given")
		os.Exit(1)
	}
}
