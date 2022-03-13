//go:build !windows && !plan9
// +build !windows,!plan9

package main

import (
	"os"
	"syscall"
	"time"
)

func getCreationTime(entry os.DirEntry) (time.Time, error) {
	info, err := entry.Info()
	if err != nil {
		return time.Now(), err
	}
	creationTime := info.Sys().(*syscall.Stat_t).Ctim
	return time.Unix(creationTime.Sec, creationTime.Nsec), err
}
