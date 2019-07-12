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

type flags struct {
	words bool
	lines bool
	chars bool
}

func (f *flags) wcString(s string) string {
	var printValues string

	if f.words {
		printValues += (countWordsInFile(s) + " ")
	}

	if f.lines {
		printValues += (countLinesInFile(s) + " ")
	}

	if f.chars {
		printValues += (countCharsInFile(s) + " ")
	}

	if !f.words && !f.lines && !f.chars {
		return countWordsInFile(s) + " " + countLinesInFile(s) + " " + countCharsInFile(s)
	}

	return printValues
}

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
	var w, l, c bool
	flag.BoolVar(&w, "w", false, "")
	flag.BoolVar(&l, "l", false, "")
	flag.BoolVar(&c, "c", false, "")
	flag.Parse()

	f := flags{words: w, lines: l, chars: c}

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

			s := string(data)
			fmt.Println(f.wcString(s))
			return

		}
		fmt.Println("Stdin is empty")
		return
	} else if flag.NArg() == 1 {
		data, err := ioutil.ReadFile(flag.Arg(0))
		if err != nil {
			fmt.Println("There was an error reading the file:", err)
		}

		s := string(data)
		fmt.Println(f.wcString(s))
		return
	}
	fmt.Println("Too many inputs")
	return
}
