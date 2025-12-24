package commands

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AddFiles(files []string) error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}

	karasuPath := filepath.Join(root, ".karasu")
	if _, err := os.Stat(karasuPath); os.IsNotExist(err) {
		return fmt.Errorf("not a karasu repository")
	}

	indexPath := filepath.Join(karasuPath, "index")
	// path -> hash
	existingIndex := make(map[string]string)

	// read existing index
	if indexData, err := os.ReadFile(indexPath); err == nil {
		lines := strings.Split(string(indexData), "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			parts := strings.SplitN(line, " ", 2)
			if len(parts) == 2 {
				existingIndex[parts[1]] = parts[0]
			}
		}
	}

	objectsPath := filepath.Join(karasuPath, "objects")

	// process each file
	for _, file := range files {
		absPath := filepath.Join(root, file)
		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			fmt.Printf("File does not exist: %s\n", file)
			continue
		}

		content, err := os.ReadFile(absPath)
		if err != nil {
			return fmt.Errorf("error reading file %s: %v", file, err)
		}

		hash := sha1.Sum(content)
		hashStr := fmt.Sprintf("%x", hash)

		// store object
		objDir := filepath.Join(objectsPath, hashStr[:2])
		if err := os.MkdirAll(objDir, 0755); err != nil {
			return err
		}
		objPath := filepath.Join(objDir, hashStr[2:])
		if _, err := os.Stat(objPath); os.IsNotExist(err) {
			if err := os.WriteFile(objPath, content, 0644); err != nil {
				return err
			}
		}

		// update index
		existingIndex[file] = hashStr
	}

	// write updated index
	var indexLines []string
	for path, hash := range existingIndex {
		indexLines = append(indexLines, fmt.Sprintf("%s %s", hash, path))
	}
	newIndexData := strings.Join(indexLines, "\n") + "\n"
	if err := os.WriteFile(indexPath, []byte(newIndexData), 0644); err != nil {
		return err
	}

	fmt.Printf("Added %d file(s) to index\n", len(files))
	return nil
}
