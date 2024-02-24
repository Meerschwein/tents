package clingo

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"

	"github.com/Meerschwein/tents/pkg/asp"
	"github.com/Meerschwein/tents/pkg/util"
)

// https://github.com/potassco/clasp/issues/42
// https://www.mat.unical.it/aspcomp2013/files/aspoutput.txt
type ExitCode int

// this is a best guess at what the exit codes are
const (
	Unknown                ExitCode = 0
	Interrupted            ExitCode = 1
	QueryIsTrue            ExitCode = 10
	QueryIsTrueInterrupted ExitCode = 11
	QueryIsFalse           ExitCode = 20
	OptimaFound            ExitCode = 30
	OptimaFoundInterrupted ExitCode = 31
	OutOfMemory            ExitCode = 33
	Error                  ExitCode = 65
	SyntaxError            ExitCode = 128
)

func (e ExitCode) Error() string {
	switch e {
	case Unknown:
		return "satisfiablity of problem not known and search not started"
	case Interrupted:
		return "run was interrupted"
	case QueryIsTrue:
		return "at least one model was found"
	case QueryIsTrueInterrupted:
		return "at least one model was found and clingo was interrupted"
	case QueryIsFalse:
		return "program is inconsistent, no model found"
	case OptimaFound:
		return "program is consistent and all models have been enumerated"
	case OptimaFoundInterrupted:
		return "program is consistent and all models have been enumerated and clingo was interrupted"
	case OutOfMemory:
		return "out of memory"
	case Error:
		return "run was interrupted by internal error"
	case SyntaxError:
		return "search not started because of syntax or command line error"
	default:
		return "Unknown exit code"
	}
}

func (cr ClingoResult) GoodExitCode() bool {
	return cr.ExitCode == QueryIsTrue || cr.ExitCode == OptimaFound
}

type ClingoResult struct {
	ExitCode   ExitCode
	Delimiter  string
	Predicates []asp.Predicate
	Out        bytes.Buffer
	Err        bytes.Buffer
}

func Run(stdin io.Reader) (ClingoResult, error) {
	c := exec.Command("clingo", "--outf=1", "-t", fmt.Sprint(runtime.NumCPU()))

	cr := ClingoResult{}

	c.Stdin = stdin
	c.Stderr = &cr.Err
	c.Stdout = &cr.Out

	c.Run()

	cr.ExitCode = ExitCode(c.ProcessState.ExitCode())

	if !cr.GoodExitCode() {
		return cr, nil
	}

	out := []string{}
	for _, line := range strings.Split(cr.Out.String(), "\n") {
		if util.IsBlank(line) || strings.HasPrefix(line, "%") {
			continue
		}

		switch line {
		case "ANSWER", "COST", "INCONSISTENT", "UNKNOWN", "OPTIMUM":
			cr.Delimiter = line
			continue
		}
		out = append(out, line)
	}
	if len(out) == 0 {
		return ClingoResult{}, fmt.Errorf("no output")
	}

	preds, err := asp.ParsePredicates(strings.Split(out[0], " "))
	if err != nil {
		return ClingoResult{}, err
	}

	cr.Predicates = preds

	return cr, nil
}
