package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/ryanburnette/go-hash-assets/htmlassetref"
)

type Markup struct {
	filePath string
	dirPath  string
}

func NewMarkup(filePath string) (*Markup, error) {
	return &Markup{
		filePath: filePath,
		dirPath:  filepath.Dir(filePath),
	}, nil
}

type MarkupProc struct {
	opts    Opts
	fsys    fs.FS
	markups []Markup
}

func NewMarkupProc(opts Opts, fsys fs.FS) (*MarkupProc, error) {
	return &MarkupProc{
		opts:    opts,
		fsys:    fsys,
		markups: []Markup{},
	}, nil
}

func (mp *MarkupProc) findMarkups() error {
	files, err := FindFilesByExt(mp.fsys, mp.opts.markupExtensions)
	if err != nil {
		return err
	}

	for _, filePath := range files {
		mu, err := NewMarkup(filePath)
		if err != nil {
			return err
		}
		mp.markups = append(mp.markups, *mu)
	}

	return nil
}

func (mp *MarkupProc) updateRefs() error {
	return nil
}

func (mu *Markup) updateRefs(opts Opts, mp *MarkupProc, ap *AssetProc) error {
	dryRunMessage := ""
	if opts.dryRun {
		dryRunMessage = "Dry run: "
	}

	fmt.Printf("%vUpdate markup: %v\n", dryRunMessage, mu.filePath)

	// Open the file from the filesystem
	file, err := mp.fsys.Open(mu.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//var builder strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		_ = htmlassetref.UpdateAssetRefs(line, func(ref string) string {
			// absPath is absolute to the file system, thats how we find the asset
			absPath := filepath.Join(mu.dirPath, ref)
			asset := ap.findAsset(absPath)

			fmt.Printf(".. %v ** %v ** %v %#v \n", mu.filePath, ref, absPath, asset)
			return ref
		})
	}

	return nil
}
