package fileprocessor

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ryanburnette/hashref/internal/util"
	"github.com/ryanburnette/hashref/pkg/config"
	"github.com/ryanburnette/hashref/pkg/hasher"
	"github.com/ryanburnette/hashref/pkg/htmlparser"
)

// Processor handles file processing and renaming.
type Processor struct {
	parser    *htmlparser.Parser
	hasher    *hasher.Hasher
	cfg       *config.Config
	changes   []string
	changesMu sync.Mutex
	assets    map[string]string // absPath -> newPath
	assetsMu  sync.Mutex
}

// NewProcessor creates a new Processor.
func NewProcessor(parser *htmlparser.Parser, hasher *hasher.Hasher, cfg *config.Config) *Processor {
	return &Processor{
		parser: parser,
		hasher: hasher,
		cfg:    cfg,
		assets: make(map[string]string),
	}
}

// ProcessDir processes all files in the directory.
func (p *Processor) ProcessDir(dir string) error {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 10) // Limit concurrency

	// First pass: Collect assets and HTML updates
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".html" {
			return nil
		}

		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			if err := p.processHTMLFile(path); err != nil {
				log.Printf("Error processing %s: %v", path, err)
			}
		}()
		return nil
	})

	wg.Wait()
	if err != nil {
		return err
	}

	// Second pass: Rename assets
	if !p.cfg.DryRun {
		for absPath, newPath := range p.assets {
			log.Printf("Renaming asset: %s to %s", absPath, newPath)
			if err := os.Rename(absPath, newPath); err != nil {
				log.Printf("Error renaming %s to %s: %v", absPath, newPath, err)
				continue
			}
			p.changesMu.Lock()
			p.changes = append(p.changes, fmt.Sprintf("Renamed %s to %s", absPath, newPath))
			p.changesMu.Unlock()
		}
	}

	// Output changes
	for _, change := range p.changes {
		fmt.Println(change)
	}
	return nil
}

// processHTMLFile processes a single HTML file.
func (p *Processor) processHTMLFile(path string) error {
	log.Printf("Processing HTML file: %s, DryRun: %v", path, p.cfg.DryRun)
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	assets, err := p.parser.ParseHTML(f)
	if err != nil {
		return err
	}

	// Map old paths to new paths
	updates := make(map[string]string)
	for _, asset := range assets {
		absPath, err := util.ResolveAssetPath(path, asset, p.cfg.Dir)
		if err != nil {
			log.Printf("Warning: Skipping asset %s in %s: %v", asset, path, err)
			continue
		}

		p.assetsMu.Lock()
		// Check if asset is already processed
		if newPath, exists := p.assets[absPath]; exists {
			newAsset := util.RelativizePath(newPath, path, p.cfg.Dir)
			updates[asset] = newAsset
			p.assetsMu.Unlock()
			continue
		}

		// Check if asset is already hashed
		hash, err := p.hasher.HashFile(absPath)
		if err != nil {
			log.Printf("Warning: Skipping asset %s in %s: %v", asset, path, err)
			p.assetsMu.Unlock()
			continue
		}
		newPath := p.hasher.NewFilename(absPath, hash)
		if filepath.Base(absPath) == filepath.Base(newPath) {
			// Asset is already hashed
			newAsset := util.RelativizePath(absPath, path, p.cfg.Dir)
			updates[asset] = newAsset
			p.assets[absPath] = absPath
			p.assetsMu.Unlock()
			continue
		}

		// Plan rename
		p.changesMu.Lock()
		p.changes = append(p.changes, fmt.Sprintf("Would rename %s to %s", absPath, newPath))
		p.changesMu.Unlock()
		p.assets[absPath] = newPath
		newAsset := util.RelativizePath(newPath, path, p.cfg.Dir)
		updates[asset] = newAsset
		p.assetsMu.Unlock()
	}

	if len(updates) == 0 {
		log.Printf("No updates needed for %s", path)
		return nil
	}

	// Log HTML update
	p.changesMu.Lock()
	p.changes = append(p.changes, fmt.Sprintf("Would update %s with new asset paths", path))
	p.changesMu.Unlock()

	// Update HTML file
	if !p.cfg.DryRun {
		log.Printf("Updating HTML file: %s", path)
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		var buf bytes.Buffer
		if _, err := io.Copy(&buf, f); err != nil {
			return err
		}

		content := buf.String()
		for old, new := range updates {
			content = strings.ReplaceAll(content, old, new)
		}

		tmpPath := path + ".tmp"
		if err := os.WriteFile(tmpPath, []byte(content), 0644); err != nil {
			return err
		}
		if err := os.Rename(tmpPath, path); err != nil {
			return err
		}
		p.changesMu.Lock()
		p.changes = append(p.changes, fmt.Sprintf("Updated %s with new asset paths", path))
		p.changesMu.Unlock()
	}

	return nil
}
