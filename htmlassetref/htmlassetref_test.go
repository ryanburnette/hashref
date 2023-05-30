package htmlassetref

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestUpdateHTMLAssetRefs(t *testing.T) {
	// Read the content from the "_content.html" file
	contentBytes, err := ioutil.ReadFile("_content.html")
	if err != nil {
		t.Fatalf("failed to read _content.html file: %v", err)
	}
	content := string(contentBytes)

	// Define the expected modified content from the "_modified.html" file
	expectedModifiedBytes, err := ioutil.ReadFile("_modified.html")
	if err != nil {
		t.Fatalf("failed to read _modified.html file: %v", err)
	}
	expectedModified := string(expectedModifiedBytes)

	// Callback function to add "_x" before the extension of each file name
	callback := func(reference, filePath string) string {
		updatedFilePath := strings.TrimSuffix(filePath, ".") + "_x."
		return strings.Replace(reference, filePath, updatedFilePath, 1)
	}

	// Update the HTML asset references
	modifiedContent, err := UpdateHTMLAssetRefs(content, callback)
	if err != nil {
		t.Fatalf("failed to update HTML asset references: %v", err)
	}

	// Compare the modified content with the expected modified content
	if modifiedContent != expectedModified {
		fmt.Printf("# modified:\n%v\n# expected:\n%v\n", modifiedContent, expectedModified)
		t.Errorf("modified content does not match the expected result")
	}
}
