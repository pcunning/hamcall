package downloader

import (
	"archive/zip"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFetchHttp(t *testing.T) {
	// Create a test server
	testContent := "This is test content for HTTP download"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, testContent)
	}))
	defer server.Close()

	// Create temp directory for test
	tempDir, err := os.MkdirTemp("", "downloader_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change to temp directory
	originalDir, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)

	// Test download
	testFile := "test_download.txt"
	err = FetchHttp(testFile, server.URL)
	if err != nil {
		t.Fatalf("FetchHttp failed: %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Downloaded file does not exist")
	}

	// Verify file content
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read downloaded file: %v", err)
	}

	if string(content) != testContent {
		t.Fatalf("Expected content '%s', got '%s'", testContent, string(content))
	}
}

func TestFetchHttpInvalidURL(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "downloader_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	originalDir, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)

	// Test with invalid URL
	err = FetchHttp("test.txt", "http://invalid-url-that-does-not-exist.com")
	if err == nil {
		t.Fatalf("Expected error for invalid URL, got nil")
	}
}

func TestUnzip(t *testing.T) {
	// Create temp directory for test
	tempDir, err := os.MkdirTemp("", "downloader_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test zip file
	zipPath := filepath.Join(tempDir, "test.zip")
	destDir := filepath.Join(tempDir, "extracted")

	// Create zip file with test content
	zipFile, err := os.Create(zipPath)
	if err != nil {
		t.Fatalf("Failed to create zip file: %v", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)

	// Add a file to the zip
	testContent := "This is test content inside zip"
	fileWriter, err := zipWriter.Create("test.txt")
	if err != nil {
		t.Fatalf("Failed to create file in zip: %v", err)
	}
	_, err = fileWriter.Write([]byte(testContent))
	if err != nil {
		t.Fatalf("Failed to write to zip file: %v", err)
	}

	// Add a directory to the zip
	dirWriter, err := zipWriter.Create("testdir/")
	if err != nil {
		t.Fatalf("Failed to create directory in zip: %v", err)
	}
	_ = dirWriter

	// Add a file in the directory
	subFileWriter, err := zipWriter.Create("testdir/subfile.txt")
	if err != nil {
		t.Fatalf("Failed to create subfile in zip: %v", err)
	}
	_, err = subFileWriter.Write([]byte("Content in subdirectory"))
	if err != nil {
		t.Fatalf("Failed to write to subfile: %v", err)
	}

	zipWriter.Close()
	zipFile.Close()

	// Test extraction
	filenames, err := Unzip(zipPath, destDir)
	if err != nil {
		t.Fatalf("Unzip failed: %v", err)
	}

	// Check that files were extracted
	if len(filenames) != 3 {
		t.Fatalf("Expected 3 extracted items, got %d", len(filenames))
	}

	// Verify extracted file content
	extractedFile := filepath.Join(destDir, "test.txt")
	content, err := os.ReadFile(extractedFile)
	if err != nil {
		t.Fatalf("Failed to read extracted file: %v", err)
	}

	if string(content) != testContent {
		t.Fatalf("Expected content '%s', got '%s'", testContent, string(content))
	}

	// Verify subdirectory file
	subFile := filepath.Join(destDir, "testdir", "subfile.txt")
	subContent, err := os.ReadFile(subFile)
	if err != nil {
		t.Fatalf("Failed to read extracted subfile: %v", err)
	}

	if string(subContent) != "Content in subdirectory" {
		t.Fatalf("Expected subcontent 'Content in subdirectory', got '%s'", string(subContent))
	}

	// Verify directory exists
	testDirPath := filepath.Join(destDir, "testdir")
	if info, err := os.Stat(testDirPath); err != nil || !info.IsDir() {
		t.Fatalf("Expected directory %s to exist", testDirPath)
	}
}

func TestUnzipInvalidZip(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "downloader_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a non-zip file
	invalidZip := filepath.Join(tempDir, "invalid.zip")
	err = os.WriteFile(invalidZip, []byte("This is not a zip file"), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid zip file: %v", err)
	}

	// Test extraction of invalid zip
	_, err = Unzip(invalidZip, filepath.Join(tempDir, "dest"))
	if err == nil {
		t.Fatalf("Expected error for invalid zip file, got nil")
	}
}

func TestUnzipNonExistentFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "downloader_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test extraction of non-existent file
	_, err = Unzip(filepath.Join(tempDir, "nonexistent.zip"), filepath.Join(tempDir, "dest"))
	if err == nil {
		t.Fatalf("Expected error for non-existent zip file, got nil")
	}
}

func TestUnzipZipSlipProtection(t *testing.T) {
	// Create temp directory for test
	tempDir, err := os.MkdirTemp("", "downloader_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a zip file with path traversal attempt
	zipPath := filepath.Join(tempDir, "malicious.zip")
	destDir := filepath.Join(tempDir, "extracted")

	zipFile, err := os.Create(zipPath)
	if err != nil {
		t.Fatalf("Failed to create zip file: %v", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)

	// Add a file with path traversal attempt
	maliciousPath := "../../../etc/passwd"
	fileWriter, err := zipWriter.Create(maliciousPath)
	if err != nil {
		t.Fatalf("Failed to create malicious file in zip: %v", err)
	}
	_, err = fileWriter.Write([]byte("malicious content"))
	if err != nil {
		t.Fatalf("Failed to write malicious content: %v", err)
	}

	zipWriter.Close()
	zipFile.Close()

	// Test extraction - should fail due to path traversal protection
	_, err = Unzip(zipPath, destDir)
	if err == nil {
		t.Fatalf("Expected error for zip slip attempt, got nil")
	}

	// Error message should mention illegal file path
	if !strings.Contains(err.Error(), "illegal file path") {
		t.Fatalf("Expected 'illegal file path' error, got: %v", err)
	}
}