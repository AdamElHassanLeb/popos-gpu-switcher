package main

import (
	"encoding/json"
	"fmt"

	"github.com/AdamElHassanLeb/popos-gpu-switcher/assets"
)

type Messages struct {
	Languages []string                     `json:"Languages"`
	Errors    map[string]map[string]string `json:"Errors"`
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
