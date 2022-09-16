package filesmanagement

import "testing"

func TestListFiles(t *testing.T) {
	fileList, _ := ListFiles("../../samples")
	if len(fileList) != 7 {
		t.Errorf("The list of files are not correct")
	}
}
