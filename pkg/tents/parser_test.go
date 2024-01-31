package tents

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/txtar"
)

func TestParseTestFiles(t *testing.T) {
	testdataPath := "./testdata"

	files, err := os.ReadDir(testdataPath)
	assert.NoError(t, err, "Failed to read testdata directory")

	for _, file := range files {
		t.Run(file.Name(), func(t *testing.T) {
			content, err := os.ReadFile(filepath.Join(testdataPath, file.Name()))
			assert.NoError(t, err, "Failed to read file %s", file.Name())

			archive := txtar.Parse(content)

			var puzzleData, parsedPuzzleData []byte

			for _, f := range archive.Files {
				if f.Name == "puzzle" {
					puzzleData = f.Data
				} else if f.Name == "parsed" {
					parsedPuzzleData = f.Data
				}
			}

			assert.NotNil(t, puzzleData, "Missing puzzle in file %s", file.Name())
			assert.NotNil(t, parsedPuzzleData, "Missing parsed puzzle in file %s", file.Name())

			acutalPuzzle, err := Parse(string(puzzleData))
			assert.NoError(t, err, "Parsing error in file %s", file.Name())

			var expectedPuzzle Puzzle
			err = json.Unmarshal([]byte(parsedPuzzleData), &expectedPuzzle)
			assert.NoError(t, err, "Failed to unmarshal expected puzzle in file %s", file.Name())

			assert.Equal(t, expectedPuzzle, acutalPuzzle, "Parsed puzzle does not match expected puzzle in file %s", file.Name())
		})
	}
}
