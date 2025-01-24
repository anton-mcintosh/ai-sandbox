package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Functions struct {
	Container fyne.CanvasObject
	functions []string
}

func NewFunctions() *Functions {
	functions := []string{
    "Example func 1",
    "Example func 2",
    "Example func 3",
	}

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
