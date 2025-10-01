package downloader

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/kothar/go-backblaze.v0"
)

// BackupDownloader handles backup operations with B2
type BackupDownloader struct {
	bucket *backblaze.Bucket
	path   string
}

// NewBackupDownloader creates a new backup downloader
func NewBackupDownloader(keyID, applicationKey, backupPath string) (*BackupDownloader, error) {
	if keyID == "" || applicationKey == "" || backupPath == "" {
		return nil, fmt.Errorf("backup credentials or path not configured")
	}

	b, err := backblaze.NewB2(backblaze.Credentials{
		KeyID:          keyID,
		ApplicationKey: applicationKey,
	})
	if err != nil {
		return nil, err
	}

	bucket, err := b.Bucket("hamcall")
	if err != nil {
		return nil, err
	}

	return &BackupDownloader{
		bucket: bucket,
		path:   backupPath,
	}, nil
}

// UploadBackup uploads a file to the backup location
func (bd *BackupDownloader) UploadBackup(localFile, remoteFile string) error {
	file, err := os.Open(localFile)
	if err != nil {
		return err
	}
	defer file.Close()

	backupPath := filepath.Join(bd.path, remoteFile)
	_, err = bd.bucket.UploadFile(backupPath, nil, file)
	return err
}

// DownloadBackup downloads a file from the backup location
func (bd *BackupDownloader) DownloadBackup(remoteFile, localFile string) error {
	backupPath := filepath.Join(bd.path, remoteFile)
	
	_, reader, err := bd.bucket.DownloadFileByName(backupPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	out, err := os.Create(localFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, reader)
	return err
}

// FetchWithBackup tries to fetch from primary URL, falls back to backup, and updates backup on success
func FetchWithBackup(localFile, primaryURL, backupFile string, backup *BackupDownloader) error {
	// Try primary download first
	err := FetchHttp(localFile, primaryURL)
	if err == nil {
		// Primary succeeded, upload to backup if backup system is available
		if backup != nil {
			if uploadErr := backup.UploadBackup(localFile, backupFile); uploadErr != nil {
				fmt.Printf("Warning: Failed to upload %s to backup: %v\n", backupFile, uploadErr)
			} else {
				fmt.Printf("Successfully backed up %s to B2\n", backupFile)
			}
		}
		return nil
	}

	// Primary failed, try backup if available
	fmt.Printf("Primary download failed for %s: %v\n", localFile, err)
	if backup != nil {
		fmt.Printf("Attempting to download %s from backup...\n", localFile)
		if backupErr := backup.DownloadBackup(backupFile, localFile); backupErr == nil {
			fmt.Printf("Successfully restored %s from backup\n", localFile)
			return nil
		} else {
			fmt.Printf("Backup download also failed for %s: %v\n", localFile, backupErr)
			return backupErr
		}
	}

	// Primary failed and no backup system available
	return err
}