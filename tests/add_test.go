package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ericklopezdev/karasu/internal/commands"
	"github.com/stretchr/testify/assert"
)

func TestAddFilesSuccess(t *testing.T) {
	tempDir, karasuPath := setupTest(t)

	assert.NoError(t, commands.InitRepository())

	testFilePath := filepath.Join(tempDir, "test.txt")
	testContent := []byte("hello world")
	assert.NoError(t, os.WriteFile(testFilePath, testContent, 0644))

	assert.NoError(t, commands.AddFiles([]string{"test.txt"}))

	indexPath := filepath.Join(karasuPath, "index")
	assert.FileExists(t, indexPath)

	indexData, err := os.ReadFile(indexPath)
	assert.NoError(t, err)
	assert.Contains(t, string(indexData), "test.txt")
}

func TestAddFilesNonExistent(t *testing.T) {
	_, _ = setupTest(t)

	assert.NoError(t, commands.InitRepository())

	assert.NoError(t, commands.AddFiles([]string{"nonexistent.txt"}))
}

func TestAddFilesNotRepo(t *testing.T) {
	_, _ = setupTest(t)

	assert.Error(t, commands.AddFiles([]string{"test.txt"}))
}
