package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func countWordsInFile(s string) string {
	words := strings.Fields(s)
	l := len(words)
	o := strconv.Itoa(l)
	return o
}

func countLinesInFile(s string) string {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	o := strconv.Itoa(count)
	return o
}

func countCharsInFile(s string) string {
	l := len(s)
	o := strconv.Itoa(l)
	return o
}

func main() {

	words := flag.Bool("w", false, "")
	lines := flag.Bool("l", false, "")
	chars := flag.Bool("c", false, "")
	flag.Parse()

	if flag.NArg() == 0 {
		file := os.Stdin
		fi, err := file.Stat()
		if err != nil {
			fmt.Println("There was an error retrieving file description:", err)
		}

		size := fi.Size()
		if size > 0 {
			data, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println("There was an error reading from stdin:", err)
			}

			var printValues string
			s := string(data)

			if *words {
				printValues += (countWordsInFile(s) + " ")
			}

			if *lines {
				printValues += (countLinesInFile(s) + " ")
			}

			if *chars {
				printValues += (countCharsInFile(s) + " ")
			}

			if !*words && !*lines && !*chars {
				fmt.Println(countWordsInFile(s), countLinesInFile(s), countCharsInFile(s))
			}

			fmt.Println(printValues)

		} else {
			fmt.Println("Stdin is empty")
		}
	} else if flag.NArg() == 1 {
		data, err := ioutil.ReadFile(flag.Arg(0))
		if err != nil {
			fmt.Println("There was an error reading the file:", err)
		}

		var printValues string
		s := string(data)

		if *words {
			printValues += (countWordsInFile(s) + " ")
		}

		if *lines {
			printValues += (countLinesInFile(s) + " ")
		}

		if *chars {
			printValues += (countCharsInFile(s) + " ")
		}

		if !*words && !*lines && !*chars {
			fmt.Println(countWordsInFile(s), countLinesInFile(s), countCharsInFile(s))
		}

		fmt.Println(printValues)
	}
}
