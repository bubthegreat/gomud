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

	
	commands.RegisterCommands()
	commands.RegisterShortcuts()


	// Load the area
	area, err := world.LoadArea("./internal/areas/default.json")
	if err != nil {
		panic(err)
	}

	// Initialize global state with the loaded area
	for id, room := range area.Rooms {
		world.GlobalState.Rooms[id] = room
		fmt.Println("Rooms[%d]: %v", id, room)
	}


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


