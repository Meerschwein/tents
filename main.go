package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Meerschwein/tents/pkg/asp"
	"github.com/Meerschwein/tents/pkg/asp/solution"
	"github.com/Meerschwein/tents/pkg/clingo"
	"github.com/Meerschwein/tents/pkg/tents"
	"github.com/alexflint/go-arg"
)

type Args struct {
	InFormat  string `arg:"-f" default:"puzzle"  help:"puzzle | asp"`
	OutFormat string `arg:"-o" default:"puzzle"  help:"puzzle | asp"`
	Solution  string `arg:"-s" default:"choices" help:"choices | disjunction | negation"`
	File      string `arg:"positional"           help:"stdin if not given"`
}

func init() {
	args := Args{}

	{
		parser := arg.MustParse(&args)

		switch args.InFormat {
		case "puzzle", "asp":
		default:
			fmt.Println("Invalid informat:", args.InFormat)
			parser.WriteHelp(os.Stdout)
			os.Exit(1)
		}

		switch args.OutFormat {
		case "puzzle", "asp":
		default:
			fmt.Println("Invalid outformat:", args.OutFormat)
			parser.WriteHelp(os.Stdout)
			os.Exit(1)
		}

		switch args.Solution {
		case "choices", "disjunction", "negation":
		default:
			fmt.Println("Invalid solution:", args.Solution)
			parser.WriteHelp(os.Stdout)
			os.Exit(1)
		}
	}

	var puzzleData []byte
	var err error

	switch args.File {
	case "":
		puzzleData, err = io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	default:
		f, err := os.Open(args.File)
		if err != nil {
			panic(err)
		}
		puzzleData, err = io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		f.Close()
	}

	switch args.InFormat {
	case "puzzle":
		puzzle, err = tents.ParsePuzzle(string(puzzleData))
		if err != nil {
			panic(err)
		}
	case "asp":
		preds, err := asp.ParsePredicates(strings.Split(string(puzzleData), "\n"))
		if err != nil {
			panic(err)
		}
		puzzle, err = tents.ParseAsp(preds)
		if err != nil {
			panic(err)
		}
	}

	program = solution.Solutions[args.Solution]
	outformat = args.OutFormat
}

var (
	puzzle    tents.Puzzle
	outformat string
	program   string
)

func main() {
	for _, p := range puzzle.ToAsp() {
		program += p.String() + "\n"
	}

	cr, err := clingo.Run(strings.NewReader(program))
	if err != nil {
		panic(err)
	}

	if !cr.GoodExitCode() {
		fmt.Println(cr.ExitCode)
		os.Exit(1)
	}

	puzzle, err = tents.ParseAsp(cr.Predicates)
	if err != nil {
		panic(err)
	}

	switch outformat {
	case "puzzle":
		fmt.Println(puzzle.ToPuzzle())
	case "asp":
		for _, p := range puzzle.ToAsp() {
			fmt.Println(p.String())
		}
	}
}
