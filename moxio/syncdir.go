package moxio

import (
	"fmt"
	"os"
)

// SyncDir opens a directory and syncs its contents to disk.
func SyncDir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return fmt.Errorf("open directory: %v", err)
	}
	xerr := d.Sync()
	d.Close()
	return xerr
}
