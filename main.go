package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Meerschwein/tents/pkg/asp"
	"github.com/Meerschwein/tents/pkg/asp/solution"
	"github.com/Meerschwein/tents/pkg/clingo"
	"github.com/Meerschwein/tents/pkg/tents"
	"golang.org/x/tools/txtar"
)

var (
	mode         = flag.String("mode", "puzzle", "format of the input file (puzzle | asp)\npuzzle: tents puzzle\nasp: asp encoding of a tents puzzle")
	solutionType = flag.String("solution", "choices", "solution type (choices | disjunction)")
	puzzleFrom   io.Reader
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] <file>\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "If no file is given the input is read from stdin.")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *mode != "puzzle" && *mode != "asp" && *mode != "txtar" {
		flag.Usage()
		os.Exit(1)
	}

	if *solutionType != "choices" && *solutionType != "disjunction" {
		flag.Usage()
		os.Exit(1)
	}

	switch flag.NArg() {
	case 0:
		puzzleFrom = os.Stdin
	case 1:
		var err error
		puzzleFrom, err = os.Open(flag.Arg(0))
		if err != nil {
			panic(err)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	puzzleData, err := io.ReadAll(puzzleFrom)
	if err != nil {
		panic(err)
	}
	switch puzzleFrom := puzzleFrom.(type) {
	case io.Closer:
		puzzleFrom.Close()
	}

	var p tents.Puzzle
	if *mode == "puzzle" {
		p, err = tents.ParsePuzzle(string(puzzleData))
	} else if *mode == "asp" {
		preds, err2 := asp.ParsePredicates(strings.Split(string(puzzleData), "\n"))
		if err2 != nil {
			panic(err)
		}
		p, err = tents.ParseAsp(preds)
	} else if *mode == "txtar" {
		ar := txtar.Parse(puzzleData)
		for _, f := range ar.Files {
			if f.Name == "puzzle" {
				p, err = tents.ParsePuzzle(string(f.Data))
				goto found
			}
		}
		panic("no puzzle found in txtar")
	found:
	}
	if err != nil {
		panic(err)
	}

	all := solution.Solutions[*solutionType]
	for _, p := range p.ToAsp() {
		println(p.String())
		all += p.String() + "\n"
	}

	cr, err := clingo.Run(strings.NewReader(all))
	if err != nil {
		panic(err)
	}

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
