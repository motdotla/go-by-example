package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println(err)
	}
	defer dir.Close()

	fmt.Println(dir)

	filenames := make([]string, 0)

	fis, err := dir.Readdir(-1) // -1 returns all the Fileinfos
	if err != nil {
		fmt.Println(err)
	}

	for _, fileinfo := range fis {
		if !fileinfo.IsDir() {
			filenames = append(filenames, fileinfo.Name())
		}
	}

	fmt.Println(filenames)
}
