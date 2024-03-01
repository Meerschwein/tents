package tents

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParsePuzzle(input string) (Puzzle, error) {
	var lines []string
	for _, l := range strings.Split(input, "\n") { // remove empty lines
		if l != "" {
			lines = append(lines, l)
		}
	}

	if len(lines) < 3 {
		// first line: number of rows and columns
		// board
		// last line: number of tents per column
		return Puzzle{}, errors.New("input too short")
	}

	rowColNum := strings.Split(lines[0], " ")
	if len(rowColNum) != 2 {
		return Puzzle{}, errors.New("first line must contain two numbers")
	}

	numRows, err := strconv.Atoi(rowColNum[0])
	if err != nil {
		return Puzzle{}, errors.New("first number must be an integer")
	} else if len(lines) != numRows+2 { // +2 because of the first line and the last line
		return Puzzle{}, errors.New("number of rows does not match first number")
	}

	numCols, err := strconv.Atoi(rowColNum[1])
	if err != nil {
		return Puzzle{}, errors.New("second number must be an integer")
	}

	puzzle := new(numRows, numCols)

	for currRow, line := range lines[1 : numRows+1] {
		/* allow for one space between cells like this
		   A T A 2
		   T T . 0
		   A T A 2
		*/
		line = strings.ReplaceAll(line, " ", "")

		board := line[:numCols]
		if len(board) != numCols {
			return Puzzle{}, errors.New("number of columns does not match second number")
		}

		tentNum, err := strconv.Atoi(line[numCols:])
		if err != nil {
			return Puzzle{}, errors.New("tent number must be an integer")
		}
		puzzle.RowTentCount[currRow] = tentNum

		for currCol, cell := range board {
			switch cell {
			case '.':
				puzzle.Board[currRow][currCol] = Empty
			case 'T':
				puzzle.Board[currRow][currCol] = Tree
			case 'A':
				puzzle.Board[currRow][currCol] = Tent
			default:
				return Puzzle{}, errors.New("board must contain only '.', 'T' and 'A'")
			}
		}
	}

	numTentCols := strings.Split(lines[len(lines)-1], " ")
	if len(numTentCols) != numCols {
		return Puzzle{}, errors.New("number of tent columns does not match second number")
	}
	for currCol, tentNumStr := range numTentCols {
		tentNum, err := strconv.Atoi(tentNumStr)
		if err != nil {
			return Puzzle{}, errors.New("tent number must be an integer")
		}
		puzzle.ColumnTentCount[currCol] = tentNum
	}

	return puzzle, nil
}

var PrintSpaces = false

func (p Puzzle) ToPuzzle() string {
	s := fmt.Sprint(p.Rows, " ", p.Columns, "\n")

	for i, rows := range p.Board {
		for _, cell := range rows {
			s += cell.String()
			if PrintSpaces {
				s += " "
			}
		}
		if !PrintSpaces {
			s += " "
		}
		s += fmt.Sprint(p.RowTentCount[i], "\n")
	}

	for i, count := range p.ColumnTentCount {
		s += fmt.Sprint(count)
		if i != len(p.ColumnTentCount)-1 { // not last column
			s += " "
		}
	}

	return s
}
