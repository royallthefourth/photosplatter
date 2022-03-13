//go:build !windows && !plan9
// +build !windows,!plan9

package gallery

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
	// these typecasts are necessary for 32 bit compile targets
	return time.Unix(int64(creationTime.Sec), int64(creationTime.Nsec)), err
}
