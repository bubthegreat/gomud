package world


// Room represents a room in the game
type Room struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Exits       map[string]string `json:"exits"` // e.g., {"north": "room2"}
	Items       []string          `json:"items"`
	Doors       map[string]bool   `json:"doors"` // e.g., {"north": true} for open doors
	Players  	map[string]*Player `json:"players"`
}

// Describe generates a string representation of the room
func (r *Room) Describe() string {
	description := "You are in " + r.Name + ".\n"
	description += "Exits: "
	for direction, _ := range r.Exits {
		description += direction + " "
	}
	description += "\n"
	// You can also add more details, like items or NPCs in the room, if desired.
	return description
}

func (gs GameState) NewRoom(id, name string) *Room {
	return &Room{
		ID:          id,
		Name:        name,
		Description: "A generic room.",
		Exits:       make(map[string]string),
		Players:     make(map[string]*Player),
	}
}
