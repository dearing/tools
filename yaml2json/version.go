package main

import "fmt"

const (
	MAJOR     = 1
	MINOR     = 0
	PATCH     = 0
	BUILD     = "1437183466716714282"
	CODEWORD  = "enormous esmenet"
)

func print_version() {
	fmt.Printf("%d.%d.%d+%s : '%s'\n ", MAJOR, MINOR, PATCH, BUILD, CODEWORD)
}

