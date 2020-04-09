# Scrappy

In its current form, Scrappy is a basic command line application for downloading textual information from web pages. With Scrappy, you can:

- Save to .txt files the following from a web page:
  - page link URLs
  - image link URLs
  - text
  - HTML copy of the page

## Contents

- [Use](##-use)
- [TODO](##-todo)
- [Building](##-building)
- [Dependencies](##-dependencies)

## Use

Simply [build](##-building) and execute the program. Choose menu options by entering their corresponding numbers and enter the address of the web page you want.

## TODO

The following list contains features or developments I would like to implement at some point in time:

- User-facing features:
  - Add web site-wide functionality similar the web page functionality mentioned above
- Internal developments:
  - Implement bluemonday sanitation for all input
  - Implement robotstxt to get information from robotstxt files, for e.g. determining and setting appropriate crawl delays - a big blocker for automating scraping

## Building

Run `go build -o scrappy main.go` to build the project into a binary called scrappy

## Dependencies

The following dependencies are required for this application:

- Go's [html pacakge](https://godoc.org/golang.org/x/net/html) for HTML functionality
- temoto's [robotstxt package](https://github.com/temoto/robotstxt) for robotstxt parsing functionality
- microcosm's [bluemonday package](https://github.com/microcosm-cc/bluemonday)for sanitation

To download them for development, run the following command in the root directory:

```@bash
go get golang.org/x/net/html; go get github.com/temoto/robotstxt; go get github.com/microcosm-cc/bluemonday
```
