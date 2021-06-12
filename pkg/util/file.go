package util

import "os"

func doesExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsExist(err); err != nil {
		return false
	}
	return true
}
