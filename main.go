package main

import (
	"fmt"
	"mcmanager/internal/curse"
	"net/url"
)

func main() {
	c := &curse.DefaultConfig

	url := &url.URL{}

	newUrl := c.ToURLParams(url)

	fmt.Printf("%v", newUrl)
}
