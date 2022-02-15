package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("Please provide an arguments!")

		return
	}

	files := arguments[1:]

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	pathLists := make(map[string][]string)

	for _, file := range files {
		var pathList []string
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
						pathList = append(pathList, fullPath)
					}
				}
			}
		}
		pathLists[file] = pathList
	}

	for k, v := range pathLists {
		if len(v) == 0 {
			fmt.Println("=================")
			fmt.Printf("%s Teu kapanggih euy teuing di mana\n", k)
			fmt.Println("=================")
			continue
		}
		fmt.Println("=================")
		fmt.Printf("Mun keur %s Tah deuleu di dieu mang:\n", k)

		for _, str := range v {
			fmt.Println(str)
		}
	}
}
