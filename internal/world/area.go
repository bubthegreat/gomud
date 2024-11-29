package world

import (
	"encoding/json"
	"os"
	"fmt"
)



// Area represents a set of rooms
type Area struct {
	Name string `json:"name"`
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
	fmt.Println("Area: ", area)
	return area, nil
}