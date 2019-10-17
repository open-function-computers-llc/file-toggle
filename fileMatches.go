package main

func fileMatches(pathA, pathB string) bool {
	hashA, _ := getFileHash(pathA)
	hashB, _ := getFileHash(pathB)
	return hashA == hashB
}
