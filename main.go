package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var fileVersions map[int]string
var fileToWrite string

func main() {
	err := godotenv.Load()
	fileToWrite = os.Getenv("FILE_TO_WRITE")

	fileVersions = map[int]string{}
	versions := strings.Split(os.Getenv("FILE_VERSIONS"), "|")
	for i := 0; i < len(versions); i++ {
		fileVersions[i] = versions[i]
	}

	// all file and file versions must exist
	if !fileExists(fileToWrite) {
		panic("target files must exist: " + fileToWrite)
	}
	for _, path := range fileVersions {
		if !fileExists(path) {
			panic("file version must exist: " + path)
		}
	}

	_, err = currentVersion()
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/", checkFileStatusHandler)
	http.HandleFunc("/change", incrementVersionHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
