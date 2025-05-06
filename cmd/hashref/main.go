package main

import (
	"fmt"
	"log"

	"github.com/ryanburnette/hashref/pkg/config"
	"github.com/ryanburnette/hashref/pkg/fileprocessor"
	"github.com/ryanburnette/hashref/pkg/hasher"
	"github.com/ryanburnette/hashref/pkg/htmlparser"
)

func main() {
	cfg := config.ParseFlags()

	// Initialize components
	hashGen := hasher.NewMD5Hasher(cfg.HashLength)
	parser := htmlparser.NewParser(cfg.Extensions)
	processor := fileprocessor.NewProcessor(parser, hashGen, cfg)

	// Process directory
	if err := processor.ProcessDir(cfg.Dir); err != nil {
		log.Fatalf("Error processing directory: %v", err)
	}

	if cfg.DryRun {
		fmt.Println("Dry run complete. No files were modified.")
	} else {
		fmt.Println("Processing complete. Files updated.")
	}
}
