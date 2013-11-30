package main

import (
	"fmt"
	"regexp"
)

var validURL = regexp.MustCompile(`^http(s)?://`)

func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.


	fmt.Println(validURL.MatchString("http://google.com"))
	fmt.Println(validURL.MatchString("https://google.com"))
	fmt.Println(validURL.MatchString("/foobar"))
	fmt.Println(validURL.MatchString("random"))

  if validURL.MatchString("http://google.com") {
    fmt.Println("valid")
  }
}
