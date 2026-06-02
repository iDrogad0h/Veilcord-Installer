// Package injector contains Veilcord's core install/uninstall logic.
//
// How it works (same proven method Vencord/Equicord use — no memory or DLL
// injection, just files):
//
//   Discord's desktop app is an Electron app. Its real code lives in
//   <resources>/app.asar. To inject, we:
//     1. close Discord,
//     2. rename app.asar  ->  _app.asar      (the untouched original),
//     3. create an <resources>/app/ folder with our loader (index.js),
//        which re-points Electron at _app.asar and then starts Discord normally.
//   To uninstall we just delete app/ and rename _app.asar back to app.asar.
//
// This package uses ONLY the Go standard library, so it cross-compiles to a
// Windows .exe with no external dependencies.
package injector

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// The files we drop into <resources>/app/. For this first build it's a tiny,
// safe loader that proves injection works. Later this folder is replaced by the
// real Veilcord build (the Equicord fork's dist/).
//
//go:embed all:payload
var payloadFS embed.FS

// Discord describes one detected installation (Stable, PTB, Canary, ...).
type Discord struct {
	Branch    string // e.g. "Discord", "DiscordPTB"
	Resources string // path to the .../resources folder
	Proc      string // process name to close (e.g. "Discord.exe")
}

func (d Discord) asar() string   { return filepath.Join(d.Resources, "app.asar") }
func (d Discord) backup() string { return filepath.Join(d.Resources, "_app.asar") }
func (d Discord) appDir() string { return filepath.Join(d.Resources, "app") }

// IsInstalled reports whether Veilcord is currently injected here.
func (d Discord) IsInstalled() bool {
	if _, err := os.Stat(d.backup()); err != nil {
		return false
	}
	_, err := os.Stat(d.appDir())
	return err == nil
}

// FindDiscords returns every Discord installation we can patch on this OS.
func FindDiscords() []Discord {
	switch runtime.GOOS {
	case "windows":
		return findWindows()
	case "darwin":
		return findMac()
	default:
		return nil // Linux: use your distro package / Equibop for now
	}
}

func findWindows() []Discord {
	local := os.Getenv("LOCALAPPDATA")
	if local == "" {
		return nil
	}
	branches := []struct{ dir, proc string }{
		{"Discord", "Discord.exe"},
		{"DiscordPTB", "DiscordPTB.exe"},
		{"DiscordCanary", "DiscordCanary.exe"},
		{"DiscordDevelopment", "DiscordDevelopment.exe"},
	}
	var out []Discord
	for _, b := range branches {
		res := latestResources(filepath.Join(local, b.dir))
		if res != "" {
			out = append(out, Discord{Branch: b.dir, Resources: res, Proc: b.proc})
		}
	}
	return out
}

func findMac() []Discord {
	branches := []struct{ path, branch, proc string }{
		{"/Applications/Discord.app/Contents/Resources", "Discord", "Discord"},
		{"/Applications/Discord PTB.app/Contents/Resources", "DiscordPTB", "Discord PTB"},
		{"/Applications/Discord Canary.app/Contents/Resources", "DiscordCanary", "Discord Canary"},
	}
	var out []Discord
	for _, b := range branches {
		if usableResources(b.path) {
			out = append(out, Discord{Branch: b.branch, Resources: b.path, Proc: b.proc})
		}
	}
	return out
}

// latestResources finds the newest app-*/resources folder under root.
// Discord keeps several app-<version> folders; the most recent one is live.
func latestResources(root string) string {
	entries, err := os.ReadDir(root)
	if err != nil {
		return ""
	}
	best := ""
	var bestMod time.Time
	for _, e := range entries {
		if !e.IsDir() || !strings.HasPrefix(e.Name(), "app-") {
			continue
		}
		res := filepath.Join(root, e.Name(), "resources")
		if !usableResources(res) {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		if best == "" || info.ModTime().After(bestMod) {
			best, bestMod = res, info.ModTime()
		}
	}
	return best
}

func usableResources(res string) bool {
	for _, name := range []string{"app.asar", "_app.asar", "app"} {
		if _, err := os.Stat(filepath.Join(res, name)); err == nil {
			return true
		}
	}
	return false
}

func closeDiscord(d Discord) {
	switch runtime.GOOS {
	case "windows":
		_ = exec.Command("taskkill", "/F", "/IM", d.Proc).Run()
	case "darwin":
		_ = exec.Command("pkill", "-x", d.Proc).Run()
	}
	time.Sleep(800 * time.Millisecond) // give the OS a moment to release the files
}

// Install injects Veilcord into the given Discord installation.
func Install(d Discord) error {
	closeDiscord(d)

	// 1) Make sure we have the untouched original backed up as _app.asar.
	if _, err := os.Stat(d.backup()); err != nil {
		if _, err := os.Stat(d.asar()); err != nil {
			return fmt.Errorf("no encontré app.asar en %s (¿Discord está instalado y se abrió al menos una vez?)", d.Resources)
		}
		if err := os.Rename(d.asar(), d.backup()); err != nil {
			return fmt.Errorf("no pude renombrar app.asar (cerrá Discord y reintenta): %w", err)
		}
	}

	// 2) (Re)create the app/ folder from our embedded loader.
	if err := os.RemoveAll(d.appDir()); err != nil {
		return err
	}
	if err := os.MkdirAll(d.appDir(), 0o755); err != nil {
		return err
	}
	return extractPayload(d.appDir())
}

// Uninstall removes Veilcord and restores a clean Discord.
func Uninstall(d Discord) error {
	closeDiscord(d)

	if _, err := os.Stat(d.appDir()); err == nil {
		if err := os.RemoveAll(d.appDir()); err != nil {
			return fmt.Errorf("no pude borrar la carpeta app: %w", err)
		}
	}
	if _, err := os.Stat(d.backup()); err == nil {
		_ = os.Remove(d.asar()) // remove a stray app.asar if one exists
		if err := os.Rename(d.backup(), d.asar()); err != nil {
			return fmt.Errorf("no pude restaurar app.asar: %w", err)
		}
	}
	return nil
}

// extractPayload copies the embedded payload/ files into dst (the app/ folder).
func extractPayload(dst string) error {
	return fs.WalkDir(payloadFS, "payload", func(p string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel := strings.TrimPrefix(strings.TrimPrefix(p, "payload"), "/")
		if rel == "" {
			return nil
		}
		target := filepath.Join(dst, rel)
		if entry.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		data, err := payloadFS.ReadFile(p)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		return os.WriteFile(target, data, 0o644)
	})
}
