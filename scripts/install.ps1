# Servverse Installer for Windows
# Usage: irm https://raw.githubusercontent.com/vyuvaraj/servverse/main/scripts/install.ps1 | iex

param(
    [string]$Version = "latest",
    [string]$InstallDir = "$env:USERPROFILE\.servverse",
    [switch]$Uninstall
)

$ErrorActionPreference = "Stop"
$Repo = "vyuvaraj/servverse"
$BinDir = Join-Path $InstallDir "bin"

function Write-Status($msg) { Write-Host "  [servverse] $msg" -ForegroundColor Cyan }
function Write-Ok($msg) { Write-Host "  [✓] $msg" -ForegroundColor Green }
function Write-Err($msg) { Write-Host "  [✗] $msg" -ForegroundColor Red }

# --- Uninstall ---
if ($Uninstall) {
    Write-Status "Uninstalling Servverse..."
    if (Test-Path $InstallDir) {
        Remove-Item -Recurse -Force $InstallDir
        Write-Ok "Removed $InstallDir"
    }
    # Remove from PATH
    $userPath = [Environment]::GetEnvironmentVariable("PATH", "User")
    if ($userPath -like "*$BinDir*") {
        $newPath = ($userPath -split ";" | Where-Object { $_ -ne $BinDir }) -join ";"
        [Environment]::SetEnvironmentVariable("PATH", $newPath, "User")
        Write-Ok "Removed from PATH"
    }
    Write-Ok "Servverse uninstalled."
    return
}

# --- Detect architecture ---
$arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }
$os = "windows"
Write-Status "Platform: ${os}/${arch}"

# --- Resolve version ---
if ($Version -eq "latest") {
    Write-Status "Fetching latest release..."
    $release = Invoke-RestMethod "https://api.github.com/repos/$Repo/releases/latest"
    $Version = $release.tag_name
} else {
    $release = Invoke-RestMethod "https://api.github.com/repos/$Repo/releases/tags/$Version"
}
Write-Status "Version: $Version"

# --- Find archive asset ---
$assetName = "servverse-${Version}-${os}-${arch}.zip"
$asset = $release.assets | Where-Object { $_.name -eq $assetName }

if (-not $asset) {
    # Fallback: try without 'v' prefix
    $assetName = "servverse-$($Version.TrimStart('v'))-${os}-${arch}.zip"
    $asset = $release.assets | Where-Object { $_.name -eq $assetName }
}

if (-not $asset) {
    Write-Err "Could not find asset: $assetName"
    Write-Err "Available assets:"
    $release.assets | ForEach-Object { Write-Host "    - $($_.name)" }
    exit 1
}

# --- Download ---
$tempFile = Join-Path $env:TEMP $assetName
Write-Status "Downloading $assetName..."
Invoke-WebRequest -Uri $asset.browser_download_url -OutFile $tempFile -UseBasicParsing

# --- Extract ---
Write-Status "Installing to $BinDir..."
if (Test-Path $BinDir) { Remove-Item -Recurse -Force $BinDir }
New-Item -ItemType Directory -Path $BinDir -Force | Out-Null
Expand-Archive -Path $tempFile -DestinationPath $BinDir -Force
Remove-Item $tempFile -Force

# --- Add to PATH ---
$userPath = [Environment]::GetEnvironmentVariable("PATH", "User")
if ($userPath -notlike "*$BinDir*") {
    [Environment]::SetEnvironmentVariable("PATH", "$BinDir;$userPath", "User")
    $env:PATH = "$BinDir;$env:PATH"
    Write-Ok "Added $BinDir to PATH"
} else {
    Write-Ok "Already in PATH"
}

# --- Verify ---
$binaries = Get-ChildItem $BinDir -Filter "*.exe" | Select-Object -ExpandProperty Name
Write-Ok "Installed $($binaries.Count) binaries:"
$binaries | ForEach-Object { Write-Host "    $_" -ForegroundColor DarkGray }

# --- Test serv ---
$servPath = Join-Path $BinDir "pranor.exe"
if (Test-Path $servPath) {
    Write-Host ""
    Write-Ok "Servverse $Version installed successfully!"
    Write-Host ""
    Write-Host "  Quick start:" -ForegroundColor White
    Write-Host "    servverse up          # Start all services" -ForegroundColor DarkGray
    Write-Host "    servverse status      # Check service health" -ForegroundColor DarkGray
    Write-Host "    pranor run app.pnr      # Run a .pnr file" -ForegroundColor DarkGray
    Write-Host ""
    Write-Host "  Open a new terminal for PATH changes to take effect." -ForegroundColor Yellow
} else {
    Write-Err "Installation completed but pranor.exe not found in $BinDir"
}
