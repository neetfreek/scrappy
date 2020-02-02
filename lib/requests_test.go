package lib

import (
	"reflect"
	"testing"
)

var testURL = "http://google.com"

func TestPageData(t *testing.T) {
	got := PageData(testURL)

	typeGot := reflect.TypeOf(got).String()

	if typeGot != "[]uint8" {
		t.Errorf("GetPageResponseData(%s) expected type []uint8, got %s", testURL, typeGot)
	}
}

func TestPageString(t *testing.T) {
	got := PageString(testURL)

	typeGot := reflect.TypeOf(got).String()

	if typeGot != "string" {
		t.Errorf("GetPageResponseString(%s) expected type string, got %s", testURL, typeGot)
	}
}
