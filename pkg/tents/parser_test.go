package tents

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
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

			var puzzleData, parsedPuzzleData, aspData []byte

			for _, f := range archive.Files {
				if f.Name == "puzzle" {
					puzzleData = f.Data
				} else if f.Name == "parsed" {
					parsedPuzzleData = f.Data
				} else if f.Name == "asp" {
					aspData = f.Data
				}
			}

			assert.NotNil(t, puzzleData, "Missing puzzle in file %s", file.Name())
			assert.NotNil(t, parsedPuzzleData, "Missing parsed puzzle in file %s", file.Name())
			assert.NotNil(t, aspData, "Missing asp in file %s", file.Name())

			acutalPuzzle, err := ParsePuzzle(string(puzzleData))
			assert.NoError(t, err, "Parsing error in file %s", file.Name())

			var expectedPuzzle Puzzle
			err = json.Unmarshal([]byte(parsedPuzzleData), &expectedPuzzle)
			assert.NoError(t, err, "Failed to unmarshal expected puzzle in file %s", file.Name())

			assert.Equal(t, expectedPuzzle, acutalPuzzle, "Parsed puzzle does not match expected puzzle in file %s", file.Name())

			aspLines := strings.Split(string(aspData), "\n")
			assertEqualAsp(t, aspLines, acutalPuzzle.ToAspProgram())

			actualAspPuzzle, err := ParseAsp(string(aspData))
			assert.NoError(t, err, "Parsing error in file %s", file.Name())
			assert.Equal(t, acutalPuzzle, actualAspPuzzle, "Parsed puzzle does not match expected puzzle in file %s", file.Name())
		})
	}
}

func assertEqualAsp(t *testing.T, expected, actual []string) {
	// remove all whitespace and sort then compare
	var expectedCleaned, actualCleaned []string
	for i := range expected {
		clean := strings.TrimSpace(expected[i])
		if clean != "" {
			expectedCleaned = append(expectedCleaned, clean)
		}
	}

	for i := range actual {
		clean := strings.TrimSpace(actual[i])
		if clean != "" {
			actualCleaned = append(actualCleaned, clean)
		}
	}

	sort.Strings(expectedCleaned)
	sort.Strings(actualCleaned)

	assert.ElementsMatch(t, expectedCleaned, actualCleaned)
}
