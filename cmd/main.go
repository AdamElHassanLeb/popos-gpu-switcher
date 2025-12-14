package main

import (
	assetmanager "github.com/AdamElHassanLeb/popos-gpu-switcher/assetmanager"
	services "github.com/AdamElHassanLeb/popos-gpu-switcher/internal/services"
)

var AppMessages *assetmanager.Messages

func main() {
	var err error
	AppMessages, err = assetmanager.LoadMessages()
	if err != nil {
		panic(err)
	}

	service, err := services.GetGPUService()

	if err != nil {
		StartUI(service, false)
	} else {
		StartUI(service, true)
	}
}
