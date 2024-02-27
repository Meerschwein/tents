package tents

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/Meerschwein/tents/pkg/asp"
	"github.com/Meerschwein/tents/pkg/asp/solution"
	"github.com/Meerschwein/tents/pkg/clingo"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/txtar"
)

func TestParseTestFiles(t *testing.T) {
	testdataPath := "./testdata"

	files, err := os.ReadDir(testdataPath)
	assert.NoError(t, err, "Failed to read testdata directory")

	for _, file := range files {
		file := file
		t.Run(file.Name(), func(t *testing.T) {
			t.Parallel()
			content, err := os.ReadFile(filepath.Join(testdataPath, file.Name()))
			assert.NoError(t, err, "Failed to read file %s", file.Name())

			archive := txtar.Parse(content)

			var puzzleData, jsonData, aspData, solutionData []byte

			for _, f := range archive.Files {
				switch f.Name {
				case "puzzle":
					puzzleData = f.Data
				case "json":
					jsonData = f.Data
				case "asp":
					aspData = f.Data
				case "solution":
					solutionData = f.Data
				}
			}

			assert.NotNil(t, puzzleData)

			p, err := ParsePuzzle(string(puzzleData))
			assert.NoError(t, err)

			if jsonData != nil {
				t.Run("json", func(t *testing.T) { jsonPuzzleTest(t, p, jsonData) })
			}
			if aspData != nil {
				t.Run("asp", func(t *testing.T) { aspPuzzleTest(t, p, aspData) })
			}
			if solutionData != nil {
				t.Run("solution", func(t *testing.T) { solutionTest(t, p, solutionData) })
			}
		})
	}
}

func jsonPuzzleTest(t *testing.T, p Puzzle, jsonData []byte) {
	var expectedJson Puzzle
	err := json.Unmarshal([]byte(jsonData), &expectedJson)
	assert.NoError(t, err)

	assert.Equal(t, expectedJson, p)
}

func aspPuzzleTest(t *testing.T, p Puzzle, aspData []byte) {
	asp, err := asp.ParsePredicates(strings.Split(string(aspData), "\n"))
	assert.NoError(t, err)
	assert.ElementsMatch(t, asp, p.ToAsp())
}

func solutionTest(t *testing.T, p Puzzle, solutionData []byte) {
	pstr := ""
	for _, p := range p.ToAsp() {
		pstr += p.String()
	}

	for name, program := range solution.Solutions {
		t.Run(name, func(t *testing.T) {
			program += pstr
			cr, err := clingo.Run(strings.NewReader(program))
			assert.NoError(t, err)

			if strings.TrimSpace(string(solutionData)) == "INCONSISTENT" {
				assert.Equal(t, "INCONSISTENT", cr.Delimiter)
				if cr.Delimiter == "ANSWER" {
					p, err := ParseAsp(cr.Predicates)
					assert.NoError(t, err)
					t.Log("\n",p.ToPuzzle())
				}
				assert.Equal(t, clingo.QueryIsFalse.Error(), cr.ExitCode.Error())
				return
			}

			assert.True(t, cr.GoodExitCode(), cr.ExitCode)
			assert.Equal(t, "ANSWER", cr.Delimiter)

			pa, err := ParseAsp(cr.Predicates)
			assert.NoError(t, err)

			pe, err := ParsePuzzle(string(solutionData))
			assert.NoError(t, err)
			assertPuzzleEqual(t, pe, pa)
		})
	}
}

func assertPuzzleEqual(t *testing.T, expected, actual Puzzle) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected\n%v\nGot\n%v", expected.ToPuzzle(), actual.ToPuzzle())
	}
}
