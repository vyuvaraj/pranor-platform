# Build All Pranor Components
# Usage: .\scripts\build-all.ps1 -Version v0.1.0 -OS windows -Arch amd64
# Output: dist/pranor-<version>-<os>-<arch>/

param(
    [Parameter(Mandatory=$true)][string]$Version,
    [Parameter(Mandatory=$true)][string]$OS,
    [Parameter(Mandatory=$true)][string]$Arch,
    [string]$OutputDir = "dist"
)

$ErrorActionPreference = "Stop"
$ext = if ($OS -eq "windows") { ".exe" } else { "" }
$archiveDir = "$OutputDir/pranor-${Version}-${OS}-${Arch}"

Write-Host "Building Pranor $Version for ${OS}/${Arch}..." -ForegroundColor Cyan

# Component definitions: [name, source_dir, go_module_path]
$components = @(
    @("pranor",         "../pranor-lang",     "."),
    @("pranor-gate",     "../pranor-gate",      "."),
    @("pranor-vault",    "../pranor-vault",     "./cmd/pranor-vault/"),
    @("pranor-pulse",    "../pranor-pulse",     "."),
    @("pranor-console",  "../pranor-console",   "."),
    @("pranor-cache",    "../pranor-cache",     "."),
    @("pranor-mesh",     "../pranor-mesh",      "."),
    @("pranor-chrono",     "../pranor-chrono",      "."),
    @("pranor-deploy",    "../pranor-deploy",     "."),
    @("pranor-trace",    "../pranor-trace",     "."),
    @("pranor-tunnel",   "../pranor-tunnel",    "."),
    @("pranor-auth",     "../pranor-auth",      "."),
    @("servdb",       "../pranor-pool",      "."),
    @("pranor-notify",     "../pranor-notify",      "."),
    @("pranor-flow",     "../pranor-flow",      "."),
    @("pranor-hub", "../pranor-hub",  "."),
    @("servlock",     "../pranor-lock",      ".")
)

# Prepare output directory
if (Test-Path $archiveDir) { Remove-Item -Recurse -Force $archiveDir }
New-Item -ItemType Directory -Path $archiveDir -Force | Out-Null

# Set Go cross-compilation env
$env:GOOS = $OS
$env:GOARCH = $Arch
$env:CGO_ENABLED = "0"

$success = 0
$failed = 0

foreach ($comp in $components) {
    $name = $comp[0]
    $srcDir = $comp[1]
    $module = $comp[2]
    $output = Join-Path $archiveDir "${name}${ext}"

    if (-not (Test-Path $srcDir)) {
        Write-Host "  SKIP $name (directory not found: $srcDir)" -ForegroundColor Yellow
        $failed++
        continue
    }

    Write-Host "  Building $name..." -NoNewline
    try {
        $absArchiveDir = Resolve-Path $archiveDir
        $argList = "build -ldflags `"-s -w -X main.version=$Version`" -o `"$absArchiveDir/${name}${ext}`" $module"
        $proc = Start-Process -FilePath "go" -ArgumentList $argList -WorkingDirectory $srcDir -NoNewWindow -Wait -PassThru -RedirectStandardError "$env:TEMP\go-err.txt"
        if ($proc.ExitCode -eq 0) {
            Write-Host " OK" -ForegroundColor Green
            $success++
        } else {
            $errMsg = Get-Content "$env:TEMP\go-err.txt" -Raw
            Write-Host " FAILED" -ForegroundColor Red
            Write-Host "    $errMsg" -ForegroundColor DarkRed
            $failed++
        }
    } catch {
        Write-Host " ERROR: $_" -ForegroundColor Red
        $failed++
    }
}

# Copy the launcher (built separately)
$launcherSrc = "cmd/pranor/main.go"
if (Test-Path $launcherSrc) {
    Write-Host "  Building pranor launcher..." -NoNewline
    go build -ldflags "-s -w -X main.version=$Version" -o "$archiveDir/pranor${ext}" ./cmd/pranor/
    Write-Host " OK" -ForegroundColor Green
    $success++
}

Write-Host ""
Write-Host "Results: $success built, $failed failed" -ForegroundColor $(if ($failed -eq 0) { "Green" } else { "Yellow" })
Write-Host "Output:  $archiveDir/" -ForegroundColor Cyan
