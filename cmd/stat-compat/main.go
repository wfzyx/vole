package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	args := os.Args[1:]
	var format string
	var files []string
	followSymlinks := false

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "-L" {
			followSymlinks = true
		} else if strings.HasPrefix(arg, "-f") {
			if len(arg) > 2 {
				format = arg[2:]
			} else if i+1 < len(args) {
				format = args[i+1]
				i++
			}
		} else if !strings.HasPrefix(arg, "-") {
			files = append(files, arg)
		}
	}

	for _, file := range files {
		var stat syscall.Stat_t
		var err error
		if followSymlinks {
			err = syscall.Stat(file, &stat)
		} else {
			err = syscall.Lstat(file, &stat)
		}
		if err != nil {
			os.Exit(1)
		}

		out := format
		out = strings.ReplaceAll(out, "%z", strconv.FormatInt(stat.Size, 10))
		out = strings.ReplaceAll(out, "%d", strconv.FormatUint(uint64(stat.Dev), 10))
		out = strings.ReplaceAll(out, "%i", strconv.FormatUint(stat.Ino, 10))
		out = strings.ReplaceAll(out, "%u", strconv.FormatUint(uint64(stat.Uid), 10))
		out = strings.ReplaceAll(out, "%l", strconv.FormatUint(uint64(stat.Nlink), 10))
		out = strings.ReplaceAll(out, "%m", strconv.FormatInt(stat.Mtim.Sec, 10))
		out = strings.ReplaceAll(out, "%Lp", fmt.Sprintf("%03o", stat.Mode&0o7777))

		if strings.Contains(out, "%Su") {
			username := strconv.FormatUint(uint64(stat.Uid), 10)
			if u, err := user.LookupId(username); err == nil {
				username = u.Username
			}
			out = strings.ReplaceAll(out, "%Su", username)
		}

		fmt.Println(out)
	}
}
