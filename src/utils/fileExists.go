package utils

import "os"

// FileExists function check if the given file exists in the system and if it is not a directory.
func FileExists(f string) bool {
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
