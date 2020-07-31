package commands_test

import (
	"fmt"
	"github.com/cyberfox1002/groo/commands"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenRemote(t *testing.T) {
	// error: invalid remote name
	err := commands.OpenRemote("missing")
	fmt.Printf("Error: %s", err)
	assert.Error(t, err)

	// success: available remote name
	err = commands.OpenRemote("origin")
	assert.NoError(t, err)
}
