package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitRepository() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}

	karasuPath := filepath.Join(root, ".karasu")

	if _, err := os.Stat(karasuPath); err == nil {
		fmt.Println("Repository already exists")
		return nil
	}

	if err := os.MkdirAll(filepath.Join(karasuPath, "objects"), 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(karasuPath, "refs", "heads"), 0755); err != nil {
		return err
	}

	// HEAD -> current branch
	headContent := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(filepath.Join(karasuPath, "HEAD"), headContent, 0644); err != nil {
		return err
	}

	// Create empty main branch
	if err := os.WriteFile(
		filepath.Join(karasuPath, "refs", "heads", "main"),
		[]byte{},
		0644,
	); err != nil {
		return err
	}

	// Create empty index
	if err := os.WriteFile(
		filepath.Join(karasuPath, "index"),
		[]byte{},
		0644,
	); err != nil {
		return err
	}

	printInitSuccess(karasuPath)
	return nil
}

func printInitSuccess(path string) {
	fmt.Println("    __")
	fmt.Println("   /_ 0>  からす - repository initialized!")
	fmt.Println("   `- '   at:", path)
}
