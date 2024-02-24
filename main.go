package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Meerschwein/tents/pkg/asp"
	"github.com/Meerschwein/tents/pkg/clingo"
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

	p, err := tents.ParsePuzzle(string(puzzleFileContent))
	if err != nil {
		panic(err)
	}

	println(string(puzzleFileContent) + "\n")

	all := asp.TentsSolution
	for _, p := range p.ToAsp() {
		all += p.String() + "\n"
	}

	cr, err := clingo.Run(strings.NewReader(all))
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%+v", cr)

	if !cr.GoodExitCode() {
		fmt.Println(cr.ExitCode)
		return
	}

	p, err = tents.ParseAsp(cr.Predicates)
	if err != nil {
		panic(err)
	}

	println(p.ToPuzzle())
}
