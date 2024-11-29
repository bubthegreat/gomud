package commands

import (
	"bubtaylor.com/gomud/internal/world"
)

// Execute the 'up' movement
func executeUp(player *world.Player) string {
	return move(player, "up")
}
