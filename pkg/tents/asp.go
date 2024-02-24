package tents

import (
	"fmt"
	"strings"
)

func appendFact(facts []string, name string, args ...int) []string {
	argsStrs := []string{}
	for _, arg := range args {
		argsStrs = append(argsStrs, fmt.Sprint(arg))
	}
	fact := fmt.Sprintf("%s(%s).", name, strings.Join(argsStrs, ","))
	return append(facts, fact)
}

func (p Puzzle) ToAspProgram() (facts []string) {
	facts = appendFact(facts, "lines", p.Rows)
	facts = appendFact(facts, "columns", p.Columns)

	for i, row := range p.Board {
		for j, cell := range row {
			name := ""
			switch cell {
			case Empty:
				name = "free"
			case Tree:
				name = "tree"
			case Tent:
				name = "tent"
			}
			facts = appendFact(facts, name, i+1, j+1)
		}
	}

	for i, count := range p.RowTentCount {
		facts = appendFact(facts, "line", i+1)
		facts = appendFact(facts, "rowsum", i+1, count)
	}

	for i, count := range p.ColumnTentCount {
		facts = appendFact(facts, "column", i+1)
		facts = appendFact(facts, "colsum", i+1, count)
	}

	return
}
