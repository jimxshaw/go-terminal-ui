package main

import "github.com/rivo/tview"

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

func main() {
	if err := app.SetRoot(tview.NewBox(), true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
