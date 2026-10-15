//go:build darwin

package main

import (
	"syscall"
	"time"
)

func statAccessTime(stat *syscall.Stat_t) time.Time {
	return time.Unix(stat.Atimespec.Sec, stat.Atimespec.Nsec)
}
