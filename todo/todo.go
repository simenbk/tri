package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	// Placeholder implementation
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	err = json.Unmarshal(b, &items)
	if err != nil {
		return []Item{}, err
	}
	return items, nil
}
