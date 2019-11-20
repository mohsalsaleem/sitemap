package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("url", "", "Enter the url of the website to proceed")
	flag.Parse()
	fmt.Printf("Building the sitemap of %s\n", *url)
}
