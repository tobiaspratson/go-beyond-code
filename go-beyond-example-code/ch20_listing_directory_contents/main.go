package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type FileInfo struct {
	Name    string
	Size    int64
	IsDir   bool
	ModTime time.Time
}

func listDirectory(dirPath string) ([]FileInfo, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", dirPath, err)
	}

	var files []FileInfo
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue // Skip files we can't get info for
		}

		files = append(files, FileInfo{
			Name:    entry.Name(),
			Size:    info.Size(),
			IsDir:   entry.IsDir(),
			ModTime: info.ModTime(),
		})
	}

	// Sort by name
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})

	return files, nil
}

func main() {
	dir := "." // Current directory
	files, err := listDirectory(dir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	fmt.Printf("Contents of %s:\n", dir)
	fmt.Printf("%-20s %-10s %-10s %s\n", "Name", "Size", "Type", "Modified")
	fmt.Println(strings.Repeat("-", 60))

	for _, file := range files {
		fileType := "File"
		if file.IsDir {
			fileType = "Directory"
		}

		fmt.Printf("%-20s %-10d %-10s %s\n",
			file.Name,
			file.Size,
			fileType,
			file.ModTime.Format("2006-01-02 15:04:05"))
	}
}
