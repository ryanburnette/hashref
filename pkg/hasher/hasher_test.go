package hasher

import (
	"os"
	"testing"
)

func TestHashFile(t *testing.T) {
	filePath := "../../example/css/style.css"

	// Expected hash for the file
	expectedHash := "e3b0c442"

	// Compute the hash of the file
	hash, err := HashFile(filePath, len(expectedHash))
	if err != nil {
		t.Fatalf("HashFile failed: %v", err)
	}

	// Verify the hash matches the expected value
	if hash != expectedHash {
		t.Errorf("Expected hash %s, got %s", expectedHash, hash)
	}
}

func TestRenameFile(t *testing.T) {
	filePath := "../../example/css/style.css"

	// Expected hash for the file
	expectedHash := "e3b0c442"

	// Act: Rename the file
	newPath, err := RenameFile(filePath, expectedHash)
	if err != nil {
		t.Fatalf("RenameFile failed: %v", err)
	}

	// Verify the file was renamed correctly
	expectedNewPath := "../../example/css/style.e3b0c442.css"
	if newPath != expectedNewPath {
		t.Errorf("Expected new path %s, got %s", expectedNewPath, newPath)
	}

	// Verify that the renamed file exists
	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		t.Errorf("Renamed file does not exist: %s", newPath)
	}

	// Rename the file back to its original name for future tests
	err = os.Rename(newPath, filePath)
	if err != nil {
		t.Fatalf("Failed to rename file back to original: %v", err)
	}
}
