package main

import (
	"os"
	"path/filepath"
	"testing"
)

// Test our static file discovery logic in isolation
func TestStaticFileDiscovery(t *testing.T) {
	staticDir := "public"
	
	// Check that the public directory exists
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		t.Skipf("Static directory %s does not exist", staticDir)
	}

	// Find all files in the static directory (same logic as in WriteStaticAssets)
	var files []string
	err := filepath.Walk(staticDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, err := filepath.Rel(staticDir, path)
			if err != nil {
				return err
			}
			files = append(files, relPath)
		}
		return nil
	})
	
	if err != nil {
		t.Fatalf("Error walking directory: %v", err)
	}
	
	if len(files) == 0 {
		t.Error("No files found in static directory")
	}
	
	// Verify we found index.html
	found := false
	for _, file := range files {
		if file == "index.html" {
			found = true
			break
		}
	}
	
	if !found {
		t.Error("Expected to find index.html in static files")
	}
	
	t.Logf("Found %d static files: %v", len(files), files)
}