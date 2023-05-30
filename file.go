package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// FindFilesByExt recursively finds all files with the specified extensions in the given root
func FindFilesByExt(fsys fs.FS, extensions []string) ([]string, error) {
	var list []string

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && hasExtAny(d.Name(), extensions) {
			list = append(list, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return list, nil
}

// hasExtAny checks if the given name has any of the specified suffixes.
func hasExtAny(filePath string, extensions []string) bool {
	for _, extension := range extensions {
		if strings.HasSuffix(filePath, extension) {
			return true
		}
	}
	return false
}

// GetFileHash generates an MD5 hash for the given file
func GetFileHash(fsys fs.FS, filePath string, hashLen int) (string, error) {
	input, err := fsys.Open(filePath)
	if err != nil {
		return "", err
	}

	hasher := md5.New()
	{
		_, err := io.Copy(hasher, input)
		if err != nil {
			return "", fmt.Errorf("failed to generate hash: %w", err)
		}
	}

	hash := hex.EncodeToString(hasher.Sum(nil))

	// Take the first (hashLen) characters of the hash
	if len(hash) > hashLen {
		hash = hash[:hashLen]
	}

	return hash, nil
}

func CreateHashedFilePath(filePath string, hash string) string {
	ext := filepath.Ext(filePath)
	base := strings.TrimSuffix(filePath, ext)
	newFilePath := fmt.Sprintf("%s.%s%s", base, hash, ext)
	return newFilePath
}

func RenameFile(directory string, filePath string, newFilePath string) error {
	oldFullPath := filepath.Join(directory, filePath)
	newFullPath := filepath.Join(directory, newFilePath)

	err := os.Rename(oldFullPath, newFullPath)
	if err != nil {
		return err
	}

	return nil
}
