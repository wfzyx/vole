//go:build linux

package main

import (
	"syscall"
	"time"
)

func statAccessTime(stat *syscall.Stat_t) time.Time {
	return time.Unix(stat.Atim.Sec, stat.Atim.Nsec)
}
