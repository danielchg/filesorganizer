package filesmanagement

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestListFiles(t *testing.T) {
	fileList, _ := ListFiles("../../samples")
	if len(fileList) != 8 {
		t.Errorf("The list of files are not correct")
	}
}

func TestGetDestPath(t *testing.T) {
	desiredPath := "tests/2019/January"
	desiredTime := time.Date(2019, time.January, 27, 6, 7, 2, 0, time.UTC)
	err := os.Chtimes("../../samples/2019-01-27 06.08.07-1-2.jpg", desiredTime, desiredTime)
	if err != nil {
		fmt.Println("Error chaning time of file")
	}
	file, _ := os.Stat("../../samples/2019-01-27 06.08.07-1-2.jpg")
	testedPath, _ := GetDestPath(file, "tests")
	if testedPath != desiredPath {
		t.Errorf("The destination path has not been generated correctly, %v differ from %v", testedPath, desiredPath)
	}

}
