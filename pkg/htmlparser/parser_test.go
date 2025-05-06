package htmlparser

import (
	"strings"
	"testing"
)

func TestParseHTML(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		extensions []string
		want       []string
	}{
		{
			name: "Basic CSS and JS",
			input: `
			<html>
			<head>
				<link href="styles.css" rel="stylesheet">
			</head>
			<body>
				<script src="app.js"></script>
			</body>
			</html>`,
			extensions: []string{".css", ".js"},
			want:       []string{"styles.css", "app.js"},
		},
		{
			name: "Additional extensions",
			input: `
			<html>
			<body>
				<img src="logo.png">
			</body>
			</html>`,
			extensions: []string{".png"},
			want:       []string{"logo.png"},
		},
		{
			name: "Ignore inline tags",
			input: `
			<html>
			<head>
				<style>body {}</style>
			</head>
			<body>
				<script>console.log('hi');</script>
			</body>
			</html>`,
			extensions: []string{".css", ".js"},
			want:       []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.extensions)
			assets, err := p.ParseHTML(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("ParseHTML error: %v", err)
			}
			if len(assets) != len(tt.want) {
				t.Errorf("Got %v, want %v", assets, tt.want)
			}
			for i, asset := range assets {
				if asset != tt.want[i] {
					t.Errorf("Got asset %v, want %v", asset, tt.want[i])
				}
			}
		})
	}
}
