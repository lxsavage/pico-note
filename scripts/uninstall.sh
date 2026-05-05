#!/usr/bin/env bash

BINARY="pn"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"

BINARY_PATH="$INSTALL_DIR/$BINARY"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')

if [[ ! -x "$BINARY_PATH" ]]; then
  echo "$BINARY not found in $INSTALL_DIR; nothing to uninstall."
  exit 0
fi

CONFIG_DIR=$HOME/.config
if [[ $OS -eq darwin ]]; then
  CONFIG_DIR="$HOME/Library/Application Support"
fi

rm -f "$BINARY_PATH"
echo "$BINARY removed."
