package main

import "fmt"

func main() {
L:
	for {
		for {
			break L
		}
		fmt.Println("Am I dead?")
	}
}
