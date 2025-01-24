package formulas

import (
	"os"
	"os/exec"
	"runtime"
)

func InstallK9s() {
	switch runtime.GOOS {
	case "darwin":
		installK9sMac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installK9sMac() {
	boldGreen.Println("Starting K9s installation 🚀")

	yellow.Println("Cloning K9s repository...")
	download := exec.Command("git", "clone", "https://github.com/derailed/k9s.git")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading K9s:", err)
		return
	}

	yellow.Println("Entering K9s directory...")
	if err := os.Chdir("k9s"); err != nil {
		redBold.Println("Error changing to K9s directory:", err)
		return
	}

	yellow.Println("Building K9s...")
	build := exec.Command("go", "build", "-o", "k9s", "./cmd")
	if err := build.Run(); err != nil {
		redBold.Println("Error building K9s:", err)
		return
	}

	yellow.Println("Moving K9s binary to /usr/local/bin...")
	if err := os.Rename("k9s", "/usr/local/bin/k9s"); err != nil {
		redBold.Println("Error moving K9s binary:", err)
		return
	}

	boldGreen.Println("K9s installed successfully 🎉")
}
