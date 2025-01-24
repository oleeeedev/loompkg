package formulas

import (
	"os/exec"
	"runtime"
)

func InstallDocker() {
	switch runtime.GOARCH {
	case "amd64":
		installDockerMac("https://desktop.docker.com/mac/main/amd64/Docker.dmg?utm_source=docker&utm_medium=webreferral&utm_campaign=docs-driven-download-mac-amd64")
	case "arm64":
		installDockerMac("https://desktop.docker.com/mac/main/arm64/Docker.dmg?utm_source=docker&utm_medium=webreferral&utm_campaign=docs-driven-download-mac-arm64")
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installDockerMac(url string) {
	boldGreen.Println("Starting Docker installation 🚀")
	yellow.Println("Downloading Docker...")

	download := exec.Command("curl", "-L", url, "-o", "Docker.dmg")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Docker:", err)
		return
	}

	yellow.Println("Mounting the disk image...")
	mount := exec.Command("sudo", "hdiutil", "attach", "Docker.dmg")
	if err := mount.Run(); err != nil {
		redBold.Println("Error mounting the disk image:", err)
		return
	}

	yellow.Println("Running the Docker installer...")
	install := exec.Command("sudo", "/Volumes/Docker/Docker.app/Contents/MacOS/install")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Docker:", err)
		if err := exec.Command("sudo", "rm", "-rf", "/Volumes/Docker").Run(); err != nil {
			redBold.Println("Error cleaning up the disk image:", err)
			return
		}
		return
	}

	yellow.Println("Unmounting the disk image...")
	detach := exec.Command("sudo", "hdiutil", "detach", "/Volumes/Docker")
	if err := detach.Run(); err != nil {
		redBold.Println("Error unmounting the disk image:", err)
		return
	}

	boldGreen.Println("Docker installed successfully 🎉")
}
