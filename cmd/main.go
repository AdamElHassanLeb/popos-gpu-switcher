package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var AppMessages *Messages

const SelectedLanguage string = "en"

func main() {
	var err error
	AppMessages, err = LoadMessages()
	if err != nil {
		panic(err)
	}

	a := app.New()
	w := a.NewWindow("Test Window")
	w.SetContent(container.NewVBox(
		widget.NewLabel("If you can read this, Fyne is working." + SelectedLanguage + AppMessages.Errors[SelectedLanguage]["ErrUnsupportedOS"]),
	))
	w.Resize(fyne.NewSize(420, 240))
	w.CenterOnScreen()
	w.ShowAndRun()
}
