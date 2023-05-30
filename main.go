package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Opts struct {
	directory        string
	dryRun           bool
	assetExtensions  []string
	markupExtensions []string
	hashLen          int
}

func main() {
	// Command line flags
	var directory string
	flag.StringVar(&directory, "dir", "", "Working directory path, required")
	dryRun := flag.Bool("d", false, "Dry run")
	assetExtArg := flag.String("asset-ext", ".css,.js", "Asset file extensions")
	assetExtensions := strings.Split(*assetExtArg, ",")
	markupExtArg := flag.String("markup-ext", ".html", "Markup file extensions")
	markupExtensions := strings.Split(*markupExtArg, ",")
	hashLenArg := flag.String("hash-len", "8", "Hash string length")
	flag.Parse()

	// Check if the directory flag is empty
	if directory == "" {
		fmt.Printf("Error: -dir is required\n")
		showUsage()
		os.Exit(1)
	}

	// Check if hash len arg is number, and make hashLen int
	hashLen, err := strconv.Atoi(*hashLenArg)
	if err != nil {
		fmt.Printf("Error: -hash-len must be a number\n")
		os.Exit(1)
	}

	// init opts
	opts := Opts{
		directory:        directory,
		dryRun:           *dryRun,
		assetExtensions:  assetExtensions,
		markupExtensions: markupExtensions,
		hashLen:          hashLen,
	}
	// fmt.Printf("%#v\n", opts)

	fsys := os.DirFS(opts.directory)

	ap, err := NewAssetProc(opts, fsys)
	if err != nil {
		fmt.Printf("Error: %#v", err)
		os.Exit(1)
	}

	err = ap.findAssets()
	if err != nil {
		fmt.Printf("Error: %#v", err)
		os.Exit(1)
	}

	err = ap.renameAssets()
	if err != nil {
		fmt.Printf("Error: %#v", err)
		os.Exit(1)
	}

	mp, err := NewMarkupProc(opts, fsys)
	if err != nil {
		fmt.Printf("Error: %#v", err)
		os.Exit(1)
	}

	err = mp.findMarkups()
	if err != nil {
		fmt.Printf("Error: %#v", err)
		os.Exit(1)
	}

	err = mp.markups[0].updateRefs(opts)
	if err != nil {
		fmt.Printf("Error: %#v", err)
		os.Exit(1)
	}

	// fmt.Printf("%#v\n", mp)
}

func showUsage() {
	fmt.Printf("Usage:\n")
	flag.PrintDefaults()
}
