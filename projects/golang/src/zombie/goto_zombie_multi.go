package main

import "fmt"

func main() {
	goto A
A:
	goto D
B:
	fmt.Println("Am I dead?")
	goto C
C:
	fmt.Println("Am I dead?")
	goto B
D:
	return
}
