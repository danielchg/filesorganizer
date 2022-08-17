package filesmanagement

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func ListFiles(files_src_path string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(files_src_path)
	if err != nil {
		log.Fatal(err)
	}
	return files, nil
}

func GetDestPath(file os.FileInfo, files_dest_path string) (string, error) {

	year := strconv.Itoa(file.ModTime().Year())
	month := file.ModTime().Month().String()
	new_dest_path := files_dest_path + "/" + year + "/" + month
	return new_dest_path, nil
}

func CreateNewPath(new_path string) error {
	if _, err := os.Stat(new_path); errors.Is(err, os.ErrNotExist) {
		fmt.Println("New folder created:", new_path)
		err := os.MkdirAll(new_path, os.ModePerm)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func CopyFile(src, dest string) (int64, error) {

	fmt.Println("Copying file ", src, "to destination ", dest)

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
