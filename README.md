# hashref

**hashref** is a command-line tool for processing static website assets. It identifies references to static assets (e.g., `.css`, `.js`) in `.html` files, renames the assets by appending a content-based hash to their filenames for cache-busting, and updates the HTML to reflect the new filenames. It operates in-place on a `dist` directory (post-build, static-generated files) and supports a dry-run mode to preview changes.

## Features
- Processes `.css` and `.js` assets in `<link>` and `<script>` tags.
- Supports absolute (`/assets/styles.css`) and relative (`../../assets/styles.css`) paths.
- Appends MD5 hash to filenames (e.g., `styles.abcdef123.css`).
- Configurable hash length (default: 8 characters).
- Optional additional extensions (e.g., `.png`, `.jpg` for `<img>` tags).
- Dry-run mode to preview changes.
- Safe for `dist` directories (no source code modifications).

## Installation
```bash
go install github.com/ryanburnette/hashref@latest
```

## Usage
Run `hashref` on a `dist` directory:
```bash
hashref --dir dist
```

Preview changes without modifying files:
```bash
hashref --dir dist --dry-run
```

Customize hash length:
```bash
hashref --dir dist --hash-length 16
```

Process additional extensions (e.g., images):
```bash
hashref --dir dist --extensions png,jpg
```

### CLI Flags
- `--dir`: Directory to process (default: `./dist`).
- `--dry-run`: Preview changes without modifying files.
- `--hash-length`: Length of hash in filenames (default: 8).
- `--extensions`: Comma-separated additional file extensions (e.g., `png,jpg`).

## Demo
The `testdata/demo` directory contains sample files:
- `index.html`: References assets with absolute paths.
- `about.html`: References assets with relative paths.
- `assets/styles.css`: Sample CSS file.
- `assets/app.js`: Sample JS file.

Run the demo:
```bash
hashref --dir testdata/demo --dry-run
```

## Notes
- Only `<link href>` and `<script src>` tags are processed. Inline `<style>`, `<script>`, or `<meta>` tags are not supported.
- Missing assets are logged as warnings and skipped.
- Malformed HTML is handled gracefully.

## Contributing
1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/foo`).
3. Commit changes (`git commit -m 'Add foo'`).
4. Push to the branch (`git push origin feature/foo`).
5. Open a pull request.

Run tests:
```bash
go test ./...
```

## License
MIT License. See [LICENSE](LICENSE) for details.
