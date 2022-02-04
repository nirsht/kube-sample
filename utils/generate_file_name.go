package utils

import (
	"fmt"
	"os"
)

func GenerateFileName(wantedFileName string) ([]byte, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	index := 0
	fileName := wantedFileName
	for {
		if index != 0 {
			fileName = fmt.Sprintf("%v (%v)", wantedFileName, index)
		}
		if !IsFileExists(path, fileName) {
			return []byte(fileName), nil
		}
		index += 1
	}
}
