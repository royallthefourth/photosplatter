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
	creationTime := info.Sys().(*syscall.Win32FileAttributeData).CreationTime
	// I don't have a good way to test the Windows details. Help wanted!
	return time.Unix(creationTime.Sec, creationTime.Nsec), err
}
