package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type BackupOptions struct {
	SourceDir    string
	BackupDir    string
	ExcludeDirs  []string
	ExcludeFiles []string
	DryRun       bool
	Verbose      bool
	Force        bool
}

type FileInfo struct {
	Path     string
	Size     int64
	ModTime  time.Time
	Checksum string
}

type BackupManager struct {
	options     BackupOptions
	sourceFiles map[string]FileInfo
	backupFiles map[string]FileInfo
}

func NewBackupManager(options BackupOptions) *BackupManager {
	return &BackupManager{
		options:     options,
		sourceFiles: make(map[string]FileInfo),
		backupFiles: make(map[string]FileInfo),
	}
}

func (bm *BackupManager) calculateChecksum(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func (bm *BackupManager) shouldExclude(path string) bool {
	// Check directory exclusions
	for _, excludeDir := range bm.options.ExcludeDirs {
		if strings.Contains(path, excludeDir) {
			return true
		}
	}

	// Check file exclusions
	for _, excludeFile := range bm.options.ExcludeFiles {
		if strings.Contains(filepath.Base(path), excludeFile) {
			return true
		}
	}

	return false
}

func (bm *BackupManager) scanDirectory(dir string, files map[string]FileInfo) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if bm.options.Verbose {
				fmt.Printf("Warning: %v\n", err)
			}
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if bm.shouldExclude(path) {
			if bm.options.Verbose {
				fmt.Printf("Excluding: %s\n", path)
			}
			return nil
		}

		checksum, err := bm.calculateChecksum(path)
		if err != nil {
			if bm.options.Verbose {
				fmt.Printf("Warning: failed to calculate checksum for %s: %v\n", path, err)
			}
			return nil
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		files[relPath] = FileInfo{
			Path:     path,
			Size:     info.Size(),
			ModTime:  info.ModTime(),
			Checksum: checksum,
		}

		return nil
	})
}

func (bm *BackupManager) needsBackup(sourceFile, backupFile FileInfo) bool {
	if bm.options.Force {
		return true
	}

	// Check if backup file doesn't exist
	if backupFile.Path == "" {
		return true
	}

	// Check if sizes are different
	if sourceFile.Size != backupFile.Size {
		return true
	}

	// Check if modification times are different
	if !sourceFile.ModTime.Equal(backupFile.ModTime) {
		return true
	}

	// Check if checksums are different
	if sourceFile.Checksum != backupFile.Checksum {
		return true
	}

	return false
}

func (bm *BackupManager) copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create destination directory if it doesn't exist
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}

func (bm *BackupManager) performBackup() error {
	// Scan source directory
	if bm.options.Verbose {
		fmt.Printf("Scanning source directory: %s\n", bm.options.SourceDir)
	}
	err := bm.scanDirectory(bm.options.SourceDir, bm.sourceFiles)
	if err != nil {
		return fmt.Errorf("failed to scan source directory: %w", err)
	}

	// Scan backup directory
	if bm.options.Verbose {
		fmt.Printf("Scanning backup directory: %s\n", bm.options.BackupDir)
	}
	err = bm.scanDirectory(bm.options.BackupDir, bm.backupFiles)
	if err != nil {
		return fmt.Errorf("failed to scan backup directory: %w", err)
	}

	// Determine files to backup
	var filesToBackup []string
	var filesToUpdate []string
	var filesToDelete []string

	for relPath, sourceFile := range bm.sourceFiles {
		backupFile, exists := bm.backupFiles[relPath]

		if !exists {
			filesToBackup = append(filesToBackup, relPath)
		} else if bm.needsBackup(sourceFile, backupFile) {
			filesToUpdate = append(filesToUpdate, relPath)
		}
	}

	// Find files to delete (files in backup but not in source)
	for relPath := range bm.backupFiles {
		if _, exists := bm.sourceFiles[relPath]; !exists {
			filesToDelete = append(filesToDelete, relPath)
		}
	}

	// Print summary
	fmt.Printf("Backup Summary:\n")
	fmt.Printf("  Files to backup: %d\n", len(filesToBackup))
	fmt.Printf("  Files to update: %d\n", len(filesToUpdate))
	fmt.Printf("  Files to delete: %d\n", len(filesToDelete))
	fmt.Printf("  Total operations: %d\n\n", len(filesToBackup)+len(filesToUpdate)+len(filesToDelete))

	if bm.options.DryRun {
		fmt.Println("Dry run - no files were actually copied or deleted.")
		return nil
	}

	// Perform backup operations
	for _, relPath := range filesToBackup {
		srcPath := filepath.Join(bm.options.SourceDir, relPath)
		dstPath := filepath.Join(bm.options.BackupDir, relPath)

		if bm.options.Verbose {
			fmt.Printf("Backing up: %s -> %s\n", srcPath, dstPath)
		}

		err := bm.copyFile(srcPath, dstPath)
		if err != nil {
			fmt.Printf("Error backing up %s: %v\n", relPath, err)
		}
	}

	for _, relPath := range filesToUpdate {
		srcPath := filepath.Join(bm.options.SourceDir, relPath)
		dstPath := filepath.Join(bm.options.BackupDir, relPath)

		if bm.options.Verbose {
			fmt.Printf("Updating: %s -> %s\n", srcPath, dstPath)
		}

		err := bm.copyFile(srcPath, dstPath)
		if err != nil {
			fmt.Printf("Error updating %s: %v\n", relPath, err)
		}
	}

	for _, relPath := range filesToDelete {
		dstPath := filepath.Join(bm.options.BackupDir, relPath)

		if bm.options.Verbose {
			fmt.Printf("Deleting: %s\n", dstPath)
		}

		err := os.Remove(dstPath)
		if err != nil {
			fmt.Printf("Error deleting %s: %v\n", relPath, err)
		}
	}

	return nil
}

func main() {
	var options BackupOptions
	var excludeDirsStr, excludeFilesStr string

	flag.StringVar(&options.SourceDir, "source", "", "Source directory to backup")
	flag.StringVar(&options.BackupDir, "backup", "", "Backup directory")
	flag.StringVar(&excludeDirsStr, "exclude-dirs", "", "Comma-separated directories to exclude")
	flag.StringVar(&excludeFilesStr, "exclude-files", "", "Comma-separated file patterns to exclude")
	flag.BoolVar(&options.DryRun, "dry-run", false, "Show what would be done without actually doing it")
	flag.BoolVar(&options.Verbose, "verbose", false, "Verbose output")
	flag.BoolVar(&options.Force, "force", false, "Force backup of all files")

	flag.Parse()

	if options.SourceDir == "" || options.BackupDir == "" {
		fmt.Println("Usage: go run main.go -source <source_dir> -backup <backup_dir> [options]")
		flag.PrintDefaults()
		return
	}

	if excludeDirsStr != "" {
		options.ExcludeDirs = strings.Split(excludeDirsStr, ",")
	}
	if excludeFilesStr != "" {
		options.ExcludeFiles = strings.Split(excludeFilesStr, ",")
	}

	manager := NewBackupManager(options)

	start := time.Now()
	err := manager.performBackup()
	if err != nil {
		fmt.Printf("Error performing backup: %v\n", err)
		return
	}
	duration := time.Since(start)

	fmt.Printf("Backup completed in %v\n", duration)
}
