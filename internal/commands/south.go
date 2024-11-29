package commands

import (
	"bubtaylor.com/gomud/internal/world"
)

// Execute the 'south' movement
func executeSouth(player *world.Player) string {
	return move(player, "south")
}
