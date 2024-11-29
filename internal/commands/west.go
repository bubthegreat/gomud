package commands

import (
	"bubtaylor.com/gomud/internal/world"
)

// Execute the 'west' movement
func executeWest(player *world.Player) string {
	return move(player, "west")
}
