package commands_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bubtaylor.com/gomud/internal/commands"
)


func TestCommands(t *testing.T) {
	result := commands.ListCommands()
	var expected []commands.Command = nil
	assert.Equal(t, result, expected, "The slices should be equal")
}