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
func NewPlayer(username string, conn net.Conn) *Player {

	newPlayer := &Player{
		Username:   username,
		Conn:     	conn,
	}

	fmt.Printf("New player instance created for %s\n", username)

	GlobalState.Players[newPlayer.Username] = newPlayer

	fmt.Printf("New player instance added to global state for %s\n", username)


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
		fmt.Println("player: " + otherPlayer.Username)
		if otherPlayer.Username != player.Username {
			fmt.Println("Sending message to player: " + otherPlayer.Username)
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
	fmt.Printf("Trying to set room %v for player %v\n", player, roomID)
	// remove from old room
	oldRoom, exists := GlobalState.Rooms[player.RoomID]

	if exists {
		fmt.Printf("Removing player %v from room %v\n", player, oldRoom)
		delete(oldRoom.Players, player.Username)
	}

	newRoom := GlobalState.Rooms[roomID]
	fmt.Printf("Assinging player %v to room %v\n", player, newRoom)

	player.RoomID = newRoom.ID

	fmt.Printf("Player room ID updated to %s\n", newRoom.ID)

	fmt.Printf("Room %v, players: %v", newRoom, newRoom.Players)

	newRoom.Players[player.Username] = player
}