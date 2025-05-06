package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

// Hasher generates content-based hashes for files.
type Hasher struct {
	hashLength int
}

// NewMD5Hasher creates a new Hasher with MD5 and specified length.
func NewMD5Hasher(hashLength int) *Hasher {
	return &Hasher{hashLength: hashLength}
}

// HashFile generates a hash for a file's content.
func (h *Hasher) HashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	hashBytes := hash.Sum(nil)
	fullHash := hex.EncodeToString(hashBytes)
	if h.hashLength > 0 && h.hashLength < len(fullHash) {
		return fullHash[:h.hashLength], nil
	}
	return fullHash, nil
}

// NewFilename generates a new filename with the hash.
func (h *Hasher) NewFilename(path, hash string) string {
	ext := filepath.Ext(path)
	name := path[:len(path)-len(ext)]
	return name + "." + hash + ext
}
