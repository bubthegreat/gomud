package commands

import (
	"bubtaylor.com/gomud/internal/world"
)

// Execute the 'east' movement
func executeEast(player *world.Player) string {
	return move(player, "east")
}
