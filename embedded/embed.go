package main

//static char* getCChar() {
//   return "CChar";
//}
import "C"

// No newline above import "C"

import (
	"fmt"
	"os"
)

func main() {
	char, err := C.getCChar()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", C.GoString(char))
}
