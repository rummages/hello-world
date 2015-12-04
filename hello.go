package main

// Package declaration, every program must start with one and they identify and help reuse code 

import "fmt"

// Import is a keyword that tells go to import code from other packages or libraries

// This is obviously comments for single lines 

/* and this 
is multiline
comments */ 

func main() {
// functions are the building block of go and main is special as it gets called when executed
    fmt.Printf("hello, world\n")
// https://golang.org/pkg/fmt/ Package fmt implements formatted I/O with functions analogous to C's printf and scanf
}

/* Created an executable program that imports the 'fmt' package ( docs can be found at https://golang.org/pkg/fmt/ ) 
Several practice comment lines to help me learn, and to get into the habit of documenting codebase
Function Main does one thing, print out the string hello world as is standard starting programming practice
You can also learn more by using godoc at the command line i.e. godoc fmt Println and it will provide documentation 
$ godoc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.
Also did a go run to compile and run the executable, everything went fine. 
*/ 
