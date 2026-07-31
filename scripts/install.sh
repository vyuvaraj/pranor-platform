#!/bin/sh
# Pranor Installer for macOS/Linux
# Usage: curl -fsSL https://raw.githubusercontent.com/vyuvaraj/pranor/main/scripts/install.sh | sh

set -e

REPO="vyuvaraj/pranor"
INSTALL_DIR="${SERVVERSE_HOME:-$HOME/.pranor}"
BIN_DIR="$INSTALL_DIR/bin"
VERSION="${1:-latest}"

info() { printf "  \033[36m[pranor]\033[0m %s\n" "$1"; }
ok()   { printf "  \033[32m[✓]\033[0m %s\n" "$1"; }
err()  { printf "  \033[31m[✗]\033[0m %s\n" "$1"; exit 1; }

# --- Detect platform ---
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case "$ARCH" in
    x86_64|amd64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) err "Unsupported architecture: $ARCH" ;;
esac
case "$OS" in
    linux|darwin) ;;
    *) err "Unsupported OS: $OS" ;;
esac
info "Platform: ${OS}/${ARCH}"

# --- Resolve version ---
if [ "$VERSION" = "latest" ]; then
    info "Fetching latest release..."
    VERSION=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | head -1 | cut -d'"' -f4)
fi
info "Version: $VERSION"

# --- Download ---
ASSET_NAME="pranor-${VERSION}-${OS}-${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/$REPO/releases/download/${VERSION}/${ASSET_NAME}"
TEMP_FILE="/tmp/$ASSET_NAME"

info "Downloading $ASSET_NAME..."
if ! curl -fsSL "$DOWNLOAD_URL" -o "$TEMP_FILE" 2>/dev/null; then
    # Try without 'v' prefix
    ASSET_NAME="pranor-$(echo $VERSION | sed 's/^v//')-${OS}-${ARCH}.tar.gz"
    DOWNLOAD_URL="https://github.com/$REPO/releases/download/${VERSION}/${ASSET_NAME}"
    curl -fsSL "$DOWNLOAD_URL" -o "$TEMP_FILE" || err "Download failed: $DOWNLOAD_URL"
fi

# --- Extract ---
info "Installing to $BIN_DIR..."
mkdir -p "$BIN_DIR"
tar -xzf "$TEMP_FILE" -C "$BIN_DIR"
rm -f "$TEMP_FILE"
chmod +x "$BIN_DIR"/*

# --- Add to PATH ---
SHELL_RC=""
case "$SHELL" in
    */zsh) SHELL_RC="$HOME/.zshrc" ;;
    */bash) SHELL_RC="$HOME/.bashrc" ;;
    */fish) SHELL_RC="$HOME/.config/fish/config.fish" ;;
    *) SHELL_RC="$HOME/.profile" ;;
esac

if ! echo "$PATH" | grep -q "$BIN_DIR"; then
    if [ -n "$SHELL_RC" ]; then
        echo "" >> "$SHELL_RC"
        echo "# Pranor" >> "$SHELL_RC"
        echo "export PATH=\"$BIN_DIR:\$PATH\"" >> "$SHELL_RC"
        ok "Added $BIN_DIR to PATH in $SHELL_RC"
    fi
    export PATH="$BIN_DIR:$PATH"
else
    ok "Already in PATH"
fi

# --- Verify ---
BINARY_COUNT=$(ls "$BIN_DIR" | wc -l | tr -d ' ')
ok "Installed $BINARY_COUNT binaries"

if [ -x "$BIN_DIR/serv" ]; then
    echo ""
    ok "Pranor $VERSION installed successfully!"
    echo ""
    echo "  Quick start:"
    echo "    pranor up          # Start all services"
    echo "    pranor status      # Check service health"
    echo "    pranor run app.pnr      # Run a .pnr file"
    echo ""
    echo "  Restart your shell or run: export PATH=\"$BIN_DIR:\$PATH\""
else
    err "Installation completed but 'serv' binary not found in $BIN_DIR"
fi
