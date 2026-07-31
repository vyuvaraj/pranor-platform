# PowerShell script to create and push release tag v0.2.0 for all repositories

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
    "pranor-ee",
    "pranor-operator",
    "pranor-repo"
)

$tag = "v1.0.0-rc1"
$message = "Release v1.0.0-rc1"

Write-Host "==================================================" -ForegroundColor Cyan
Write-Host "Tagging and pushing release $tag to all origins..." -ForegroundColor Cyan
Write-Host "==================================================" -ForegroundColor Cyan

foreach ($repo in $repos) {
    $repoPath = Resolve-Path (Join-Path $PSScriptRoot "../../$repo") -ErrorAction SilentlyContinue
    if (-not $repoPath) {
        continue
    }

    $gitPath = Join-Path $repoPath ".git"
    if (Test-Path $gitPath) {
        Write-Host "Processing repository: $repo..." -ForegroundColor Yellow
        Push-Location $repoPath
        try {
            # Delete tag locally and remotely if it exists
            git tag -d $tag 2>$null
            git push --delete origin $tag 2>$null

            # Tag and push
            git tag -a $tag -m $message
            git push origin $tag
            Write-Host "✅ Successfully tagged and pushed $repo as $tag" -ForegroundColor Green
        } catch {
            Write-Host "❌ Failed to tag/push $repo" -ForegroundColor Red
        }
        Pop-Location
        Write-Host ""
    }
}

Write-Host "All repositories release tagging complete!" -ForegroundColor Green
