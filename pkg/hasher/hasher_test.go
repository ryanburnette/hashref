package hasher

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHasher(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.css")
	if err := os.WriteFile(filePath, []byte("body {}"), 0644); err != nil {
		t.Fatal(err)
	}

	h := NewMD5Hasher(8)
	hash, err := h.HashFile(filePath)
	if err != nil {
		t.Fatalf("HashFile error: %v", err)
	}
	if len(hash) != 8 {
		t.Errorf("Expected hash length 8, got %d", len(hash))
	}

	newName := h.NewFilename("styles.css", hash)
	if !strings.HasPrefix(newName, "styles.") || !strings.HasSuffix(newName, ".css") {
		t.Errorf("Invalid new filename: %s", newName)
	}
}
