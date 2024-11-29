package main

import (
	"bufio"
	"fmt"
	"bubtaylor.com/gomud/internal/commands"
	"bubtaylor.com/gomud/internal/world"
	"net"
	"strings"
)

// main runs main program
func main() {
	// Load the area
	// area, err := world.LoadArea("assets/areas/sample_area.json")
	// if err != nil {
	// 	panic(err)
	// }

	// // Initialize global state with the loaded area
	// for id, room := range area.Rooms {
	// 	world.GlobalState.Rooms[id] = room
	// }

	commands.RegisterCommands()
	commands.RegisterShortcuts()

	initializeRooms()

	fmt.Println("Starting MUD server...")
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Server is running on port 4000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handlePlayerConnection(conn)
	}
}

// handlePlayerConnection handles a single player's login and interaction
func handlePlayerConnection(conn net.Conn) {
	defer conn.Close()

	// Read input from the player
	scanner := bufio.NewScanner(conn)

	// Step 1: Ask for username
	conn.Write([]byte("Enter your username: "))
	scanner.Scan()
	username := scanner.Text()

	// Step 2: Check if the username exists
	player, exists := world.GlobalState.Players[username]
	if !exists {
		conn.Write([]byte("No such player exists.\n"))
	}

	player = world.NewPlayer(username, conn)
	player.SetRoom("1")

	// Step 4: Welcome player and show the current room description
	player.Conn.Write([]byte(fmt.Sprintf("Welcome, %s!\n", player.Username)))
	player.Conn.Write([]byte("You are in the following room:\n"))
	playerRoom, exists := world.GlobalState.Rooms[player.RoomID]
	if !exists {
		player.Conn.Write([]byte("You are in a void. Something went wrong...\n"))
	}
	player.Conn.Write([]byte(playerRoom.Describe() + "\n"))

	// Step 5: Main game loop
	for {
		player.Conn.Write([]byte("> "))
		scanner.Scan()
		command := scanner.Text()

		// Check if the player typed 'quit' to exit the game
		if strings.ToLower(command) == "quit" {
			player.Conn.Write([]byte("Goodbye!\n"))
			return
		}

		// Process the command
		response := commands.HandleCommand(player, command)
		player.Conn.Write([]byte(response + "\n"))

	}
}

// initializeRooms sets up example rooms and exits
func initializeRooms() {
	// Create some rooms for players to move between
	room1 := world.NewRoom("1", "Starting Room")
	room1.Description = "This is the starting room. There are exits to the north and east."
	room1.Exits = map[string]string{
		"north": "2",
		"east":  "3",
	}

	room2 := world.NewRoom("2", "North Room")
	room2.Description = "This is the north room. You can go back to the south."
	room2.Exits = map[string]string{
		"south": "1",
	}

	room3 := world.NewRoom("3", "East Room")
	room3.Description = "This is the east room. You can go back to the west."
	room3.Exits = map[string]string{
		"west": "1",
	}

	// Add rooms to the game state
	world.GlobalState.Rooms["1"] = room1
	world.GlobalState.Rooms["2"] = room2
	world.GlobalState.Rooms["3"] = room3

}
