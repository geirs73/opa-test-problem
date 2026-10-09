#!/usr/bin/env bash
# Downloads OPA release binaries into the repo root (delegates to the Go downloader).
set -euo pipefail

if ! command -v go >/dev/null 2>&1; then
  echo "Error: go is required but not found on PATH (https://go.dev/dl/)." >&2
  exit 1
fi

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$REPO_ROOT"
exec go run ./cmd/download-opa "$@"
