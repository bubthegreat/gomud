package world

import (
	"encoding/json"
	"os"
)



// Area represents a set of rooms
type Area struct {
	Rooms map[string]*Room `json:"rooms"`
}


// LoadArea loads areas from a config file.
func LoadArea(filename string) (Area, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Area{}, err
	}
	defer file.Close()

	var area Area
	err = json.NewDecoder(file).Decode(&area)
	if err != nil {
		return Area{}, err
	}
	return area, nil
}