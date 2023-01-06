package utils

import "os"

func VerifyIfFileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		return false
	}
	return true
}
