package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Functions struct {
	Container fyne.CanvasObject
	functions []string
	// Add other fields needed
}

func NewFunctions() *Functions {
	functions := []string{
		"Search the web",
		"Search the web",
		"Search the web",
		"Search the web",
		"Search the web",
	}

	// Create a list widget for the functions
	list := widget.NewList(
		func() int { return len(functions) },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(functions[id])
		},
	)

	return &Functions{
		functions: functions,
		Container: list,
	}
}
