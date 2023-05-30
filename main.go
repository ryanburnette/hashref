package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Opts struct {
	directory       string
	dryRun          bool
	assetExtensions []string
	hashLen         int
}

func main() {
	// Command line flags
	var directory string
	flag.StringVar(&directory, "dir", "", "Working directory path, required")
	dryRun := flag.Bool("d", false, "Dry run")
	assetExtArg := flag.String("asset-ext", ".css,.js", "Asset file extensions")
	assetExtensions := strings.Split(*assetExtArg, ",")
	// markupExtensions := flag.String("markup-ext", ".html", "Markup file extensions")
	hashLenArg := flag.String("hash-len", "8", "Hash string length")
	flag.Parse()

	// Check if the directory flag is empty
	if directory == "" {
		fmt.Printf("Error: -dir is required\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check if hash len arg is number, and make hashLen int
	hashLen, err := strconv.Atoi(*hashLenArg)
	if err != nil {
		fmt.Printf("Error: -hash-len must be a number\n")
		os.Exit(1)
	}

	opts := Opts{
		directory:       directory,
		dryRun:          *dryRun,
		assetExtensions: assetExtensions,
		hashLen:         hashLen,
	}

	// fmt.Printf("%#v\n", opts)

	fsys := os.DirFS(opts.directory)
	ap := NewAssetProc(opts, fsys)

	err = ap.findAssets()
	if err != nil {

	}

	err = ap.renameAssets()
	if err != nil {
		fmt.Printf("%#v\n", err)
		fmt.Printf("Error: failed to rename assets\n")
		os.Exit(1)
	}

	// fmt.Printf("%#v\n", ap)

	// // Scan the directory for asset files (CSS and JS)
	// exts := strings.Split(*assetExtensions, ",")
	// err = p.findFilesByExt(exts)
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("p.list: %v\n", strings.Join(p.list, "\n"))
	// fmt.Printf("\n")
	// os.Exit(1)

	// var processedAssetFiles []string
	// for _, filePath := range assetFiles {

	// 	newFilePath, err := p.processAssetFile(filePath, *dryRun, hashLen)
	// 	if err != nil {
	// 		fmt.Printf("Error processing asset file: %s: %v\n", filePath, err)
	// 		continue
	// 	}
	// 	processedAssetFiles = append(processedAssetFiles, newFilePath)
	// }
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

// processAssetFile processes the given asset file by generating a new file path with a hashed name.
// If dryRun is true, it only prints the file rename operation without actually renaming the file.
// func (p AssetProc) processAssetFile(filePath string, dryRun bool, hashLength int) (string, error) {
// 	// Get the absolute file path by joining the working directory and the relative path
// 	fileName := filepath.Base(filePath)
// 	fileExt := filepath.Ext(filePath)

// 	// Generate the new file name with a hashed value
// 	hashedName, err := p.generateHash(filePath, hashLength)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to generate hashed name for file %s: %w", filePath, err)
// 	}

// 	newFileName := fmt.Sprintf("%s.%s%s", strings.TrimSuffix(fileName, fileExt), hashedName, fileExt)
// 	newFilePath := filepath.Join(filePath, newFileName)

// 	if !dryRun { // Skip file renaming if dryRun is true
// 		err = os.Rename(filePath, newFilePath)
// 		if err != nil {
// 			return "", fmt.Errorf("failed to rename file %s to %s: %w", filePath, newFilePath, err)
// 		}
// 	}

// 	fmt.Printf("Rename file: %s -> %s\n", filePath, newFilePath)
// 	return newFilePath, nil
// }
