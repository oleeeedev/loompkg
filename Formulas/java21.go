package formulas

import (
	"os"
	"os/exec"
	"runtime"
)

func InstallJava21() {
	switch runtime.GOOS {
	case "darwin":
		Java21Mac()
	}
}

func Java21Mac() {
	url := "https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21.0.5%2B11/OpenJDK21U-jdk_x64_mac_hotspot_21.0.5_11.pkg"
	boldGreen.Println("Starting Java21 installation 🚀")
	yellow.Println("Downloading Java21...")
	download := exec.Command("curl", "-L", url, "-o", "java21.pkg")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Java21:", err)
		return
	}

	yellow.Println("Installing Java21...")
	install := exec.Command("sudo", "installer", "-pkg", "java21.pkg", "-target", "/")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Java21:", err)
		return
	}

	boldGreen.Println("Java21 installed successfully 🎉")

	yellow.Println("Cleaning up...")
	if err := os.Remove("java21.pkg"); err != nil {
		redBold.Println("Error removing installer package:", err)
		return
	}

	boldGreen.Println("Cleanup completed.")
}
