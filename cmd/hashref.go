package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ryanburnette/hashref/pkg/hasher"
	"github.com/ryanburnette/hashref/pkg/updater"
)

var (
	assetMatchers  string
	markupMatchers string
	dryRun         bool
)

func main() {
	// Define CLI arguments
	flag.StringVar(&assetMatchers, "assets", "**/*.js,**/*.css", "Asset file match patterns (comma-separated)")
	flag.StringVar(&markupMatchers, "markup", "**/*.html", "Markup file match patterns (comma-separated)")
	flag.BoolVar(&dryRun, "dry-run", false, "Perform a dry run without making changes")

	// Usage message
	flag.Usage = func() {
		fmt.Println("Usage: hashref [dir] [options]")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
	}

	// Parse CLI arguments
	flag.Parse()

	// Ensure a directory is provided
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	dir := args[0]

	absDir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("Failed to resolve absolute path of directory: %v", err)
	}

	// Validate the directory
	if _, err := os.Stat(absDir); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", absDir)
	}

	// Verbose output
	fmt.Printf("Processing directory: %s\n", absDir)
	fmt.Printf("Asset match patterns: %s\n", assetMatchers)
	fmt.Printf("Markup match patterns: %s\n", markupMatchers)
	fmt.Printf("Dry run mode: %v\n", dryRun)

	// Process files
	processFiles(dir, assetMatchers, markupMatchers, dryRun)
}

func processFiles(directory, assetPatterns, markupPatterns string, dryRun bool) {
	// Parse matchers
	assetPatternList := hasher.ParsePatterns(assetPatterns)
	markupPatternList := hasher.ParsePatterns(markupPatterns)

	// Match and hash asset files
	assetFiles, err := hasher.MatchFiles(directory, assetPatternList)
	if err != nil {
		log.Fatalf("Error matching asset files: %v", err)
	}
	fmt.Printf("Matched asset files: %v\n", assetFiles)

	// Hash and rename asset files
	replacements := make(map[string]string)
	for _, file := range assetFiles {
		hash, err := hasher.HashFile(file, 8)
		if err != nil {
			log.Printf("Error hashing file %s: %v", file, err)
			continue
		}

		newPath := fmt.Sprintf("%s.%s", file, hash)
		if dryRun {
			fmt.Printf("[DRY-RUN] Would rename: %s -> %s\n", file, newPath)
		} else {
			renamedPath, err := hasher.RenameFile(file, hash)
			if err != nil {
				log.Printf("Error renaming file %s: %v", file, err)
				continue
			}
			replacements[file] = renamedPath
			fmt.Printf("Renamed: %s -> %s\n", file, renamedPath)
		}
	}

	// Match markup files
	markupFiles, err := hasher.MatchFiles(directory, markupPatternList)
	if err != nil {
		log.Fatalf("Error matching markup files: %v", err)
	}
	fmt.Printf("Matched markup files: %v\n", markupFiles)

	// Update references in markup files
	for _, markupFile := range markupFiles {
		if dryRun {
			fmt.Printf("[DRY-RUN] Would update references in: %s\n", markupFile)
		} else {
			if err := updater.UpdateHTML(markupFile, replacements); err != nil {
				log.Printf("Error updating markup file %s: %v", markupFile, err)
			} else {
				fmt.Printf("Updated references in: %s\n", markupFile)
			}
		}
	}
}
