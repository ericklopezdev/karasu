package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ericklopezdev/karasu/internal/commands"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (string, string) {
	originalWD, _ := os.Getwd()

	// create and switch to a temporary directory
	tempDir := t.TempDir()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("setup failed: could not change directory to tempDir: %v", err)
	}

	// ensure to return to the original working directory
	t.Cleanup(func() {
		os.Chdir(originalWD)
	})

	karasuPath := filepath.Join(tempDir, ".karasu")
	return tempDir, karasuPath
}

func TestInitSuccess(t *testing.T) {
	_, karasuPath := setupTest(t)

	assert.NoError(t, commands.InitRepository(), "InitRepository should not return an error on a fresh directory")

	assert.DirExists(t, karasuPath, ".karasu directory should be created")

	// check required subdirectories
	subdirs := []string{"objects", filepath.Join("refs", "heads")}
	for _, dir := range subdirs {
		path := filepath.Join(karasuPath, dir)
		assert.DirExists(t, path, "Required subdirectory %s was not created", dir)
	}

	// check HEAD file content
	headPath := filepath.Join(karasuPath, "HEAD")
	content, err := os.ReadFile(headPath)

	assert.NoError(t, err, "HEAD file should exist and be readable")
	expectedContent := "ref: refs/heads/main\n"
	assert.Equal(t, expectedContent, string(content), "HEAD content must point to 'ref: refs/heads/main\\n'")

	// check for index and main ref files
	files := []string{"index", filepath.Join("refs", "heads", "main")}
	for _, file := range files {
		path := filepath.Join(karasuPath, file)
		assert.FileExists(t, path, "Required file %s was not created", file)
	}
}

func TestInitAlreadyExists(t *testing.T) {
	_, preexistingKarasuPath := setupTest(t)

	// manually create the .karasu directory first
	if err := os.Mkdir(preexistingKarasuPath, 0755); err != nil {
		t.Fatalf("setup failed: could not create .karasu: %v", err)
	}

	// add a file to verify it's not deleted/overwritten
	testFilePath := filepath.Join(preexistingKarasuPath, "safety_file.txt")
	uniqueContent := []byte("do not touch me")
	if err := os.WriteFile(testFilePath, uniqueContent, 0644); err != nil {
		t.Fatalf("setup failed: could not create safety file: %v", err)
	}

	assert.NoError(t, commands.InitRepository(), "InitRepository call on existing repo should return nil")

	assert.FileExists(t, testFilePath, "Pre-existing repository files should not be deleted")

	content, err := os.ReadFile(testFilePath)
	assert.NoError(t, err)
	assert.Equal(t, uniqueContent, content, "Pre-existing file content should be preserved")
}
