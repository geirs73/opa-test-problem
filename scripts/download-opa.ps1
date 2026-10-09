<#
.SYNOPSIS
Downloads OPA release binaries into the repo root (delegates to the Go downloader).
#>
[CmdletBinding()]
param(
    [string[]]$Versions = @()
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    throw 'go is required but not found on PATH (https://go.dev/dl/).'
}

$repoRoot = Split-Path -Parent $PSScriptRoot
Push-Location $repoRoot
try {
    & go run ./cmd/download-opa @Versions
    if ($LASTEXITCODE -ne 0) {
        throw "download-opa failed with exit code $LASTEXITCODE."
    }
}
finally {
    Pop-Location
}
