package main

import (
	"devbook/pkg/devbook"
	"devbook/pkg/tui"
)

func main() {
	store := devbook.Store{}

	store.Add(devbook.Item{
		Name:        "item1",
		Description: "search engine",
		URL:         "https://google.pl",
	})
	store.Add(devbook.Item{
		Name:        "item2",
		Description: "other search engine",
		URL:         "https://bing.com",
	})

	if err := tui.MenuView(store); err != nil {
		panic(err)
	}
}
