# [hashref](https://github.com/ryanburnette/go-hashref)

hashref is a command-line tool designed to simplify cache busting for generated `.html`, `.css`, and `.js` files. It simplifies the process of cache busting by adding a short hash to the filenames of generated assets such as .css and .js files. By incorporating the hash in the filenames, hashref enables cache busting when the content of the assets changes.

## Features

-   Cache busting: hashref adds a hash to the filenames of generated assets, allowing for reliable cache busting when the assets change.
-   Compatible with cache strategy: hashref seamlessly integrates with cache strategies that rely on HTTP server configurations to handle asset caching. It leverages the hash in the filenames to trigger cache updates when changes occur.
-   Support for various asset references: hashref handles absolute, relative, and base style asset references within HTML, ensuring consistent cache busting across different reference types.
-   Internal references only: hashref operates exclusively on internal asset references within the generated HTML, CSS, and JS files. It does not modify or process external references.

By utilizing hashref, developers can effortlessly incorporate cache busting into their build process and enhance the efficiency of asset caching, ensuring that updated assets are served to end users when changes occur.

## Usage

```shell
Usage:
	hash-assets -dir <directory> [-d] [-hash-len] [-asset-exts] [-markup-exts]

Options:
	-dir string
		Specify the working directory path (required).
	-d
		Perform a dry run without making any changes.
	-hash-len string
		Specify the length of the hash string (default "10").
	-asset-ext string
		Specify the asset file extensions (default ".css,.js").
	-markup-ext string
		Specify the markup file extensions (default ".html,.htm").
```
