package main

import (
	"log"
	"os"
	"path/filepath"
)

const assetDirectory = "Assets"

func main() {
	assetPath := assetDirectory

	if len(assetPath) > 0 {
		fileInfo, err := os.Stat(assetPath)
		if err != nil {
			log.Fatal(err)
		}

		if !fileInfo.IsDir() {
			log.Fatal("main: ", assetPath, " is not a directory")
		}
	}

	someAsset := "level01.txt"
	if _, err := os.Stat(filepath.Join(assetPath, someAsset)); err != nil {
		log.Fatal(err)
	}
}
