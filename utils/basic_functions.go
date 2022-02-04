package utils

import (
	"errors"
	"fmt"
	"os"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IsFileExists(directory string, filename string) bool {
	filePath := fmt.Sprintf("%v/%v", directory, filename)
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return true
	}
	return false
}
