package tests

import (
	"testing"

	"github.com/ericklopezdev/karasu/internal/commands"
	"github.com/stretchr/testify/assert"
)

func TestStatusRepositoryNotRepo(t *testing.T) {
	setupTest(t)

	assert.NoError(t, commands.StatusRepository())
}

func TestStatusRepositoryInitialized(t *testing.T) {
	_, _ = setupTest(t)

	assert.NoError(t, commands.InitRepository())

	assert.NoError(t, commands.StatusRepository())
}
