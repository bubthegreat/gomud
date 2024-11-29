package world

import (
	"fmt"
	"sync"
)

// GameState Represents the general game state with rooms and players and eventually areas.
type GameState struct {
	mu      sync.Mutex
	Players map[string]*Player
	Rooms   map[string]*Room
}

var GlobalState = &GameState{
	Players: make(map[string]*Player),
	Rooms:   make(map[string]*Room),
}


// Broadcast sends a message to all players in a given room, but this should be moved to the room object.
func (gs *GameState) Broadcast(roomID string, message string) {
	fmt.Println("Attempting to broadcast some shit in room", roomID, ": ", message)
	room, exists := gs.Rooms[roomID]
	if !exists {
		
		fmt.Println("Fuckin room don't exist yo: ", roomID)
		return // If the room doesn't exist, we can't broadcast to it.
	}

	fmt.Printf("Player IDS in room: %v\n", room.Players)

	for _, player := range room.Players {
		fmt.Println("Attempting to broadcast some shit in room", roomID, " to player ", player.Username)
		_, err := player.Conn.Write([]byte(message + "\n"))
		if err != nil {
			// Handle the error
			fmt.Println("Error sending message to player:", err)
		}
	}
}