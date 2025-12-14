package assetmanager

import (
	"encoding/json"
	"fmt"

	"github.com/AdamElHassanLeb/popos-gpu-switcher/assets"
)

type Messages struct {
	Languages map[string]LanguageMessages `json:"Languages"`
}

type LanguageMessages struct {
	Errors  map[string]string         `json:"Errors"`
	Systems map[string]SystemMessages `json:"Systems"`
}

type SystemMessages struct {
	Modes map[string]ModeMessage `json:"Modes"`
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
