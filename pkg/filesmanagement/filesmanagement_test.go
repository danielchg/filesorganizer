package filesmanagement

import (
	"os"
	"testing"
)

func TestListFiles(t *testing.T) {
	fileList, _ := ListFiles("../../samples")
	if len(fileList) != 8 {
		t.Errorf("The list of files are not correct")
	}
}

func TestGetDestPath(t *testing.T) {
	desiredPath := "tests/2019/January"
	file, _ := os.Stat("../../samples/2019-01-27 06.08.07-1-2.jpg")
	testedPath, _ := GetDestPath(file, "tests")
	if testedPath != desiredPath {
		t.Errorf("The destination path has not been generated correctly, %v differ from %v", testedPath, desiredPath)
	}

}
