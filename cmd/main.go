package main

import assetmanager "github.com/AdamElHassanLeb/popos-gpu-switcher/assetmanager"

var AppMessages *assetmanager.Messages

const SelectedLanguage string = "en"

func main() {
	var err error
	AppMessages, err = assetmanager.LoadMessages()
	if err != nil {
		panic(err)
	}

	StartUI()
}
