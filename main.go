package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// Command line flags
	directory := flag.String("dir", "", "Working directory path, required")
	dryRun := flag.Bool("d", false, "Dry run")
	assetExtensions := flag.String("asset-ext", ".css,.js", "Asset file extensions")
	// markupExtensions := flag.String("markup-ext", ".html", "Markup file extensions")
	hashLenArg := flag.String("hash-len", "10", "Hash string length")
	flag.Parse()

	// Check if the directory flag is empty
	if *directory == "" {
		fmt.Printf("Error: -dir is required\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check if hash len arg is number, and make hashLen int
	hashLen, err := strconv.Atoi(*hashLenArg)
	if err != nil {
		fmt.Printf("Error: -hash-len must be a number")
		os.Exit(1)
	}
	// fmt.Printf("xx %v %v", hashLenArg, hashLen)

	// Scan the directory for asset files (CSS and JS)
	assetFiles, err := findFilesByExt(*directory, strings.Split(*assetExtensions, ","))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf(strings.Join(assetFiles, "\n"))
	// fmt.Printf("\n")
	os.Exit(1)

	var processedAssetFiles []string
	for _, file := range assetFiles {
		relPath, _ := filepath.Rel(*directory, file) // Get the relative path of the asset file with respect to the working directory

		newFilePath, err := processAssetFile(*directory, relPath, *dryRun, hashLen)
		if err != nil {
			fmt.Printf("Error processing asset file: %s: %v\n", file, err)
			continue
		}
		processedAssetFiles = append(processedAssetFiles, newFilePath)
	}
	// fmt.Printf("processedAssetFiles: %v\n", processedAssetFiles)

	// Scan the directory for markup files (HTML)
	// markupFiles, err := findFilesByExt(*directory, strings.Split(*markupExtensions, ","))
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	os.Exit(1)
	// }

	// Process markup files
	// for _, file := range markupFiles {
	// 	if err := ProcessMarkupFile(file, *dryRun); err != nil {
	// 		fmt.Printf("Error processing markup file: %s: %v\n", file, err)
	// 	}
	// }
}

// generateHash generates an MD5 hash for the given content and returns it as a string.
func generateHash(content []byte, hashLength int) (string, error) {
	hasher := md5.New()
	_, err := hasher.Write(content)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash: %w", err)
	}

	hash := hex.EncodeToString(hasher.Sum(nil))

	fmt.Printf("xx %v", hashLength)

	// Take the first (hashLength) characters of the hash
	if len(hash) > hashLength {
		hash = hash[:hashLength]
	}

	return hash, nil
}

// processAssetFile processes the given asset file by generating a new file path with a hashed name.
// If dryRun is true, it only prints the file rename operation without actually renaming the file.
func processAssetFile(directory string, relPath string, dryRun bool, hashLength int) (string, error) {
	// Get the absolute file path by joining the working directory and the relative path
	file := filepath.Join(directory, relPath)
	filePath := filepath.Dir(file)
	fileName := filepath.Base(file)
	fileExt := filepath.Ext(fileName)

	// Generate the new file name with a hashed value
	hashedName, err := generateHash([]byte(fileName), hashLength)
	if err != nil {
		return "", fmt.Errorf("failed to generate hashed name for file %s: %w", file, err)
	}

	newFileName := fmt.Sprintf("%s.%s%s", strings.TrimSuffix(fileName, fileExt), hashedName, fileExt)
	newFilePath := filepath.Join(filePath, newFileName)

	if !dryRun { // Skip file renaming if dryRun is true
		err = os.Rename(file, newFilePath)
		if err != nil {
			return "", fmt.Errorf("failed to rename file %s to %s: %w", file, newFilePath, err)
		}
	}

	fmt.Printf("Rename file: %s -> %s\n", file, newFilePath)
	return newFilePath, nil
}

// findFilesByExt recursively finds all files with the specified extensions in the given root
func findFilesByExt(root string, extensions []string) ([]string, error) {
	var files []string

	fsys := os.DirFS(root)
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && hasSuffixAny(d.Name(), extensions) {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

// hasSuffixAny checks if the given name has any of the specified suffixes.
func hasSuffixAny(name string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(name, suffix) {
			return true
		}
	}
	return false
}
