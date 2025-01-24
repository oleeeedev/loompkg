package handlers

import (
	lib "loompkg/lib"

	"github.com/fatih/color"
)

var (
	boldGreen    = color.New(color.FgGreen, color.Bold)
	redBold      = color.New(color.FgRed, color.Bold)
	yellow       = color.New(color.FgYellow)
	yellowItalic = color.New(color.FgYellow, color.BgBlack, color.Italic)
)

func Install(pkg string) {
	yellowItalic.Println("This may take a while... 😴")
	installFunc, exists := lib.GetTool(pkg)
	if !exists {
		redBold.Printf("Error: Package %s is not recognized.\n", pkg)
		return
	}

	installFunc()
}
