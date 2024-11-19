package updater

import (
	"os"
	"regexp"
)

// UpdateHTML updates asset references in a markup file.
func UpdateHTML(filePath string, replacements map[string]string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	html := string(content)
	for oldPath, newPath := range replacements {
		// Use regex to replace references accurately
		regex := regexp.MustCompile(regexp.QuoteMeta(oldPath))
		html = regex.ReplaceAllString(html, newPath)
	}

	return os.WriteFile(filePath, []byte(html), 0644)
}
