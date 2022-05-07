package main

import (
	"biogrid/internal/boot"
	"biogrid/internal/config"
	"biogrid/internal/windows"
	"fyne.io/fyne/v2/app"
)

func main() {
	runApp()
}

func runApp() {
	boot.Init()
	a := app.New()
	initConfig := config.InitConfig()
	mainWindow := windows.NewWindow(a, initConfig)
	mainWindow.ShowAndRun()
}
