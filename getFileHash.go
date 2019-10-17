package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// maybe someday i'll cache the file hashes so that we don't rehash the same file over and over
// var fileHashes map[string]string

// thanks https://www.mrwaggel.be/post/generate-md5-hash-of-a-file-in-golang/
func getFileHash(path string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open(path)
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil
}
