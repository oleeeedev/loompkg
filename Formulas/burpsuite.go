package formulas

import (
	"fmt"
	"os/exec"
	"runtime"
)

func InstallBurpSuite() {
	switch runtime.GOARCH {
	case "arm64":
		InstallBurpSuiteMac("https://portswigger.net/burp/releases/community/latest")
	case "amd64":
		InstallBurpSuiteMac("https://portswigger.net/burp/releases/community/latest")
	}
}

func InstallBurpSuiteMac(url string) {
	download := exec.Command("curl", "-L", url)
	if err := download.Run(); err != nil {
		fmt.Println("Error downloading burpsuite:", err)
		return
	}
}
