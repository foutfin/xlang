package main

import (
	"boolang/repl"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("BOOLANG BOOLANG\n\n")
	repl.Start(os.Stdin, os.Stdout)
}
