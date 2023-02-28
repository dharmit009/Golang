package main

import (
	"github.com/fyne-io/fyne/v2"
	"github.com/fyne-io/fyne/v2/app"
	"github.com/fyne-io/fyne/v2/container"
	"github.com/fyne-io/fyne/v2/dialog"
	"github.com/fyne-io/fyne/v2/widget"
)

type Credential struct {
	Name     string
	Username string
	Password string
}

type Credentials struct {
	List []Credential
}

func main() {
	a := app.New()
	w := a.NewWindow("Password Manager")

	// Create a form for entering credential details
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter name for credential")
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Enter username for credential")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password for credential")

	form := &widget.Form{
		OnCancel: func() {
			nameEntry.SetText("")
			usernameEntry.SetText("")
			passwordEntry.SetText("")
		},
		OnSubmit: func() {
			name := nameEntry.Text
			username := usernameEntry.Text
			password := passwordEntry.Text

			// Generate a unique ID for the credential
			id := generateID()

			// Create the credential and add it to the list
			credential := Credential{
				Name:     name,
				Username: username,
				Password: password,
			}
			credentials := loadCredentials()
			credentials.List = append(credentials.List, credential)
			saveCredentials(credentials)

			// Notify the user that the credential was added
			dialog.NewInformation("Credential added", "The credential was successfully added.", w).Show()
			nameEntry.SetText("")
			usernameEntry.SetText("")
			passwordEntry.SetText("")
		},
		Items: []*widget.FormItem{
			{Text: "Name", Widget: nameEntry},
			{Text: "Username", Widget: usernameEntry},
			{Text: "Password", Widget: passwordEntry},
		},
	}

	// Create a list of credentials
	credentialsList := widget.NewList(
		func() int {
			credentials := loadCredentials()
			return len(credentials.List)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			credentials := loadCredentials()
			credential := credentials.List[i]
			label := obj.(*widget.Label)
			label.SetText(credential.Name)
		},
	)

	// Create a delete button for each credential in the list
	credentialsList.OnSelected = func(id widget.ListItemID) {
		dialog.NewConfirm("Delete credential", "Are you sure you want to delete this credential?", func(b bool) {
			if b {
				credentials := loadCredentials()
				credentials.List = append(credentials.List[:id], credentials.List[id+1:]...)
				saveCredentials(credentials)
				credentialsList.Refresh()
			}
		}, w).Show()
	}

	// Add the form and list to the window
	w.SetContent(container.NewVBox(
		form,
		widget.NewLabel("Credentials:"),
		credentialsList,
	))

	w.ShowAndRun()
}
