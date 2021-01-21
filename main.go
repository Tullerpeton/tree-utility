package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getDirFiles(output io.Writer, prefix, pwd string, printFiles bool) {
	files, err := ioutil.ReadDir(pwd)
	if err != nil {
		fmt.Printf("Panic (read directory: %s): %v\n", pwd, err)
	}

	if !printFiles {
		var dirs []os.FileInfo
		for _, file := range files {
			if file.IsDir() {
				dirs = append(dirs, file)
			}
		}
		files = dirs
	}

	length := len(files)
	for i, file := range files {
		if file.IsDir() {
			var prefixNew string

			if length > i + 1 {
				fmt.Fprintf(output, prefix+"├───%s\n", file.Name())
				prefixNew = prefix + "│\t"
			} else {
				fmt.Fprintf(output, prefix+"└───%s\n", file.Name())
				prefixNew = prefix + "\t"
			}
			getDirFiles(output, prefixNew, filepath.Join(pwd, file.Name()), printFiles)
		} else {
			if file.Size() > 0 {
				if length > i + 1 {
					fmt.Fprintf(output, prefix+"├───%s (%vb)\n", file.Name(), file.Size())
				} else {
					fmt.Fprintf(output, prefix+"└───%s (%vb)\n", file.Name(), file.Size())
				}
			} else {
				if length > i + 1 {
					fmt.Fprintf(output, prefix+"├───%s (empty)\n", file.Name())
				} else {
					fmt.Fprintf(output, prefix+"└───%s (empty)\n", file.Name())
				}
			}
		}
	}
}

func dirTree(output io.Writer, path string, printFiles bool) error {
	getDirFiles(output, "", path, printFiles)
	return nil
}

func main() {
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(os.Stdout, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}