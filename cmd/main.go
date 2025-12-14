package main

var AppMessages *Messages

const SelectedLanguage string = "en"

func main() {
	var err error
	AppMessages, err = LoadMessages()
	if err != nil {
		panic(err)
	}

	StartUI()

}
