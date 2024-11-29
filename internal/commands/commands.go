package commands

import (
	"fmt"
	"sync"
	"bubtaylor.com/gomud/internal/world"
)

type CommandFunc func(player *world.Player) string

type Command struct {
	Name        string
	Description string
	Execute     CommandFunc
}

var (
	commands = make(map[string]*Command)
	mu       sync.RWMutex
)

var commandHandlers = make(map[string]func(*world.Player) string)
var shortcutHandlers = make(map[string]func(*world.Player) string)

// Register the main command and its handler
func registerCommand(command string, handler func(*world.Player) string) {
	if _, exists := commandHandlers[command]; exists {
		fmt.Printf("Command %s is already registered.\n", command)
		return
	}
	commandHandlers[command] = handler
	fmt.Printf("Registered command: %s\n", command)
}

// Register a shortcut for a command
func registerShortcut(shortcut string, handler func(*world.Player) string) {
	if _, exists := shortcutHandlers[shortcut]; exists {
		fmt.Printf("Shortcut %s is already registered.\n", shortcut)
		return
	}
	shortcutHandlers[shortcut] = handler
	fmt.Printf("Registered shortcut: %s\n", shortcut)
}

// Execute a command or a shortcut (calls the associated handler).
// It first looks for the shortcut, and if not found, looks for the command.
func HandleCommand(player *world.Player, input string) string {
	// First, check if it's a shortcut
	if handler, exists := shortcutHandlers[input]; exists {
		fmt.Println("Command entered for player ", player.Username, ": ", input)
		return handler(player)
	}

	// If not a shortcut, check if it's a full command
	if handler, exists := commandHandlers[input]; exists {
		fmt.Println("Shortcut entered for player ", player.Username, ": ", input)
		return handler(player)
	}

	// If neither is found, return an error message
	return fmt.Sprintf("Command or shortcut %s not found.", input)
}

// GetCommand retrieves a command by name
func GetCommand(name string) *Command {
	mu.RLock()
	defer mu.RUnlock()
	return commands[name]
}

// ListCommands returns all available commands
func ListCommands() []Command {
	mu.RLock()
	defer mu.RUnlock()
	var list []Command
	for _, cmd := range commands {
		list = append(list, *cmd)
	}
	return list
}

// move function handles the actual movement logic
func move(player *world.Player, direction string) string {
	fmt.Println("Player ", player.Username, " trying to go ", direction)

	oldRoom := player.GetRoom()

	nextRoomID, exists := oldRoom.Exits[direction]
	if !exists {
		return fmt.Sprintf("You can't go that way: %s.", direction)
	}

	player.SetRoom(nextRoomID)
	// Tell the old room he's gone.
	world.GlobalState.Broadcast(oldRoom.ID, player.Username + " leaves " + direction)

	// Broadcast to players in the next room that the player is entering
	player.WriteRoom(player.Username + " has arrived.")



	// TODO:Stop using the return of this to decide to print it to the user.
	//      That shouldn't happen as a consequence of the return, it should
	//		happen as a part of the function.
	nextRoom := world.GlobalState.Rooms[nextRoomID]
	return nextRoom.Describe()
}

// Register all commands
func RegisterCommands() {
	registerCommand("down", executeDown)
	registerCommand("east", executeEast)
	registerCommand("help", executeHelp)
	registerCommand("look", executeLook)
	registerCommand("north", executeNorth)
	registerCommand("south", executeSouth)
	registerCommand("up", executeUp)
	registerCommand("west", executeWest)
}

// Register shortcuts
func RegisterShortcuts() {
	registerShortcut("d", executeDown)
	registerShortcut("e", executeEast)
	registerShortcut("h", executeHelp)
	registerShortcut("l", executeLook)
	registerShortcut("n", executeNorth)
	registerShortcut("s", executeSouth)
	registerShortcut("u", executeUp)
	registerShortcut("w", executeWest)
}