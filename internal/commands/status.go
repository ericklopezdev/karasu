package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func StatusRepository() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}

	karasuPath := filepath.Join(root, ".karasu")

	if _, err := os.Stat(karasuPath); os.IsNotExist(err) {
		fmt.Println("Not a karasu repository")
		return nil
	}

	// read HEAD
	headPath := filepath.Join(karasuPath, "HEAD")
	headContent, err := os.ReadFile(headPath)
	if err != nil {
		return err
	}

	currentBranch := string(headContent)
	if len(currentBranch) > 0 && currentBranch[len(currentBranch)-1] == '\n' {
		currentBranch = currentBranch[:len(currentBranch)-1]
	}

	// extract branch name
	branchName := "unknown"
	if len(currentBranch) > 5 && currentBranch[:5] == "ref: " {
		branchName = currentBranch[5:]
		if len(branchName) > 10 && branchName[:10] == "refs/heads/" {
			branchName = branchName[10:]
		}
	}

	fmt.Printf("On branch %s\n", branchName)

	// read branch ref
	branchRefPath := filepath.Join(karasuPath, currentBranch[5:]) // remove "ref: "
	if _, err := os.Stat(branchRefPath); os.IsNotExist(err) {
		fmt.Println("No commits yet")
	} else {
		commitHash, err := os.ReadFile(branchRefPath)
		if err != nil {
			return err
		}
		if len(commitHash) == 0 {
			fmt.Println("No commits yet")
		} else {
			fmt.Printf("Latest commit: %s\n", string(commitHash))
		}
	}

	// read index and show staged files
	indexPath := filepath.Join(karasuPath, "index")
	if indexData, err := os.ReadFile(indexPath); err == nil && len(indexData) > 0 {
		fmt.Println("\nChanges to be committed:")
		lines := strings.Split(string(indexData), "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			parts := strings.SplitN(line, " ", 2)
			if len(parts) == 2 {
				fmt.Printf("  new file: %s\n", parts[1])
			}
		}
	}

	return nil
}
