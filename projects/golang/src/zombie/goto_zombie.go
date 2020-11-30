package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Cannot enter arguments.")
		goto L
	}
	return
L:
	fmt.Println("Am I dead?")
}
