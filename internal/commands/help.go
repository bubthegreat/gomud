package commands

import (
	"strings"
	"bubtaylor.com/gomud/internal/world"

)

func executeHelp(player *world.Player) string {
	var output strings.Builder
	output.WriteString("Available commands:\n")
	for _, cmd := range ListCommands() {
		output.WriteString("- " + cmd.Name + ": " + cmd.Description + "\n")
	}
	return output.String()
}
