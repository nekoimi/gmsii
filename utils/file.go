package utils

import "os"

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err == nil {
		return true
	}
	return false
}
