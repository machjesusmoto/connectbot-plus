#!/bin/bash
# ConnectBot Plus - Quick Installer Script
# This script downloads and runs the appropriate installer for your system

set -e

echo "ConnectBot Plus Installer"
echo "========================"

# Detect architecture
ARCH=$(uname -m)
case $ARCH in
    x86_64)
        BINARY="installer-linux-amd64"
        ;;
    aarch64)
        BINARY="installer-linux-arm64"
        ;;
    armv7l)
        BINARY="installer-linux-arm"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Download URL (update this after first release)
BASE_URL="https://github.com/machjesusmoto/connectbot-plus/releases/latest/download"
DOWNLOAD_URL="$BASE_URL/$BINARY"

echo "Downloading installer for $ARCH..."

# Download installer
if command -v curl &> /dev/null; then
    curl -L -o /tmp/cbplus-installer "$DOWNLOAD_URL"
elif command -v wget &> /dev/null; then
    wget -O /tmp/cbplus-installer "$DOWNLOAD_URL"
else
    echo "Error: Please install curl or wget"
    exit 1
fi

# Make executable
chmod +x /tmp/cbplus-installer

# Run installer
sudo /tmp/cbplus-installer

# Cleanup
rm -f /tmp/cbplus-installer

echo "\nInstallation complete!"