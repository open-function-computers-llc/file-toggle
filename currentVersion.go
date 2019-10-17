package main

import "errors"

func currentVersion() (int, error) {
	currentIndex := 0
	for key, path := range fileVersions {
		if fileMatches(path, fileToWrite) {
			return key, nil
		}
	}
	return currentIndex, errors.New("the file to write doesn't match any of the variations")
}
