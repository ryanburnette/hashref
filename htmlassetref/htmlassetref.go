package htmlassetref

import (
	"fmt"
	"regexp"
	"strings"
)

type AssetReferenceCallback func(reference, filePath string) string

func UpdateHTMLAssetRefs(htmlContent string, callback AssetReferenceCallback) (string, error) {
	// Regular expression pattern to match asset references in HTML tags
	pattern := `(?i)(?s)(<link[^>]*>|<img[^>]*>|<script[^>]*>|<source[^>]*>|<audio[^>]*>|<video[^>]*>|<a[^>]*>)(.*?)(href|src|poster)\s*=\s*(['"])(.*?)\{4}`

	// Compile the regular expression
	regExp, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("error compiling regular expression: %w", err)
	}

	// Find all matches of the regular expression in the HTML content
	matches := regExp.FindAllStringSubmatch(htmlContent, -1)

	// Iterate over the matches and invoke the callback function with each asset reference
	modifiedHTMLContent := htmlContent
	for _, match := range matches {
		// tag := match[1]
		attribute := match[5]

		// Invoke the callback function with the asset reference and file path, and get the updated file path
		modifiedFilePath := callback(attribute, attribute)

		// Update the asset reference with the modified file path
		modifiedTag := strings.Replace(match[0], attribute, modifiedFilePath, 1)
		modifiedHTMLContent = strings.Replace(modifiedHTMLContent, match[0], modifiedTag, 1)
	}

	return modifiedHTMLContent, nil
}
