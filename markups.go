package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
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

func (mu *Markup) updateRefs(opts Opts) error {
	dryRunMessage := ""
	if opts.dryRun {
		dryRunMessage = "Dry run: "
	}

	fmt.Printf("%vUpdate markup: %v\n", dryRunMessage, mu.filePath)

	return nil
}
