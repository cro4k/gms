//go:build windows

package fileutil

import (
	"syscall"
)

func Hide(path string) error {
	name, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	err = syscall.SetFileAttributes(name, syscall.FILE_ATTRIBUTE_HIDDEN)
	return err
}
