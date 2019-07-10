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

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	flag.Parse()

	switch flag.NArg() {
	case 0:
		data, err := ioutil.ReadAll(os.Stdin)
		errorCheck(err)
		wc := countWordsInFile(string(data))
		fmt.Println(wc)
	case 1:
		data, err := ioutil.ReadFile(flag.Arg(0))
		errorCheck(err)
		wc := countWordsInFile(string(data))
		fmt.Println(wc)
	default:
		fmt.Println("No input given")
		os.Exit(1)

	}
}
