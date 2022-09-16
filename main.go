package main

import (
	"flag"
	"fmt"
	"io/fs"

	fm "github.com/danielchg/filesorganizer/pkg/filesmanagement"
)

func worker(id int, files_src_path, files_dest_path *string, jobs <-chan fs.FileInfo, quit <-chan struct{}) {

	for {
		select {
		case file := <-jobs:
			src := *files_src_path + "/" + file.Name()
			fmt.Println("File to move: ", src, "by Worker: ", id)

			dest, _ := fm.GetDestPath(file, *files_dest_path)
			fm.CreateNewPath(dest)

			dest_file := dest + "/" + file.Name()
			fm.CopyFile(src, dest_file)

		case <-quit:
			return
		}
	}

}

func main() {
	files_src_path := flag.String("src", "./", "The path where there are the files to be organized")
	files_dest_path := flag.String("dest", "./", "The path where the files will be organized")
	workers := flag.Int("w", 1, "Number of workers to parallelize file processing")
	flag.Parse()

	jobs := make(chan fs.FileInfo)
	quit := make(chan struct{})
	for w := 1; w <= *workers; w++ {
		go worker(w, files_src_path, files_dest_path, jobs, quit)
	}

	files, _ := fm.ListFiles(*files_src_path)
	for _, file := range files {
		jobs <- file
	}

	close(quit)
}
