package formulas

import (
	"os"
	"os/exec"
	"runtime"
)

func InstallAws() {
	switch runtime.GOOS {
	case "darwin":
		installAwsMac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installAwsMac() {
	url := "https://awscli.amazonaws.com/AWSCLIV2.pkg"

	boldGreen.Println("Starting AWS CLI installation 🚀")

	yellow.Println("Downloading AWS CLI...")
	download := exec.Command("curl", "-L", url, "-o", "AWSCLIV2.pkg")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading AWS CLI:", err)
		return
	}

	yellow.Println("Installing AWS CLI...")
	install := exec.Command("sudo", "installer", "-pkg", "AWSCLIV2.pkg", "-target", "/")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing AWS CLI:", err)
		return
	}

	boldGreen.Println("AWS CLI installed successfully 🎉")

	yellow.Println("Cleaning up...")
	if err := os.Remove("AWSCLIV2.pkg"); err != nil {
		redBold.Println("Error removing installer package:", err)
		return
	}

	boldGreen.Println("Cleanup completed.")
}
