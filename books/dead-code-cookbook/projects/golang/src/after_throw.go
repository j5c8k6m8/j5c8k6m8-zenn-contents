package main

import "fmt"

func main() {
	defer func() {
		recover()
	}()
	panic("Error")
	fmt.Println("Am I dead?")
}
