package main

import (
	"fmt"
	"net/http"
)

func incrementVersion() error {
	current, err := currentVersion()
	if err != nil {
		return err
	}

	next := current + 1
	if next > len(fileVersions)-1 {
		next = 0
	}
	fmt.Println(fileToWrite)

	return copy(fileVersions[next], fileToWrite)
}

func incrementVersionHandler(w http.ResponseWriter, r *http.Request) {
	incrementVersion()
	runAdditionalTasks()
	w.Write([]byte("file changed"))
}
