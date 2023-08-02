package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	fieldWidth     int    = 30
	addKey         rune   = 97
	quitKey        rune   = 113
	menu           string = "Menu"
	addNewContact  string = "Add New Contact"
	asciiNumberOne int    = 49
)

var states = []string{"AK", "AL", "AR", "AZ", "CA", "CO", "CT", "DC", "DE", "FL", "GA",
	"HI", "IA", "ID", "IL", "IN", "KS", "KY", "LA", "MA", "MD", "ME",
	"MI", "MN", "MO", "MS", "MT", "NC", "ND", "NE", "NH", "NJ", "NM",
	"NV", "NY", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX",
	"UT", "VA", "VT", "WA", "WI", "WV", "WY"}

// Contact represents a contact.
type Contact struct {
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	city        string
	state       string
	business    bool
}

var contacts []Contact

var app = tview.NewApplication()

var pages = tview.NewPages()

var form = tview.NewForm()

var contactsList = tview.NewList().ShowSecondaryText(false)

var contactTextView = tview.NewTextView()

var flex = tview.NewFlex()

var textView = tview.NewTextView().
	SetTextColor(tcell.ColorWhite).
	SetText("(a) to add a new contact \n(q) to quit")

func main() {
	contactsList.SetSelectedFunc(func(index int, name string, secondName string, shortcut rune) {
		setDetailsText(&contacts[index])
	})

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(contactsList, 0, 1, true).
			AddItem(contactTextView, 0, 4, false),
			0, 6, false,
		).
		AddItem(textView, 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == quitKey {
			app.Stop()
		} else if event.Rune() == addKey {
			form.Clear(true)
			addNewContactForm()
			pages.SwitchToPage(addNewContact)
		}
		return event
	})

	pages.AddPage(menu, flex, true, true)
	pages.AddPage(addNewContact, form, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func addNewContactForm() *tview.Form {
	contact := Contact{}

	form.AddInputField("First Name", "", fieldWidth, nil, func(firstName string) {
		contact.firstName = firstName
	})

	form.AddInputField("Last Name", "", fieldWidth, nil, func(lastName string) {
		contact.lastName = lastName
	})

	form.AddInputField("Email", "", fieldWidth, nil, func(email string) {
		contact.email = email
	})

	form.AddInputField("Phone Number", "", fieldWidth, nil, func(phoneNumber string) {
		contact.phoneNumber = phoneNumber
	})

	form.AddInputField("City", "", fieldWidth, nil, func(city string) {
		contact.city = city
	})

	form.AddDropDown("State", states, 0, func(state string, index int) {
		contact.state = state
	})

	form.AddCheckbox("Business", false, func(business bool) {
		contact.business = business
	})

	form.AddButton("Save", func() {
		contacts = append(contacts, contact)
		addContactList()
		pages.SwitchToPage(menu)
	})

	return form
}

func addContactList() {
	contactsList.Clear()

	for index, contact := range contacts {
		contactsList.AddItem(contact.firstName+" "+contact.lastName, " ", rune(asciiNumberOne+index), nil)
	}
}

func setDetailsText(contact *Contact) {
	contactTextView.Clear()
	details := contact.firstName +
		" " +
		contact.lastName +
		"\n" +
		contact.email +
		"\n" +
		contact.phoneNumber +
		"\n" +
		contact.city +
		"\n" +
		contact.state
	contactTextView.SetText(details)
}
