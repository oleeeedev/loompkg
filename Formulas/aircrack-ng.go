package formulas

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func InstallAircrackNg() {
	switch runtime.GOOS {
	case "darwin":
		installAircrackNgMac()
	default:
		fmt.Println("This script only supports macOS (darwin).")
	}
}

func installAircrackNgMac() {
	fmt.Println("Starting Aircrack-ng installation 🚀")

	zipURL := "https://github.com/aircrack-ng/aircrack-ng/archive/refs/heads/master.zip"
	zipFile := "aircrack-ng.zip"

	if err := downloadFile(zipFile, zipURL); err != nil {
		fmt.Printf("Error downloading aircrack-ng: %v\n", err)
		return
	}
	defer os.Remove(zipFile)

	if err := unzip(zipFile, "aircrack-ng"); err != nil {
		fmt.Printf("Error extracting aircrack-ng: %v\n", err)
		return
	}

	if err := os.Chdir("aircrack-ng/aircrack-ng-master"); err != nil {
		fmt.Printf("Error changing to aircrack-ng directory: %v\n", err)
		return
	}

	fmt.Println("Building aircrack-ng...")
	if err := runCommand("autoreconf", "-i"); err != nil {
		fmt.Println("Error running autoreconf:", err)
		return
	}
	if err := runCommand("./configure", "--with-experimental"); err != nil {
		fmt.Println("Error running configure:", err)
		return
	}
	if err := runCommand("make"); err != nil {
		fmt.Println("Error running make:", err)
		return
	}
	if err := runCommand("sudo", "make", "install"); err != nil {
		fmt.Println("Error running make install:", err)
		return
	}

	fmt.Println("aircrack-ng installed successfully 🎉")
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fPath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fPath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
