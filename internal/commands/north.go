package commands

import (
	"bubtaylor.com/gomud/internal/world"
)

// Execute the 'north' movement
func executeNorth(player *world.Player) string {
	return move(player, "north")
}
