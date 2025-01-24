package formulas

import (
	"os/exec"
	"runtime"
)

func InstallNode() {
	switch runtime.GOOS {
	case "darwin":
		installNodeMac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installNodeMac() {
	url := "https://nodejs.org/dist/v20.18.0/node-v20.18.0.pkg"

	boldGreen.Println("Starting Node installation 🚀")
	yellow.Println("Downloading Node...")
	download := exec.Command("curl", "-L", url, "-o", "node.pkg")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Go:", err)
		return
	}

	yellow.Println("Installing Node...")
	install := exec.Command("sudo", "installer", "-pkg", "node.pkg", "-target", "/")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Node:", err)
		return
	}

	boldGreen.Println("Node installed successfully 🎉")
}
