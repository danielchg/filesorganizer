package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	fm "github.com/danielchg/filesorganizer/pkg/filesmanagement"
)

var wg sync.WaitGroup

func ProcessingFiles(ch chan os.FileInfo, files_src_path string, files_dest_path string) {
	defer wg.Done()
	fl := <-ch
	src := files_src_path + "/" + fl.Name()
	fmt.Println("File to move: ", fl.Name())

	dest, _ := fm.GetDestPath(fl, files_dest_path)
	fm.CreateNewPath(dest)

	dest_file := dest + "/" + fl.Name()
	fm.CopyFile(src, dest_file)
}

func main() {

	files_src_path := flag.String("src", "./", "The path where there are the files to be organized")
	files_dest_path := flag.String("dest", "./", "The path where the files will be organized")
	flag.Parse()

	files, _ := fm.ListFiles(*files_src_path)
	ch := make(chan os.FileInfo)
	for _, file := range files {
		ch <- file
		wg.Add(len(files))
		go ProcessingFiles(ch, *files_src_path, *files_dest_path)
	}
	wg.Wait()
}
