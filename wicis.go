package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var programPath string

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("Please provide an arguments!")

		return
	}

	file := arguments[1]

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)

		//Does it exists?
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()

			// Is it a regular file?
			if mode.IsRegular() {
				// Is is executable?
				if mode&0111 != 0 {
					programPath = fullPath
					fmt.Println("Tah deuleu di dieu mang:", fullPath)
					return
				}
			}
		}
	}

	if programPath == "" {
		fmt.Println("Teu kapanggih euy teuing di mana")
	}
}
