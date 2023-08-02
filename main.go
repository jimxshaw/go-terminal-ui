package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Contact represents a contact.
type Contact struct {
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	city        string
	state       string
	occupation  string
}

var contacts []Contact

var app = tview.NewApplication()
var textView = tview.NewTextView().
	SetTextColor(tcell.ColorWhite).
	SetText("Press (q) to quit")

func main() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		}
		return event
	})

	if err := app.SetRoot(textView, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
