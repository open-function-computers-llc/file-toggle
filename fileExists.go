package main

import "os"

func fileExists(path string) bool {
	var exists bool
	if _, err := os.Stat(path); err == nil {
		exists = true
	}

	return exists
}
