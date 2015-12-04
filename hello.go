package main

// Package declaration, every program must start with one and they identify and help reuse code 

import "fmt"

// Import is a keyword that tells go to import code from other packages

// This is obviously comments for single lines 

/* and this 
is multiline
comments */ 

func main() {
// functions are the building block of go and main is special as it gets called when executed
    fmt.Printf("hello, world\n")
// https://golang.org/pkg/fmt/ Package fmt implements formatted I/O with functions analogous to C's printf and scanf
}

