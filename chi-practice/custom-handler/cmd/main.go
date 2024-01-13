package main

import (
	customhandler "customhandler"
)

func main() {
	app := customhandler.NewApp()
	app.Run()
}
