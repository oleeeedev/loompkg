package formulas

import (
	"os/exec"
	"runtime"
)

func InstallPython2() {
	switch runtime.GOOS {
	case "darwin":
		installPython2Mac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installPython2Mac() {
	boldGreen.Println("Starting Python2 installation 🚀")
	yellow.Println("Downloading Python2...")
	download := exec.Command("curl", "-L", "https://www.python.org/ftp/python/2.7.18/python-2.7.18-macosx10.9.pkg", "-o", "python2.pkg")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Python2:", err)
		return
	}

	yellow.Println("Installing Python2...")
	install := exec.Command("sudo", "installer", "-pkg", "python2.pkg", "-target", "/")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Python2:", err)
		return
	}

	boldGreen.Println("Python2 installed successfully 🎉")
}
