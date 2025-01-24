package handlers

import (
	"os"
)

func Uninstall(pkg string) {
	p := "/usr/local/bin/" + pkg
	if err := os.Remove(p); err != nil {
		redBold.Println("Error deleting the binary:", err)
		return
	}

	boldGreen.Println(pkg + " Successfully removed. 🎉")
}
