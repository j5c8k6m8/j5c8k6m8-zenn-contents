package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Cannot enter arguments.")
		return
	}
	fmt.Println("Am I dead?")
}
