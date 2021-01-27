package main

import (
	"fmt"
	"lite/repl"
	"os"
)

func main() {
	fmt.Printf("Sup! This is the Lite programming language!\n")
	fmt.Printf("Feel free to type in something\n")
	repl.Start(os.Stdin, os.Stdout)
}
