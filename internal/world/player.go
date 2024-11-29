package world

import (
	"net" // For TCP connections, or you can use a websocket package if you're using websockets
	"fmt"
)

// Player represents a player in the game
type Player struct {
	Username    string
	Conn    net.Conn // Add this field to store the player's network connection
	RoomID	string
}


// NewPlayer returns a new Player instance
func NewPlayer(name string, conn net.Conn) *Player {

	newPlayer := &Player{
		Username:     name,
		Conn:     conn,
	}
	GlobalState.Players[newPlayer.Username] = newPlayer
	return newPlayer
}

// Write sends a message to all players in the same room as the source player.
func (player *Player) Write(message string) {
	_, err := player.Conn.Write([]byte(message + "\n"))
	if err != nil {
		// Handle the error
		fmt.Println("Error sending message to player:", err)
	}
}


// WriteRoom sends a message to all players in the same room as the source player.
func (player *Player) WriteRoom(message string) {

	room := player.GetRoom()

	for _, otherPlayer := range room.Players {
		fmt.Println("player: %v", otherPlayer)
		if otherPlayer.Username != player.Username {
			fmt.Println("Sending message to player: %v", otherPlayer)
			otherPlayer.Write(message)
		}
	}
}

// GetRoom Get the room that a player is in
func (player *Player) GetRoom() *Room {
	room := GlobalState.Rooms[player.RoomID]
	return room
}

// SetRoom Set the room a user should be in with a room ID
func (player *Player) SetRoom(roomID string) {
	// remove from old room
	oldRoom, exists := GlobalState.Rooms[player.RoomID]
	if exists {
		delete(oldRoom.Players, player.Username)
	}

	newRoom := GlobalState.Rooms[roomID]
	player.RoomID = newRoom.ID
	newRoom.Players[player.Username] = player
}