package formulas

import (
	"os/exec"
	"runtime"
)

func InstallCargo() {
	switch runtime.GOOS {
	case "darwin":
		installCargoMac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}
func installCargoMac() {
	boldGreen.Println("Starting Rust installation 🚀")

	yellow.Println("Downloading Rust...")

	cmd := exec.Command("curl", "--proto", "'=https'", "--tlsv1.2", "-sSf", "https://sh.rustup.rs", "-o", "rustup.sh")
	if err := cmd.Run(); err != nil {
		redBold.Println("Error downloading Rust:", err)
		return
	}

	yellow.Println("Installing Rust...")

	cmd = exec.Command("sh", "rustup.sh", "-y")
	if err := cmd.Run(); err != nil {
		redBold.Println("Error installing Cargo:", err)
		return
	}

	boldGreen.Println("Rust installed successfully 🎉")
}
