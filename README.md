# opa.test.problem
Recreate issue we have when running tests with large json files in bundle dir

## Getting OPA binaries

The reproduction compares OPA v1.20.2 and v1.21.1. Download both into the repo root
(named `opa_v<version>` / `opa_v<version>.exe`). Go is required:

```sh
go run ./cmd/download-opa
```

Or use the convenience wrappers (they delegate to the Go program and can be run
from any directory):

```sh
# macOS / Linux / Windows Git Bash
./scripts/download-opa.sh
```

```powershell
# Windows PowerShell or PowerShell 7+ (any OS)
powershell -File scripts/download-opa.ps1
pwsh scripts/download-opa.ps1
```

The downloader detects OS/architecture, verifies the published sha256 checksums,
and skips files that are already downloaded. Other versions can be requested as
arguments, e.g. `go run ./cmd/download-opa 1.19.0`. The binaries are gitignored.
