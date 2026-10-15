//go:build linux

package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/sys/unix"
)

func moveViaPlatformBinary(ctx context.Context, absPath string) error {
	var cmd *exec.Cmd
	if path, err := exec.LookPath("gio"); err == nil {
		cmd = exec.CommandContext(ctx, path, "trash", absPath)
	} else if path, err := exec.LookPath("trash-put"); err == nil {
		cmd = exec.CommandContext(ctx, path, absPath)
	} else {
		return fmt.Errorf("no supported trash utility found (gio or trash-put)")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("timeout moving to Trash")
		}
		return fmt.Errorf("failed to move to Trash: %s", strings.TrimSpace(string(output)))
	}
	return nil
}

func atomicRenameExclusive(oldpath, newpath string) error {
	return unix.Renameat2(unix.AT_FDCWD, oldpath, unix.AT_FDCWD, newpath, unix.RENAME_NOREPLACE)
}

func trashDirectoryForPath(absPath string) (string, error) {
	home, err := os.UserHomeDir()
	if err == nil {
		legacyTrash := filepath.Join(home, ".Trash")
		if _, statErr := os.Lstat(legacyTrash); statErr == nil {
			if err := ensureOwnedTrashDirectory(legacyTrash, false); err != nil {
				return "", err
			}
			return legacyTrash, nil
		}
	}

	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		if home == "" {
			var homeErr error
			home, homeErr = os.UserHomeDir()
			if homeErr != nil {
				return "", fmt.Errorf("failed to resolve home directory: %w", homeErr)
			}
		}
		dataHome = filepath.Join(home, ".local", "share")
	}

	trashDir := filepath.Join(dataHome, "Trash", "files")
	if err := ensureOwnedTrashDirectory(trashDir, true); err != nil {
		return "", err
	}
	return trashDir, nil
}

func moveToTrashFallback(absPath string) error {
	return fmt.Errorf("cannot move to trash: neither gio nor filesystem move succeeded for %s", absPath)
}
