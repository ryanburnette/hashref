package main

import (
	"fmt"
	"io/fs"
	"strings"
)

// Asset ...
type Asset struct {
	filePath    string
	hash        string
	newFilePath string
}

func NewAsset(filePath string, ap *AssetProc) (*Asset, error) {
	hash, err := GetFileHash(ap.fsys, filePath, ap.opts.hashLen)
	if err != nil {
		return nil, err
	}
	newFilePath := CreateHashedFilePath(filePath, hash)
	return &Asset{
		filePath:    filePath,
		hash:        hash,
		newFilePath: newFilePath,
	}, nil
}

// AssetProc ...
type AssetProc struct {
	opts   Opts
	fsys   fs.FS
	assets []Asset
}

func NewAssetProc(opts Opts, fsys fs.FS) (*AssetProc, error) {
	return &AssetProc{
		opts:   opts,
		fsys:   fsys,
		assets: []Asset{},
	}, nil
}

func (ap *AssetProc) findAssets() error {
	files, err := FindFilesByExt(ap.fsys, ap.opts.assetExtensions)
	if err != nil {
		return err
	}

	for _, filePath := range files {
		asset, err := NewAsset(filePath, ap)
		if err != nil {
			return err
		}
		ap.assets = append(ap.assets, *asset)
	}

	return nil
}

func (ap *AssetProc) renameAssets() error {
	for _, a := range ap.assets {
		var dryRun string
		dryRun = ""
		if !ap.opts.dryRun {
			err := RenameFile(ap.opts.directory, a.filePath, a.newFilePath)
			if err != nil {
				return err
			}
		} else {
			dryRun = "Dry run: "
		}
		fmt.Printf("%vRename asset: %v -> %v\n", dryRun, a.filePath, a.newFilePath)
	}

	return nil
}

func (ap *AssetProc) findAsset(filePath string) *Asset {
	for _, a := range ap.assets {
		if strings.EqualFold(a.filePath, filePath) {
			return &a
		}
	}
	return nil
}
