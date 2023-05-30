package htmlassetref

import (
	"regexp"
	"strings"
)

// UpdateAssetRefs updates asset references in the HTML using the provided callback.
func UpdateAssetRefs(html string, callback func(ref string) string) string {
	re := regexp.MustCompile(`(src|href|srcset)=(['"])?([^>\s]+)(['"\s>])?`)
	return re.ReplaceAllStringFunc(html, func(matched string) string {
		attr := re.ReplaceAllString(matched, "$1")
		quote := re.ReplaceAllString(matched, "$2")
		ref := re.ReplaceAllString(matched, "$3")
		afterRef := re.ReplaceAllString(matched, "$4")

		updatedRef := callback(ref)

		// Reconstruct the attribute with the updated reference
		return attr + "=" + quote + updatedRef + strings.TrimSuffix(afterRef, quote)
	})
}
