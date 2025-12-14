package main

import (
	"encoding/json"
	"fmt"

	"github.com/AdamElHassanLeb/popos-gpu-switcher/assets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Messages struct {
	Languages []string                     `json:"Languages"`
	Errors    map[string]map[string]string `json:"Errors"`
	Systems   map[string]SystemMessages    `json:"Systems"`
}

type SystemMessages struct {
	Modes map[string]map[string]ModeMessage `json:"Modes"`
}

type ModeMessage struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

func LoadMessages() (*Messages, error) {
	data, err := assets.FS.ReadFile("messages/messages.json")
	if err != nil {
		return nil, fmt.Errorf("read messages.json: %w", err)
	}

	var msgs Messages
	if err := json.Unmarshal(data, &msgs); err != nil {
		return nil, fmt.Errorf("parse messages.json: %w", err)
	}

	return &msgs, nil
}

func StartUI() {
	a := app.New()
	w := a.NewWindow("Test Window")
	w.SetContent(container.NewVBox(
		widget.NewLabel("If you can read this, Fyne is working."),
	))
	w.Resize(fyne.NewSize(420, 240))
	w.CenterOnScreen()
	w.ShowAndRun()
}
