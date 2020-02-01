package main

import (
	inputoutput "http-testing/lib/lib_inputoutput"
	requests "http-testing/lib/lib_requests"
)

func main() {
	url := "https://news.ycombinator.com/item?id=22201337"
	inputoutput.WriteToFile("test", requests.GetPageResponseString(url))
}
