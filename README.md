# [hash-assets](https://github.com/ryanburnette/go-hash-assets)

> work in progress, rewriting in Go

This tool is designed to hash your asset files (.css and .js) and update the
references in the markup (.html). It is expected that this tool be will be run
on generated assets so it updates everything in place.

My own approach is to reference all assets to the root `/`, so it only works
with assets referenced at the root level. It doesn't work with relative assets
or the base dir attribute.

## Usage

```
hash-assets -dir <directory> [-d] [-hash-len] [-asset-exts] [-markup-exts]
```

```
Usage:
	-asset-ext string
		Asset file extensions (default ".css,.js")
	-d	Dry run
	-dir string
		Working directory path, required
	-hash-len string
		Hash string length (default "10")
```
