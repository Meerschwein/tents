package tents

import (
	"fmt"

	"github.com/Meerschwein/tents/pkg/asp"
	"github.com/Meerschwein/tents/pkg/util"
)

func filterByName(preds []asp.Predicate, name string) []asp.Predicate {
	return util.Filter(preds, func(p asp.Predicate) bool { return p.Name == name })
}

func ParseAsp(input []asp.Predicate) (Puzzle, error) {
	rowsp := filterByName(input, "lines")
	if len(rowsp) != 1 {
		return Puzzle{}, fmt.Errorf("expected 1 lines fact, got %d", len(rowsp))
	}
	if len(rowsp[0].Args) != 1 {
		return Puzzle{}, fmt.Errorf("expected 1 argument for lines fact, got %v", rowsp[0])
	}
	rows := rowsp[0].Args[0]

	columnsp := filterByName(input, "columns")
	if len(columnsp) != 1 {
		return Puzzle{}, fmt.Errorf("expected 1 columns fact, got %d", len(columnsp))
	}
	if len(columnsp[0].Args) != 1 {
		return Puzzle{}, fmt.Errorf("expected 1 argument for columns fact, got %v", columnsp[0])
	}
	columns := columnsp[0].Args[0]

	p := new(rows, columns)

	colsums := filterByName(input, "colsum")
	if len(colsums) != columns {
		return Puzzle{}, fmt.Errorf("expected %d colsum facts, got %d", columns, len(colsums))
	}
	for _, colsum := range colsums {
		if len(colsum.Args) != 2 {
			return Puzzle{}, fmt.Errorf("expected 2 arguments for colsum fact, got %v", colsum)
		}
		if colsum.Args[0] < 1 || colsum.Args[0] > columns {
			return Puzzle{}, fmt.Errorf("column index out of bounds: %d", colsum.Args[0])
		}
		p.ColumnTentCount[colsum.Args[0]-1] = colsum.Args[1]
	}

	rowsums := filterByName(input, "rowsum")
	if len(rowsums) != rows {
		return Puzzle{}, fmt.Errorf("expected %d rowsum facts, got %d", rows, len(rowsums))
	}
	for _, rowsum := range rowsums {
		if len(rowsum.Args) != 2 {
			return Puzzle{}, fmt.Errorf("expected 2 arguments for rowsum fact, got %v", rowsum)
		}
		if rowsum.Args[0] < 1 || rowsum.Args[0] > rows {
			return Puzzle{}, fmt.Errorf("row index out of bounds: %d", rowsum.Args[0])
		}
		p.RowTentCount[rowsum.Args[0]-1] = rowsum.Args[1]
	}

	trees := filterByName(input, "tree")
	for _, tree := range trees {
		if len(tree.Args) != 2 {
			return Puzzle{}, fmt.Errorf("expected 2 arguments for tree fact, got %v", tree)
		}
		if tree.Args[0] < 1 || tree.Args[0] > rows {
			return Puzzle{}, fmt.Errorf("row index out of bounds: %d", tree.Args[0])
		} else if tree.Args[1] < 1 || tree.Args[1] > columns {
			return Puzzle{}, fmt.Errorf("column index out of bounds: %d", tree.Args[1])
		}
		p.Board[tree.Args[0]-1][tree.Args[1]-1] = Tree
	}

	tents := filterByName(input, "tent")
	for _, tent := range tents {
		if len(tent.Args) != 2 {
			return Puzzle{}, fmt.Errorf("expected 2 arguments for tent fact, got %v", tent)
		}
		if tent.Args[0] < 1 || tent.Args[0] > rows {
			return Puzzle{}, fmt.Errorf("row index out of bounds: %d", tent.Args[0])
		} else if tent.Args[1] < 1 || tent.Args[1] > columns {
			return Puzzle{}, fmt.Errorf("column index out of bounds: %d", tent.Args[1])
		}
		p.Board[tent.Args[0]-1][tent.Args[1]-1] = Tent
	}

	return p, nil
}

func (p Puzzle) ToAsp() []asp.Predicate {
	preds := []asp.Predicate{
		asp.NewPredicate("lines", p.Rows),
		asp.NewPredicate("columns", p.Columns),
	}

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
			preds = append(preds, asp.NewPredicate(name, i+1, j+1))
		}
	}

	for i, count := range p.RowTentCount {
		preds = append(preds, asp.NewPredicate("line", i+1))
		preds = append(preds, asp.NewPredicate("rowsum", i+1, count))
	}

	for i, count := range p.ColumnTentCount {
		preds = append(preds, asp.NewPredicate("column", i+1))
		preds = append(preds, asp.NewPredicate("colsum", i+1, count))
	}

	return preds
}
