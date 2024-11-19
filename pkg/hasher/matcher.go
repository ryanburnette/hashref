package hasher

import (
	"os"
	"path/filepath"
	"strings"
)

// MatchFiles recursively finds files matching any of the provided patterns in the directory.
func MatchFiles(directory string, patterns []string) ([]string, error) {
	var matchedFiles []string

	// Walk through the directory
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check file against all patterns
		for _, pattern := range patterns {
			matched, err := filepath.Match(pattern, filepath.Base(path))
			if err != nil {
				return err
			}
			if matched {
				matchedFiles = append(matchedFiles, path)
				break
			}
		}

		return nil
	})

	return matchedFiles, err
}

// ParsePatterns converts a comma-separated string into a slice of patterns.
func ParsePatterns(patterns string) []string {
	return strings.Split(strings.TrimSpace(patterns), ",")
}
