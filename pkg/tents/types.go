package tents

type Cell int

const (
	Empty Cell = iota
	Tree
	Tent
)

func (c Cell) String() string {
	switch c {
	case Empty:
		return "."
	case Tree:
		return "T"
	case Tent:
		return "A"
	}

	panic("invalid cell")
}

type Puzzle struct {
	Rows            int
	Columns         int
	Board           [][]Cell
	RowTentCount    []int
	ColumnTentCount []int
}

func new(rows, columns int) Puzzle {
	puzzle := Puzzle{
		Rows:            rows,
		Columns:         columns,
		Board:           make([][]Cell, rows),
		RowTentCount:    make([]int, rows),
		ColumnTentCount: make([]int, columns),
	}

	for i := range puzzle.Board {
		puzzle.Board[i] = make([]Cell, columns)
	}

	return puzzle
}
