package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Meerschwein/tents/pkg/tents"
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	puzzleFile := flag.Arg(0)

	puzzleFileContent, err := os.ReadFile(puzzleFile)
	if err != nil {
		panic(err)
	}

	puzzle, err := tents.Parse(string(puzzleFileContent))
	if err != nil {
		panic(err)
	}

	aspProgram := puzzle.ToAspProgram()

	for _, fact := range aspProgram {
		fmt.Println(fact)
	}
}
