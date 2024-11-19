package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

// HashFile computes the SHA256 hash of a file's content and returns the first n characters of the hash.
func HashFile(path string, n int) (string, error) {
	// Read the file content
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Compute the SHA256 hash
	hash := sha256.Sum256(data)

	// Convert the hash to a hexadecimal string and truncate to the desired length
	fullHash := hex.EncodeToString(hash[:])
	if n > len(fullHash) {
		n = len(fullHash) // Ensure we don't exceed the hash length
	}

	return fullHash[:n], nil
}

// RenameFile renames a file to include the hash in its name.
// For example, "style.css" with hash "abcd1234" becomes "style.abcd1234.css".
func RenameFile(path, hash string) (string, error) {
	// Split the file path into directory, file name, and extension
	dir, file := filepath.Split(path)
	ext := filepath.Ext(file)
	base := file[:len(file)-len(ext)] // Remove the extension to get the base name

	// Create the new file name with the hash
	newName := fmt.Sprintf("%s.%s%s", base, hash, ext)
	newPath := filepath.Join(dir, newName)

	// Rename the file
	err := os.Rename(path, newPath)
	if err != nil {
		return "", fmt.Errorf("failed to rename file %s to %s: %w", path, newPath, err)
	}

	return newPath, nil
}
