package main

import "fmt"

const (
	MAJOR     = 1
	MINOR     = 1
	PATCH     = 0
	BUILD     = "1437183466710838581"
	CODEWORD  = "growling dionisia"
)

func print_version() {
	fmt.Printf("%d.%d.%d+%s : '%s'\n ", MAJOR, MINOR, PATCH, BUILD, CODEWORD)
}

