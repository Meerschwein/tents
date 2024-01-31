package tents

import "fmt"

type Cell int

const (
	Empty Cell = iota
	Tree
	Tent
)

type Puzzle struct {
	Rows            int
	Columns         int
	Board           [][]Cell
	RowTentCount    []int
	ColumnTentCount []int
}

func (c Cell) String() string {
	switch c {
	case Empty:
		return "."
	case Tree:
		return "T"
	case Tent:
		return "^"
	}

	panic("invalid cell")
}

func (p Puzzle) String() string {
	s := fmt.Sprint(p.Rows, " ", p.Columns, "\n")

	for i, rows := range p.Board {
		for _, cell := range rows {
			s += cell.String()
		}
		s += fmt.Sprint(" ", p.RowTentCount[i], "\n")
	}

	for i, count := range p.ColumnTentCount {
		s += fmt.Sprint(count)
		if i != len(p.ColumnTentCount)-1 { // not last column
			s += " "
		}
	}

	return s
}
