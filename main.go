package main

import (
	"flag"
	"fmt"

	fm "github.com/danielchg/filesorganizer/pkg/filesmanagement"
)

func main() {
	files_src_path := flag.String("src", "./", "The path where there are the files to be organized")
	files_dest_path := flag.String("dest", "./", "The path where the files will be organized")
	flag.Parse()

	files, _ := fm.ListFiles(*files_src_path)
	for _, file := range files {

		src := *files_src_path + "/" + file.Name()
		fmt.Println("File to move: ", file.Name())

		dest, _ := fm.GetDestPath(file, *files_dest_path)
		fm.CreateNewPath(dest)

		dest_file := dest + "/" + file.Name()
		fm.CopyFile(src, dest_file)
	}
}
