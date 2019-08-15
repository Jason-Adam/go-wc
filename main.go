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

// wcString prints to stdout based on flags.
func (f *flags) wcString(s string) string {
	var printValues string
	if f.words {
		printValues += (wordsInFile(s) + " ")
	}
	if f.lines {
		printValues += (linesInFile(s) + " ")
	}
	if f.chars {
		printValues += (charsInFile(s) + " ")
	}
	if !f.words && !f.lines && !f.chars {
		return wordsInFile(s) + " " + linesInFile(s) + " " + charsInFile(s)
	}
	return printValues
}

// wordsInFile counts total words in the file.
func wordsInFile(s string) string {
	words := strings.Fields(s)
	l := len(words)
	return strconv.Itoa(l)
}

// linesInFile counts the total lines in the file.
func linesInFile(s string) string {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	return strconv.Itoa(count)
}

// charsInFile returns the character count.
func charsInFile(s string) string {
	l := len(s)
	return strconv.Itoa(l)
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
