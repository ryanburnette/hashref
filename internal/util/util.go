package util

import (
	"path/filepath"
	"strings"
)

// ResolveAssetPath converts an asset path (absolute or relative) to an absolute file path.
func ResolveAssetPath(htmlPath, assetPath, rootDir string) (string, error) {
	if strings.HasPrefix(assetPath, "/") {
		return filepath.Join(rootDir, assetPath[1:]), nil
	}
	dir := filepath.Dir(htmlPath)
	absPath := filepath.Join(dir, assetPath)
	cleanPath := filepath.Clean(absPath)
	return cleanPath, nil
}

// RelativizePath converts an absolute asset path to a path relative to the HTML file.
func RelativizePath(assetPath, htmlPath, rootDir string) string {
	if strings.HasPrefix(assetPath, rootDir) {
		relPath := strings.TrimPrefix(assetPath, rootDir)
		if !strings.HasPrefix(relPath, "/") {
			relPath = "/" + relPath
		}
		return relPath
	}
	rel, _ := filepath.Rel(filepath.Dir(htmlPath), assetPath)
	return rel
}
