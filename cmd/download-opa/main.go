// Command download-opa downloads OPA release binaries into the current
// directory as opa_v{VERSION}[.exe], verifying their sha256 checksums.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const releaseURL = "https://github.com/open-policy-agent/opa/releases/download"

var defaultVersions = []string{"1.20.2", "1.21.1"}

func main() {
	versions := defaultVersions
	if len(os.Args) > 1 {
		versions = os.Args[1:]
	}

	osName, arch, ext, err := platform()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	for _, v := range versions {
		if err := download(v, osName, arch, ext); err != nil {
			fmt.Fprintf(os.Stderr, "error: opa v%s: %v\n", v, err)
			os.Exit(1)
		}
	}
}

func platform() (osName, arch, ext string, err error) {
	switch runtime.GOOS {
	case "windows", "darwin", "linux":
		osName = runtime.GOOS
	default:
		return "", "", "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
	switch runtime.GOARCH {
	case "amd64", "arm64":
		arch = runtime.GOARCH
	default:
		return "", "", "", fmt.Errorf("unsupported architecture: %s", runtime.GOARCH)
	}
	if osName == "windows" {
		ext = ".exe"
		if arch == "arm64" {
			// No windows/arm64 release asset; amd64 runs under emulation.
			fmt.Fprintln(os.Stderr, "warning: no windows/arm64 OPA release asset, falling back to amd64")
			arch = "amd64"
		}
	}
	return osName, arch, ext, nil
}

func download(version, osName, arch, ext string) error {
	target := fmt.Sprintf("opa_v%s%s", version, ext)
	if _, err := os.Stat(target); err == nil {
		fmt.Printf("%s already exists, skipping.\n", target)
		return nil
	}

	asset := fmt.Sprintf("opa_%s_%s%s", osName, arch, ext)
	url := fmt.Sprintf("%s/v%s/%s", releaseURL, version, asset)

	wantSum, err := fetchChecksum(url + ".sha256")
	if err != nil {
		return fmt.Errorf("fetching checksum: %w", err)
	}

	fmt.Printf("Downloading %s -> %s ...\n", url, target)
	tmp := target + ".tmp"
	gotSum, err := downloadFile(url, tmp)
	if err != nil {
		os.Remove(tmp)
		return fmt.Errorf("downloading: %w", err)
	}

	if gotSum != wantSum {
		os.Remove(tmp)
		return fmt.Errorf("sha256 mismatch: expected %s, got %s", wantSum, gotSum)
	}

	if runtime.GOOS != "windows" {
		if err := os.Chmod(tmp, 0o755); err != nil {
			os.Remove(tmp)
			return err
		}
	}
	if err := os.Rename(tmp, target); err != nil {
		os.Remove(tmp)
		return err
	}
	fmt.Printf("%s downloaded and checksum verified.\n", target)
	return nil
}

// fetchChecksum reads a "<hash>  <filename>" sha256 file and returns the hash.
func fetchChecksum(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GET %s: %s", url, resp.Status)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1024))
	if err != nil {
		return "", err
	}
	fields := strings.Fields(string(body))
	if len(fields) == 0 || len(fields[0]) != 64 {
		return "", fmt.Errorf("unexpected checksum file content: %q", string(body))
	}
	return strings.ToLower(fields[0]), nil
}

// downloadFile saves url to path and returns the sha256 hex digest of the data.
func downloadFile(url, path string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GET %s: %s", url, resp.Status)
	}

	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	_, err = io.Copy(io.MultiWriter(f, hash), resp.Body)
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
