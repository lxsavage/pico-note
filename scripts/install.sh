#!/usr/bin/env bash

REPO="lxsavage/pico-note"
BINARY="pn"

# ---- 1. Map OS / Arch to release naming ----
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

case "$OS" in
  linux)  OS="linux" ;;
  darwin) OS="macos" ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

ASSET="${BINARY}-${OS}-${ARCH}"

# ---- 2. Pick install directory ----
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"
mkdir -p "$INSTALL_DIR"

# ---- 3. Download URL for latest release ----
URL=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" \
        | grep "browser_download_url.*$ASSET" \
        | head -n 1 \
        | cut -d '"' -f 4)

if [[ -z "$URL" ]]; then
  echo "Release asset $ASSET not found in latest release"
  exit 1
fi

# ---- 4. Download & install ----
TMP=$(mktemp -d)
trap "rm -rf $TMP" EXIT
curl -sSL "$URL" -o "$TMP/$ASSET"
chmod +x "$TMP/$ASSET"
mv "$TMP/$ASSET" "$INSTALL_DIR/$BINARY"

# ---- 5. Final message ----
echo "$BINARY installed to $INSTALL_DIR/$BINARY"
if ! command -v "$BINARY" >/dev/null 2>&1; then
  echo "Add $INSTALL_DIR to your PATH, e.g.:"
  echo "  export PATH=\"$INSTALL_DIR:\$PATH\""
fi
