package formulas

import (
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

var (
	boldGreen = color.New(color.FgGreen, color.Bold)
	redBold   = color.New(color.FgRed, color.Bold)
	yellow    = color.New(color.FgYellow)
)

func InstallNuclei() {
	switch runtime.GOOS {
	case "darwin":
		installNucleiMac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installNucleiMac() {
	boldGreen.Println("Starting Nuclei installation 🚀")

	yellow.Println("Downloading Nuclei...")

	cmd := exec.Command("go", "get", "github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest")
	if err := cmd.Run(); err != nil {
		redBold.Println("Error installing Nuclei:", err)
		return
	}

	boldGreen.Println("Nuclei installed successfully 🎉")
}
