package requests

import (
	"reflect"
	"testing"
)

var testURL = "http://google.com"

func TestGetPageResponseData(t *testing.T) {
	got := GetPageResponseData(testURL)

	typeGot := reflect.TypeOf(got).String()

	if typeGot != "[]uint8" {
		t.Errorf("GetPageResponseData(%s) expected type []uint8, got %s", testURL, typeGot)
	}
}

func TestGetPageResponseString(t *testing.T) {
	got := GetPageResponseString(testURL)

	typeGot := reflect.TypeOf(got).String()

	if typeGot != "string" {
		t.Errorf("GetPageResponseString(%s) expected type string, got %s", testURL, typeGot)
	}
}
