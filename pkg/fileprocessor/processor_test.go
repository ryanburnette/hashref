package fileprocessor

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ryanburnette/hashref/pkg/config"
	"github.com/ryanburnette/hashref/pkg/hasher"
	"github.com/ryanburnette/hashref/pkg/htmlparser"
)

func TestProcessDir(t *testing.T) {
	tmpDir := t.TempDir()

	// Create demo files
	htmlContent := `
	<html>
	<head>
		<link href="assets/styles.css" rel="stylesheet">
	</head>
	<body>
		<script src="assets/app.js"></script>
	</body>
	</html>`
	if err := os.Mkdir(filepath.Join(tmpDir, "assets"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "index.html"), []byte(htmlContent), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "assets/styles.css"), []byte("body {}"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "assets/app.js"), []byte("console.log('hi');"), 0644); err != nil {
		t.Fatal(err)
	}

	cfg := &config.Config{
		Dir:        tmpDir,
		DryRun:     true,
		HashLength: 8,
		Extensions: []string{".css", ".js"},
	}
	parser := htmlparser.NewParser(cfg.Extensions)
	hasher := hasher.NewMD5Hasher(cfg.HashLength)
	processor := NewProcessor(parser, hasher, cfg)

	if err := processor.ProcessDir(tmpDir); err != nil {
		t.Fatalf("ProcessDir error: %v", err)
	}

	// Check changes log
	if len(processor.changes) < 2 {
		t.Errorf("Expected at least 2 changes, got %d", len(processor.changes))
	}
	for _, change := range processor.changes {
		if !strings.Contains(change, "Would rename") && !strings.Contains(change, "Would update") {
			t.Errorf("Unexpected change: %s", change)
		}
	}
}
