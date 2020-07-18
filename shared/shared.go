package main

import (
	"C"
	"fmt"
)

//export Print
func Print() {
	fmt.Println("Hello world")
}

func main() {}
