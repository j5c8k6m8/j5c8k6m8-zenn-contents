package main

import "fmt"

func main() {
	goto L
	fmt.Println("Am I dead?")
L:
	return
}
