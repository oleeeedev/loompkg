package formulas

import (
	"os"
	"os/exec"
	"runtime"
)

func InstallDotnet() {
	switch runtime.GOARCH {
	case "amd64":
		installDotnetMac("https://download.visualstudio.microsoft.com/download/pr/d6b3fe61-3c0e-45da-9e37-cef64d4fd3eb/28d536e0ecfbb330ab49454a5e3962f6/dotnet-sdk-8.0.403-osx-x64.pkg")
	case "arm64":
		installDotnetMac("https://download.visualstudio.microsoft.com/download/pr/35b0fb29-cadc-4083-aa26-6cecd2e7ffa1/1a9972a435b73ffdd0b462f979ea5b23/dotnet-sdk-8.0.403-osx-arm64.pkg")
	default:
		redBold.Println("This script only supports macOS")
	}
}

func installDotnetMac(url string) {
	boldGreen.Println("Starting Dotnet installation 🚀")
	yellow.Println("Downloading Dotnet...")
	download := exec.Command("curl", "-L", url, "-o", "dotnet.pkg")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Dotnet:", err)
		return
	}
	yellow.Println("Installing Dotnet...")
	install := exec.Command("sudo", "installer", "-pkg", "dotnet.pkg", "-target", "/")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Dotnet:", err)
		return
	}
	boldGreen.Println("Dotnet installed successfully 🎉")
	yellow.Println("Cleaning up...")
	if err := os.Remove("dotnet.pkg"); err != nil {
		redBold.Println("Error removing installer package:", err)
		return
	}
}
