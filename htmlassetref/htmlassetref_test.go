package htmlassetref_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/ryanburnette/go-hash-assets/htmlassetref"
)

func TestUpdateAssetRefs(t *testing.T) {
	// Read the original HTML from _content.html
	originalHTML, err := ioutil.ReadFile("_content.html")
	if err != nil {
		t.Fatalf("failed to read _content.html: %v", err)
	}

	// Define the expected modified HTML from _modified.html
	expectedHTML, err := ioutil.ReadFile("_modified.html")
	if err != nil {
		t.Fatalf("failed to read _modified.html: %v", err)
	}

	// Define the callback function to add _x before the extension
	callback := func(ref string) string {
		// Split the reference by the dot (.) to separate the filename and extension
		parts := strings.Split(ref, ".")
		if len(parts) > 1 {
			// Add _x before the extension
			parts[0] += "_x"
		}
		// Reconstruct the reference
		return strings.Join(parts, ".")
	}

	// Update the asset references in the original HTML
	modifiedHTML := htmlassetref.UpdateAssetRefs(string(originalHTML), callback)

	// Compare the modified HTML with the expected HTML
	if modifiedHTML != string(expectedHTML) {
		t.Errorf("modified HTML does not match the expected HTML")

		// Print the expected and modified HTML to the console
		fmt.Println("--- Expected HTML ---")
		fmt.Println(string(expectedHTML))
		fmt.Println("--- Modified HTML ---")
		fmt.Println(modifiedHTML)
	}
}
