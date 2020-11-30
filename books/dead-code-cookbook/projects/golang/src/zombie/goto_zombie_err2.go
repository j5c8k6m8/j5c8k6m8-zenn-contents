package main

import "fmt"

func main() {
	goto A
A:
	goto C
B:
	fmt.Println("Am I dead?")
	goto C
C:
	return
}
