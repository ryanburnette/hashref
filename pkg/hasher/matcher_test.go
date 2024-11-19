package hasher

import (
	"path/filepath"
	"testing"
)

func TestMatchFiles(t *testing.T) {
	// Define the directory to search
	exampleDir := "../../example"

	// Define the pattern to match
	patterns := []string{"*.html"}

	// Call MatchFiles
	matchedFiles, err := MatchFiles(exampleDir, patterns)
	if err != nil {
		t.Fatalf("MatchFiles failed: %v", err)
	}

	// Expected file to be matched
	expectedFile := filepath.Join(exampleDir, "index.html")

	// Verify the matched files contain the expected file
	found := false
	for _, file := range matchedFiles {
		if file == expectedFile {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected file %s not found in matched files: %v", expectedFile, matchedFiles)
	}
}

func TestMatchFilesJS(t *testing.T) {
	// Define the directory to search
	exampleDir := "../../example"

	// Define the pattern to match
	patterns := []string{"*.js"}

	// Call MatchFiles
	matchedFiles, err := MatchFiles(exampleDir, patterns)
	if err != nil {
		t.Fatalf("MatchFiles failed: %v", err)
	}

	// Expected files to be matched
	expectedFiles := []string{
		filepath.Join(exampleDir, "js", "module.js"),
		filepath.Join(exampleDir, "js", "module2.js"),
		filepath.Join(exampleDir, "js", "foo", "module3.js"),
	}

	// Verify all expected files are found
	for _, expectedFile := range expectedFiles {
		found := false
		for _, file := range matchedFiles {
			if file == expectedFile {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected file %s not found in matched files: %v", expectedFile, matchedFiles)
		}
	}

	// Ensure no unexpected files are matched
	if len(matchedFiles) != len(expectedFiles) {
		t.Errorf("Expected %d matched files, got %d: %v", len(expectedFiles), len(matchedFiles), matchedFiles)
	}
}
