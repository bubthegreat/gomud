package commands

import (
	"strings"
	"bubtaylor.com/gomud/internal/world"
)

func executeLook(player *world.Player) string {
	player, exists := world.GlobalState.Players[player.Username]
	if !exists {
		return "Player not found."
	}

	room, exists := world.GlobalState.Rooms[player.RoomID]
	if !exists {
		return "You are in a void. Something went wrong..."
	}

	var output strings.Builder
	output.WriteString("--- " + room.Name + " ---\n")
	output.WriteString(room.Description + "\n")
	output.WriteString("Exits: " + strings.Join(keys(room.Exits), ", ") + "\n")
	if len(room.Items) > 0 {
		output.WriteString("Items: " + strings.Join(room.Items, ", ") + "\n")
	} else {
		output.WriteString("The room is empty.\n")
	}
	return output.String()
}

func keys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
