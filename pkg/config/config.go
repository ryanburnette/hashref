package config

import (
	"flag"
	"strings"
)

// Config holds CLI configuration.
type Config struct {
	Dir        string
	DryRun     bool
	HashLength int
	Extensions []string
}

// ParseFlags parses command-line flags and returns a Config.
func ParseFlags() *Config {
	dir := flag.String("dir", "./dist", "directory to process")
	dryRun := flag.Bool("dry-run", false, "preview changes without modifying files")
	hashLength := flag.Int("hash-length", 8, "length of hash in filenames")
	extensions := flag.String("extensions", "", "comma-separated additional file extensions (e.g., png,jpg)")

	flag.Parse()

	var extList []string
	if *extensions != "" {
		extList = strings.Split(*extensions, ",")
		for i, ext := range extList {
			extList[i] = strings.TrimSpace(ext)
			if !strings.HasPrefix(extList[i], ".") {
				extList[i] = "." + extList[i]
			}
		}
	}
	extList = append(extList, ".css", ".js") // Default extensions

	return &Config{
		Dir:        *dir,
		DryRun:     *dryRun,
		HashLength: *hashLength,
		Extensions: extList,
	}
}
