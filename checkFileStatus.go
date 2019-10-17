package main

import (
	"encoding/json"
	"net/http"
)

func checkFileStatusHandler(w http.ResponseWriter, r *http.Request) {
	current, err := currentVersion()
	if err != nil {
		sendError(w, err)
		return
	}
	output := map[string]interface{}{
		"path":           fileToWrite,
		"currentVersion": current,
	}
	bytes, _ := json.Marshal(output)

	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}
