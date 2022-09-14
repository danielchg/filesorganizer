package main

import (
	"flag"
	"fmt"
	"io/fs"

	fm "github.com/danielchg/filesorganizer/pkg/filesmanagement"
)

func worker(id int, files_src_path, files_dest_path *string, jobs <-chan fs.FileInfo, result chan<- string) {
	file := <-jobs
	src := *files_src_path + "/" + file.Name()
	fmt.Println("File to move: ", src, "by Worker: ", id)

	dest, _ := fm.GetDestPath(file, *files_dest_path)
	fm.CreateNewPath(dest)

	dest_file := dest + "/" + file.Name()
	fm.CopyFile(src, dest_file)

	result <- dest_file

}

func main() {
	files_src_path := flag.String("src", "./", "The path where there are the files to be organized")
	files_dest_path := flag.String("dest", "./", "The path where the files will be organized")
	flag.Parse()

	numWorkers := 3
	jobs := make(chan fs.FileInfo)
	result := make(chan string)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, files_src_path, files_dest_path, jobs, result)
	}

	files, _ := fm.ListFiles(*files_src_path)
	filesCount := len(files)
	for _, file := range files {
		jobs <- file
	}

	for a := 1; a <= filesCount; a++ {
		<-result
	}
	close(jobs)
}
