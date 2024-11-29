package commands

import (
	"bubtaylor.com/gomud/internal/world"
)


// Execute the 'down' movement
func executeDown(player *world.Player) string {
	return move(player, "down")
}
