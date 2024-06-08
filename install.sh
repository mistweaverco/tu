#!/usr/bin/env bash

DL_PATH=""
BIN_NAME="tu"
TMP_DIR=$(mktemp -d)
BIN_DIR="/usr/local/bin"

# Check if curl is installed
# if not, exit with error
if ! command -v curl &> /dev/null; then
  echo "curl is required to install tu"
  exit 1
fi

# Check OS and set download path
if [[ "$OSTYPE" == "linux"* ]]; then
  DL_PATH="https://github.com/mistweaverco/tu/releases/download/latest/tu-linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  DL_PATH="https://github.com/mistweaverco/tu/releases/download/latest/tu-macos"
else
  echo "Unsupported OS"
  exit 1
fi

# Check if default BIN_DIR exists,
# if not, use /usr/bin
if [ ! -d "$BIN_DIR" ]; then
  BIN_DIR="/usr/bin"
fi

# Check if tu is already installed
if [ -f "$BIN_DIR/$BIN_NAME" ]; then
  echo "tu is already installed"
  exit 0
fi

# Download
curl -L "$DL_PATH" -o "$TMP_DIR/$BIN_NAME"

# Make it executable
chmod +x "$TMP_DIR/$BIN_NAME"

# if user is root, do not use sudo
if [ "$EUID" -eq 0 ]; then
  mv "$TMP_DIR/$BIN_NAME" "$BIN_DIR/$BIN_NAME"
else
  sudo mv "$TMP_DIR/$BIN_NAME" "$BIN_DIR/$BIN_NAME"
fi

# Clean up
rm -rf "$TMP_DIR"
