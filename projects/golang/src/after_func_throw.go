package main

import "fmt"

func main() {
	defer func() {
		recover()
	}()
	f()
	fmt.Println("Am I dead?")
}

func f() {
	panic("Error")
}
