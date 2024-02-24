package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/Meerschwein/tents/pkg/tents"
)

//go:embed tents.asp
var tentsAsp string

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

	all := tentsAsp + strings.Join(p.ToAspProgram(), "\n")

	println(string(puzzleFileContent))

	// call clingo
	c := exec.Command("clingo", "--outf=1", "-t", fmt.Sprint(runtime.NumCPU()))
	c.Stdin = strings.NewReader(all)

	out, err := c.CombinedOutput()
	if err != nil &&
		c.ProcessState.ExitCode() != 30 &&
		c.ProcessState.ExitCode() != 10 {
		println(string(out))
		panic(err)
	}
	println(string(out))

	lines := strings.Split(string(out), "\n")
	// remove all lines with the % prefix
	var result []string
	for _, line := range lines {
		if !strings.HasPrefix(line, "%") && line != "" && line != "ANSWER" {
			result = append(result, line)
		}
	}
	if len(result) != 1 {
		println("i dont know what to do")
		os.Exit(1)
	}

	p, err = tents.ParseAsp(strings.Join(strings.Split(result[0], " "), "\n"))
	if err != nil {
		panic(err)
	}

	println(p.String())
}
