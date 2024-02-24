package tents

import (
	"fmt"
	"strconv"
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

func ParseAsp(input string) (Puzzle, error) {
	in := strings.Split(input, "\n")
	for i := range in {
		in[i] = strings.TrimSpace(in[i])
	}

	rowsf, err := findFact(in, "lines")
	if err != nil {
		return Puzzle{}, err
	}
	rows := rowsf[0]

	columnsf, err := findFact(in, "columns")
	if err != nil {
		return Puzzle{}, err
	}
	columns := columnsf[0]

	p := new(rows, columns)

	colsums, err := findFacts(in, "colsum")
	if err != nil {
		return Puzzle{}, err
	} else if len(colsums) != columns {
		return Puzzle{}, fmt.Errorf("expected %d colsum facts, got %d", columns, len(colsums))
	}
	for _, colsum := range colsums {
		if colsum[0] < 1 || colsum[0] > columns {
			return Puzzle{}, fmt.Errorf("column index out of bounds: %d", colsum[0])
		}
		p.ColumnTentCount[colsum[0]-1] = colsum[1]
	}

	rowsums, err := findFacts(in, "rowsum")
	if err != nil {
		return Puzzle{}, err
	} else if len(rowsums) != rows {
		return Puzzle{}, fmt.Errorf("expected %d rowsum facts, got %d", rows, len(rowsums))
	}
	for _, rowsum := range rowsums {
		if rowsum[0] < 1 || rowsum[0] > rows {
			return Puzzle{}, fmt.Errorf("row index out of bounds: %d", rowsum[0])
		}
		p.RowTentCount[rowsum[0]-1] = rowsum[1]
	}

	trees, err := findFacts(in, "tree")
	if err != nil {
		return Puzzle{}, err
	}
	for _, tree := range trees {
		if tree[0] < 1 || tree[0] > rows {
			return Puzzle{}, fmt.Errorf("row index out of bounds: %d", tree[0])
		} else if tree[1] < 1 || tree[1] > columns {
			return Puzzle{}, fmt.Errorf("column index out of bounds: %d", tree[1])
		}
		p.Board[tree[0]-1][tree[1]-1] = Tree
	}

	tents, err := findFacts(in, "tent")
	if err != nil {
		return Puzzle{}, err
	}
	for _, tent := range tents {
		if tent[0] < 1 || tent[0] > rows {
			return Puzzle{}, fmt.Errorf("row index out of bounds: %d", tent[0])
		} else if tent[1] < 1 || tent[1] > columns {
			return Puzzle{}, fmt.Errorf("column index out of bounds: %d", tent[1])
		}
		p.Board[tent[0]-1][tent[1]-1] = Tent
	}

	return p, nil
}

func findFact(facts []string, name string) ([]int, error) {
	for _, fact := range facts {
		if strings.HasPrefix(fact, name) {
			argsStr := strings.TrimPrefix(fact, name+"(")
			argsStr = strings.TrimSuffix(argsStr, ").")
			args := strings.Split(argsStr, ",")
			intArgs := []int{}
			for _, arg := range args {
				i, err := strconv.Atoi(arg)
				if err != nil {
					return nil, err
				}
				intArgs = append(intArgs, i)
			}
			return intArgs, nil
		}
	}
	return nil, fmt.Errorf("fact %s not found", name)
}

func findFacts(facts []string, name string) ([][]int, error) {
	found := [][]int{}
	for _, fact := range facts {
		if strings.HasPrefix(fact, name) {
			argsStr := strings.TrimPrefix(fact, name+"(")
			argsStr = strings.TrimSuffix(argsStr, ").")
			args := strings.Split(argsStr, ",")
			intArgs := []int{}
			for _, arg := range args {
				i, err := strconv.Atoi(arg)
				if err != nil {
					return nil, err
				}
				intArgs = append(intArgs, i)
			}
			found = append(found, intArgs)
		}
	}

	return found, nil
}
