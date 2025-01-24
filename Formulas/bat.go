package formulas

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func InstallBat() {

	switch runtime.GOOS {
	case "darwin":
		installBazelMac()
	default:
		fmt.Println("OS No Support")
	}

}

func InstallBatMac() {
	url := "https://github.com/sharkdp/bat/releases/download/v0.24.0/bat-v0.24.0-x86_64-apple-darwin.tar.gz"
	download := exec.Command("curl", "-L", url, "-o", "bat.tar.gz")
	if err := download.Run(); err != nil {
		fmt.Println("Error downloading bat:", err)
		return
	}
	defer os.Remove("bat.tar.gz")

	extract := exec.Command("tar", "-xzf", "bat.tar.gz")
	if err := extract.Run(); err != nil {
		fmt.Println("Error extracting bat:", err)
		return
	}
}
