//go:build darwin

package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Regression for discussion #474: deleting from analyze over SSH appeared to do
// nothing. The only Trash path was Finder AppleScript, which raises a dialog on
// the physical machine that a remote user cannot answer, so every delete sat for
// trashTimeout and then failed. trash(8) needs no Finder, so it must be tried
// first and must actually move the file.
func TestMoveToTrashViaBinaryMovesFile(t *testing.T) {
	if _, err := os.Stat(trashBinary); err != nil {
		t.Skipf("%s not present on this macOS version", trashBinary)
	}

	dir := t.TempDir()
	probe, err := os.CreateTemp(dir, "mole-trash-binary-probe-*.txt")
	if err != nil {
		t.Fatalf("failed to create unique target: %v", err)
	}
	target := probe.Name()
	if _, err := probe.WriteString("probe"); err != nil {
		_ = probe.Close()
		t.Fatalf("failed to seed target: %v", err)
	}
	if err := probe.Close(); err != nil {
		t.Fatalf("failed to close target: %v", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("failed to resolve user home: %v", err)
	}
	trashCopy := filepath.Join(home, ".Trash", filepath.Base(target))
	if _, err := os.Lstat(trashCopy); !os.IsNotExist(err) {
		t.Fatalf("refusing Trash probe basename collision at %s", trashCopy)
	}
	t.Cleanup(func() {
		_ = os.Remove(trashCopy)
	})

	if err := moveViaPlatformBinary(context.Background(), target); err != nil {
		t.Fatalf("moveViaPlatformBinary failed: %v", err)
	}

	if _, err := os.Lstat(target); !os.IsNotExist(err) {
		t.Fatalf("expected %s to be gone, Lstat returned %v", target, err)
	}
}

// trash(8) takes no "--" separator. Passing one makes it report a missing file
// named "--" and exit non-zero while still trashing the real target, which would
// make moveToTrash fall through to Finder and delete a second time. Absolute
// paths cannot be mistaken for options, so no separator is used.
func TestMoveToTrashViaBinaryUsesAbsolutePathWithoutSeparator(t *testing.T) {
	data, err := os.ReadFile("trash_darwin.go")
	if err != nil {
		t.Fatalf("failed to read trash_darwin.go: %v", err)
	}
	if strings.Contains(string(data), `trashBinary, "--"`) {
		t.Error(`trashBinary must not pass "--" to trash(8); it trashes the target but exits non-zero`)
	}
}
