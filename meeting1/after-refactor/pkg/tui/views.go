package tui

import (
	"fmt"

	"devbook/pkg/devbook"
)

type ItemStore interface {
	Add(devbook.Item)
	GetByName(string) (devbook.Item, bool)
	DeleteByName(string)
	List() []devbook.Item
}

func addItemView(store ItemStore) error {
	name, err := AskForString("Name")
	if err != nil {
		return err
	}

	description, err := AskForString("Description")
	if err != nil {
		return err
	}

	url, err := AskForString("URL")
	if err != nil {
		return err
	}

	store.Add(devbook.Item{
		Name:        name,
		Description: description,
		URL:         url,
	})
	return nil
}

func deleteItemView(store ItemStore) error {
	name, err := AskForString("Name")
	if err != nil {
		return err
	}

	if _, ok := store.GetByName(name); !ok {
		fmt.Println("Item does not exist")
		return nil
	}

	store.DeleteByName(name)
	return nil

	// path?a=1&a=2

}

func printItems(store ItemStore) {
	for _, item := range store.List() {
		fmt.Printf("%s : %s : %s\n", item.Name, item.URL, item.Description)
	}
}

func MenuView(store ItemStore) error {
	choices := map[string]string{
		"a": "Add",
		"e": "Edit",
		"d": "Delete",
		"l": "List",
		"q": "Exit",
	}

	for {
		choice, err := AskForMapChoice("Choose option", choices)
		if err != nil {
			return err
		}

		switch choice {
		case "a":
			if err = addItemView(store); err != nil {
				return err
			}
		case "e":
			fmt.Println("Not yet implemented")
		case "d":
			if err = deleteItemView(store); err != nil {
				return err
			}
		case "l":
			printItems(store)
		case "q":
			return nil
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println()
	}
}
