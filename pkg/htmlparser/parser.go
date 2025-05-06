package htmlparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Parser extracts asset references from HTML.
type Parser struct {
	extensions []string
}

// NewParser creates a new Parser with supported extensions.
func NewParser(extensions []string) *Parser {
	return &Parser{extensions: extensions}
}

// ParseHTML parses an HTML reader and returns asset references.
func (p *Parser) ParseHTML(r io.Reader) ([]string, error) {
	var assets []string
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return assets, nil
			}
			return nil, z.Err()
		case html.StartTagToken, html.SelfClosingTagToken:
			tag, hasAttr := z.TagName()
			if !hasAttr {
				continue
			}
			tagName := string(tag)
			if tagName != "link" && tagName != "script" && tagName != "img" {
				continue
			}

			var href, src string
			for hasAttr {
				key, val, more := z.TagAttr()
				if tagName == "link" && string(key) == "href" {
					href = string(val)
				} else if tagName == "script" && string(key) == "src" {
					src = string(val)
				} else if tagName == "img" && string(key) == "src" {
					src = string(val)
				}
				hasAttr = more
			}

			asset := href
			if tagName == "script" || tagName == "img" {
				asset = src
			}
			if asset == "" {
				continue
			}

			for _, ext := range p.extensions {
				if strings.HasSuffix(strings.ToLower(asset), ext) {
					assets = append(assets, asset)
					break
				}
			}
		}
	}
}
