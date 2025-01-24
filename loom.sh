#!/bin/bash

BOLD_GREEN='\033[1;32m'
BOLD_RED='\033[1;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

if ! command -v go &> /dev/null; then
    echo -e "${YELLOW}Go is not installed. Downloading and installing Go...${NC}"

    ARCH=$(uname -m)

    if [ "$ARCH" = "arm64" ]; then
        echo -e "${YELLOW}Detected Apple Silicon (M1/M2)...${NC}"
        GO_VERSION="1.20.4"
        GO_TAR="go$GO_VERSION.darwin-arm64.tar.gz"
    else
        echo -e "${YELLOW}Detected Intel...${NC}"
        GO_VERSION="1.20.4"
        GO_TAR="go$GO_VERSION.darwin-amd64.tar.gz"
    fi

    curl -LO "https://golang.org/dl/$GO_TAR"


    sudo tar -C /usr/local -xzf "$GO_TAR"

    echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bash_profile
    source ~/.bash_profile

    if ! command -v go &> /dev/null; then
        echo -e "${BOLD_RED}Error: Failed to install Go.${NC}"
        exit 1
    fi
else
    echo -e "${BOLD_GREEN}Go is already installed.${NC}"
fi

echo -e "${YELLOW}Cloning the Loom repository...${NC}"
git clone https://github.com/oleeeedev/loompkg
cd loompkg || { echo -e "${BOLD_RED}Error: Failed to enter the Loom directory.${NC}"; exit 1; }

echo -e "${YELLOW}Building Loom...${NC}"
go build -o loom

echo -e "${YELLOW}Moving Loom to /usr/local/bin...${NC}"
sudo mv loom /usr/local/bin

if command -v loom &> /dev/null; then
    echo -e "${BOLD_GREEN}Loom installed successfully! 🎉${NC}"
else
    echo -e "${BOLD_RED}Error: Loom installation failed.${NC}"
fi