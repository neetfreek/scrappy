package lib

import (
	"os"
	"testing"
)

var testString = "test"
var testFileExtension = ".html"

func TestBytesBuffer(t *testing.T) {
	got := BytesBuffer([]byte(testString))
	if got.String() != testString {
		t.Errorf("TestBytesBuffer expected %s, got %s", testString, got.String())
	}
}

func TestWriteToFile(t *testing.T) {
	WriteToFile(testString, "")

	got, err := os.Stat(testString + testFileExtension)
	if err != nil {
		t.Errorf("writeToFile failed to find %s", testString+testFileExtension)
	} else if got != nil {
		if os.IsNotExist(err) {
			t.Errorf("writeToFile failed to find %s", testString+testFileExtension)
		}
	}

	os.Remove(testString + testFileExtension)
}
