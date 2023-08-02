package views

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jimxshaw/go-terminal-ui/models"
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

var contacts []models.Contact

var app = tview.NewApplication()

var pages = tview.NewPages()

var form = tview.NewForm()

var contactsList = tview.NewList().ShowSecondaryText(false)

var contactTextView = tview.NewTextView()

var flex = tview.NewFlex()

var textView = tview.NewTextView().
	SetTextColor(tcell.ColorWhite).
	SetText("(a) to add a new contact \n(q) to quit")

func StartApplication() {
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
	contact := models.Contact{}

	form.AddInputField("First Name", "", fieldWidth, nil, func(firstName string) {
		contact.FirstName = firstName
	})

	form.AddInputField("Last Name", "", fieldWidth, nil, func(lastName string) {
		contact.LastName = lastName
	})

	form.AddInputField("Email", "", fieldWidth, nil, func(email string) {
		contact.Email = email
	})

	form.AddInputField("Phone Number", "", fieldWidth, nil, func(phoneNumber string) {
		contact.PhoneNumber = phoneNumber
	})

	form.AddInputField("City", "", fieldWidth, nil, func(city string) {
		contact.City = city
	})

	form.AddDropDown("State", states, 0, func(state string, index int) {
		contact.State = state
	})

	form.AddCheckbox("Business", false, func(business bool) {
		contact.Business = business
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
		contactsList.AddItem(contact.FirstName+" "+contact.LastName, " ", rune(asciiNumberOne+index), nil)
	}
}

func setDetailsText(contact *models.Contact) {
	contactTextView.Clear()
	details := contact.FirstName +
		" " +
		contact.LastName +
		"\n" +
		contact.Email +
		"\n" +
		contact.PhoneNumber +
		"\n" +
		contact.City +
		"\n" +
		contact.State
	contactTextView.SetText(details)
}
