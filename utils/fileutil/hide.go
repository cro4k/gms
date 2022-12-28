//go:build !windows

package fileutil

import (
	"os"
	"strings"
)

func Hide(path string) error {
	if strings.HasPrefix(path, ".") {
		return nil
	}
	return os.Rename(path, "."+path)
}
