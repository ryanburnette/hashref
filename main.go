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
}
