# PowerShell script to push all Pranor ecosystem repositories to origin main

$repos = @(
    "pranor-lang",
    "pranor-gate",
    "pranor-vault",
    "pranor-pulse",
    "pranor-console",
    "pranor-cache",
    "pranor-mesh",
    "pranor-chrono",
    "pranor-deploy",
    "pranor-trace",
    "pranor-tunnel",
    "pranor-auth",
    "pranor-pool",
    "pranor-notify",
    "pranor-flow",
    "pranor-hub",
    "pranor-core",
    "ServDocs",
    "pranor-repo"
)

Write-Host "==================================================" -ForegroundColor Cyan
Write-Host "Pushing all Pranor ecosystem repositories to origin..." -ForegroundColor Cyan
Write-Host "==================================================" -ForegroundColor Cyan

foreach ($repo in $repos) {
    # Check if the folder exists in parent or current dir
    $repoPath = Join-Path ".." $repo
    if (-not (Test-Path $repoPath)) {
        $repoPath = Join-Path "." $repo
        if (-not (Test-Path $repoPath)) {
            continue
        }
    }

    $gitPath = Join-Path $repoPath ".git"
    if (Test-Path $gitPath) {
        Write-Host "Processing repository: $repo..." -ForegroundColor Yellow
        Push-Location $repoPath
        try {
            # Execute git push origin main
            git push origin main
            Write-Host "✅ Successfully processed $repo" -ForegroundColor Green
        } catch {
            Write-Host "❌ Failed to push $repo" -ForegroundColor Red
        }
        Pop-Location
        Write-Host ""
    }
}

Write-Host "Done pushing all repositories!" -ForegroundColor Green
